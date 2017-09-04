package scaler

import (
	"github.com/astaxie/beego/orm"
	"weibo.com/opendcp/orion/models"
)

func onlineNodesList(pid int) (int, []*models.NodeState, []*models.NodeState, []*models.NodeState, []*models.NodeState, []*models.NodeState, error) {
	var (
		init    = make([]*models.NodeState, 0)
		ok      = make([]*models.NodeState, 0)
		failed  = make([]*models.NodeState, 0)
		running = make([]*models.NodeState, 0)
		stopped = make([]*models.NodeState, 0)
		online  = make([]*models.NodeState, 0)
		o       = orm.NewOrm()
	)

	count, err := o.QueryTable(&models.NodeState{}).Filter("deleted", false).
		Filter("Pool", pid).Filter("NodeType", models.Crontab).
		All(&online)

	if err != nil && err == orm.ErrNoRows {
		return int(count), init, ok, failed, running, stopped, nil
	}

	//get status of nodeStates
	for _, n := range online {
		switch n.Status {
		case models.STATUS_INIT:
			init = append(init, n)
		case models.STATUS_RUNNING:
			running = append(running, n)
		case models.STATUS_SUCCESS:
			ok = append(ok, n)
		case models.STATUS_FAILED:
			failed = append(failed, n)
		case models.STATUS_STOPPED:
			stopped = append(stopped, n)
		}
	}

	return int(count), init, ok, failed, running, stopped, err
}
