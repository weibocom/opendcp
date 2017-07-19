package scaler

import (
	"fmt"

	"weibo.com/opendcp/orion/models"

	"github.com/astaxie/beego/orm"
)

// blame: who designed the table ?
type Node struct {
	n models.Node
	s models.NodeState
}

func onlineNodesList(pid int) ([]*Node, error) {
	var (
		nodes  []models.Node
		online []*Node
		o      = orm.NewOrm()
	)

	if _, err := o.QueryTable(&models.Node{}).
		Filter("Pool", pid).Filter("NodeType", models.Crontab).
		All(&nodes); err != nil {
		if err == orm.ErrNoRows { // ignore no row found
			err = nil
		}
		return nil, err
	}

	for _, n := range nodes {
		nn := &Node{n: n}
		if err := o.QueryTable(&models.NodeState{}).Filter("Node", n.Id).
			Filter("Pool", pid).OrderBy("-CreatedTime").One(&nn.s); err != nil {
			return nil, fmt.Errorf("db query NodeState failed: %v", err)
		}
		online = append(online, nn)
	}
	return online, nil
}
