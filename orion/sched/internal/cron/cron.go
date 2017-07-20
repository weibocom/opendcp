package cron

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/astaxie/beego"
	"github.com/davecgh/go-spew/spew"
	"github.com/robfig/cron"

	"weibo.com/opendcp/orion/models"
)

type configs interface {
	Config(id int) *models.ExecTask
}

// Cron represents a task which is like crontable
type Cron struct {
	mu     sync.Mutex
	wg     sync.WaitGroup
	id     int
	cfgs   configs
	cron   *cron.Cron
	cancel context.CancelFunc
}

// New return a new instance of Cron
func New(id int, cfgs configs) *Cron {
	return &Cron{
		id:   id,
		cfgs: cfgs,
		cron: cron.New(),
	}
}

func (r *Cron) Start() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	beego.Info(spew.Sprintf("crontable task(%d) start", r.id))
	ctx, cancel := context.WithCancel(context.Background())
	r.cancel = cancel
	ctx = context.WithValue(ctx, "wg", &r.wg)

	if err := r.addJobs(ctx); err != nil {
		cancel()
		return err
	}

	r.cron.Start()
	return nil
}

func (r *Cron) Stop() {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.cancel != nil {
		r.cancel()
	}
	r.cron.Stop()
	r.wg.Wait()
}

func (r *Cron) addJobs(ctx context.Context) error {
	cfg := r.cfgs.Config(r.id)
	if cfg == nil {
		return fmt.Errorf("cannot get %d config", r.id)
	}

	l := len(cfg.CronItems)
	if l == 0 {
		beego.Warn(spew.Sprintf("%+v Start a empty cronitem runner", cfg))
		return nil
	}

	sort.Sort(models.CronItemSlice(cfg.CronItems))

	if idx := findNonIgnoredCronItem2Check(cfg.CronItems); !cfg.CronItems[idx].Ignore {
		// must run once at now to check there are enough machines
		(&cronJob{ctx: ctx, cfgs: r.cfgs, id: r.id, idx: idx}).Run()
	}

	for i, item := range cfg.CronItems {
		if item.Ignore {
			continue
		}
		if err := r.cron.AddJob(toCronExpression(item), &cronJob{
			ctx:  ctx,
			cfgs: r.cfgs,
			id:   r.id,
			idx:  i,
		}); err != nil {
			return fmt.Errorf("cron add job failed: %v", err)
		}
	}
	return nil
}

func findNonIgnoredCronItem2Check(items []*models.CronItem) int {
	l := len(items)
	n := time.Now()
	weekday := int(n.Weekday()) + 1
	idx := sort.Search(l, func(i int) bool {
		if wd := items[i].WeekDay; wd != 0 && weekday != wd {
			return weekday < items[i].WeekDay
		}
		token := strings.Split(items[i].Time, ":")
		ih, _ := strconv.Atoi(token[0])
		im, _ := strconv.Atoi(token[1])
		if ih == n.Hour() {
			return n.Minute() < im
		}
		return n.Hour() < ih
	})

	// pick up which cron item should be checked
	idx += l - 1
	for i := 0; i < l; i++ {
		if !items[idx%l].Ignore {
			break
		}
		idx--
	}
	return idx % l
}

// the every field of models.CronItem must be validated by input api,
// otherwise it return a bad expression.
func toCronExpression(item *models.CronItem) string {
	tokens := strings.Split(item.Time, ":")
	h, _ := strconv.Atoi(tokens[0])
	m, _ := strconv.Atoi(tokens[1])
	weekday := "*"
	if item.WeekDay != 0 {
		weekday = strconv.Itoa(item.WeekDay - 1)
	}
	return fmt.Sprintf("00 %02d %02d * * %s", m, h, weekday)
}
