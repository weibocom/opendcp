package sched

import (
	"errors"
	"fmt"
	"sync"

	"weibo.com/opendcp/orion/models"
	"weibo.com/opendcp/orion/sched/internal/cron"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/davecgh/go-spew/spew"
)

// task represents a background running task
type task interface {
	Start() error
	Stop()
}

var (
	ErrStopped      = errors.New("already stopped")
	ErrAlreadyExist = errors.New("already exist")

	Scheduler *scheduler
)

func Initial() error {
	Scheduler = &scheduler{
		tasks: make(map[int]task),
	}
	if err := Scheduler.load(); err != nil {
		return err
	}
	return nil
}

type scheduler struct {
	mu      sync.Mutex
	tasks   map[int]task
	stopped bool
}

func (s *scheduler) load() error {
	var (
		cfgs []*models.ExecTask
		o    = orm.NewOrm()
	)

	if _, err := o.QueryTable(&models.ExecTask{}).
		RelatedSel().All(&cfgs); err != nil {
		return fmt.Errorf("db load ExecTask failed: %v", err)
	}

	for _, cfg := range cfgs {
		if _, err := o.LoadRelated(cfg, "CronItems"); err != nil {
			return fmt.Errorf("db load %d CronItems failed: %v", cfg.Id, err)
		}
		if _, err := o.LoadRelated(cfg, "DependItems"); err != nil {
			return fmt.Errorf("db load %d DependItems failed: %v", cfg.Id, err)
		}
		for _, dep := range cfg.DependItems {
			if err := o.Read(dep.Pool); err != nil {
				return fmt.Errorf("db load %d DependItems Pool failed: %v", err)
			}
		}
	}

	for _, cfg := range cfgs {
		beego.Info(spew.Sprintf("load exec task: %+v", cfg))
		if cfg.ExecType == models.Crontab {
			if err := s.addTask(cfg); err != nil {
				return fmt.Errorf("add task %d failed %v", cfg.Id, err)
			}
		}
	}
	return nil
}

// Create creates a new task by given config, it will update the database
// and launch the backgroud task.
func (s *scheduler) Create(cfg *models.ExecTask) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.stopped {
		return ErrStopped
	}

	if cfg.Id != 0 {
		return ErrAlreadyExist
	}

	if err := s.upsertCfg(cfg); err != nil {
		return fmt.Errorf("db insert failed: %v", err)
	}

	beego.Debug(spew.Sprintf("create task %+v", cfg))

	if cfg.ExecType == models.Crontab {
		if err := s.addTask(cfg); err != nil {
			return fmt.Errorf("add task %d failed: %v", cfg.Id, err)
		}
	}
	return nil
}

// Delete deletes this ExecTask
func (s *scheduler) Delete(cfg *models.ExecTask) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.stopped {
		return ErrStopped
	}

	old, err := s.delCfg(cfg)
	if err != nil {
		return err
	}

	beego.Debug(spew.Sprintf("delete task %+v", old))

	if old.ExecType == models.Crontab {
		s.delTask(old.Id)
	}
	return nil
}

// Update updates the config of the task and restart the task to apply config
func (s *scheduler) Update(cfg *models.ExecTask) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.stopped {
		return ErrStopped
	}

	old, err := s.config(cfg.Id)
	if err != nil {
		return fmt.Errorf("get config failed %v", err)
	}

	if err := s.upsertCfg(cfg); err != nil {
		return fmt.Errorf("update cfg failed: %v", err)
	}

	if old.ExecType == models.Crontab {
		beego.Debug("delete task(%d) ...", cfg.Id)
		s.delTask(old.Id)
		beego.Debug("delete task(%d) done", cfg.Id)
	}

	beego.Debug(spew.Sprintf("update task from (%+v) to (%+v)", old, cfg))

	if cfg.ExecType == models.Crontab {
		if err := s.addTask(cfg); err != nil {
			return err
		}
	}
	return nil
}

// Shutdown stops all running tasks
func (s *scheduler) Shutdown() error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.stopped {
		return ErrStopped
	}
	for _, t := range s.tasks {
		t.Stop()
	}
	s.stopped = true
	return nil
}

func (s *scheduler) addTask(cfg *models.ExecTask) error {
	var t task
	switch cfg.ExecType { // only one type so far
	case models.Crontab:
		t = cron.New(cfg.Id, s)
	case models.Mock:
		beego.Info(spew.Sprintf("add mock task (%+v)", cfg))
	}
	if t != nil {
		if err := t.Start(); err != nil {
			return err
		}
		s.tasks[cfg.Id] = t
	}
	return nil
}

func (s *scheduler) delTask(id int) {
	if t, ok := s.tasks[id]; ok {
		t.Stop()
		delete(s.tasks, id)
	}
}

func (s *scheduler) upsertCfg(cfg *models.ExecTask) (err error) {
	o := orm.NewOrm()
	if err = o.Begin(); err != nil {
		return fmt.Errorf("tx begin failed: %v", err)
	}

	defer func() {
		if err != nil {
			if e := o.Rollback(); e != nil {
				err = fmt.Errorf("%v, rollback failed %v", err, e)
			}
		}
	}()

	if cfg.Id != 0 {
		if _, err = o.QueryTable(&models.CronItem{}).
			Filter("ExecTask", cfg.Id).Delete(); err != nil {
			return
		}
		if _, err = o.QueryTable(&models.DependItem{}).
			Filter("ExecTask", cfg.Id).Delete(); err != nil {
			return
		}
		if _, err = o.Update(cfg); err != nil {
			return
		}
	} else {
		if _, err = o.Insert(cfg); err != nil {
			return
		}
	}

	for _, item := range cfg.CronItems {
		item.Id = 0
		item.ExecTask = cfg
		if _, err = o.Insert(item); err != nil {
			return
		}
	}

	for _, item := range cfg.DependItems {
		item.Id = 0
		item.ExecTask = cfg
		if _, err = o.Insert(item); err != nil {
			return
		}
	}

	err = o.Commit()
	return
}

func (s *scheduler) delCfg(cfg *models.ExecTask) (*models.ExecTask, error) {
	old, err := s.config(cfg.Id)
	if err != nil {
		return nil, fmt.Errorf("db get old failed: %v", err)
	}
	if _, err := orm.NewOrm().Delete(cfg); err != nil {
		return nil, fmt.Errorf("db delete failed: %v", err)
	}
	return old, nil
}

func (s *scheduler) Config(id int) *models.ExecTask {
	for i := 0; i < 3; i++ {
		if cfg, err := s.config(id); err == nil {
			return cfg
		}
	}
	beego.Error("get config %d retry max", id)
	return nil
}

func (s *scheduler) config(id int) (*models.ExecTask, error) {
	var (
		cfg = &models.ExecTask{}
		o   = orm.NewOrm()
	)

	if err := o.QueryTable(cfg).Filter("id", id).RelatedSel().One(cfg); err != nil {
		return nil, err
	}

	if _, err := o.LoadRelated(cfg, "CronItems"); err != nil {
		return nil, err
	}

	if _, err := o.LoadRelated(cfg, "DependItems"); err != nil {
		return nil, err
	}

	for _, dep := range cfg.DependItems {
		if err := o.Read(dep.Pool); err != nil {
			return nil, fmt.Errorf("db load %d DependItems Pool failed: %v", err)
		}
	}

	return cfg, nil
}
