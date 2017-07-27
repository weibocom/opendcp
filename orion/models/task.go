/**
 *    Copyright (C) 2016 Weibo Inc.
 *
 *    This file is part of Opendcp.
 *
 *    Opendcp is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU General Public License as published by
 *    the Free Software Foundation; version 2 of the License.
 *
 *    Opendcp is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU General Public License for more details.
 *
 *    You should have received a copy of the GNU General Public License
 *    along with Opendcp; if not, write to the Free Software
 *    Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA 02110-1301  USA
 */

package models

import (
	"strconv"
	"strings"
	"time"
)

const (
	STATUS_INIT = iota
	STATUS_RUNNING
	STATUS_SUCCESS
	STATUS_FAILED
	STATUS_STOPPED
)

const (
	Manual  = "manual"
	Crontab = "crontab"
	Depend  = "depend"
	Mock    = "mock"
)

const (
	TaskExpend  = "expand"
	TaskShrink  = "shrink"
	TaskDdeploy = "deploy"
)

type ExecTask struct {
	Id          int           `json:"id" orm:"pk;auto"`
	Pool        *Pool         `json:"pool" orm:"rel(fk)"`               //服务池id
	CronItems   []*CronItem   `json:"cron_items" orm:"reverse(many)"`   //定时任务列表
	DependItems []*DependItem `json:"depend_items" orm:"reverse(many)"` //依赖任务列表
	Type        string        `json:"type"`                             //模版任务类型 expand/upload
	ExecType    string        `json:"exec_type"`                        //执行任务类型 crontab/depend
}

type CronItem struct {
	Id           int       `json:"id" orm:"pk;auto"`
	ExecTask     *ExecTask `json:"-" orm:"rel(fk)"`           //定时任务
	InstanceNum  int       `json:"instance_num"`              //扩容缩容使用作为机器的数量
	ConcurrRatio int       `json:"concurr_ratio"`             //上线使用作为最大并发比例
	ConcurrNum   int       `json:"concurr_num"`               //上线使用作为最大并发数
	WeekDay      int       `json:"week_day" orm:"default(0)"` //每周第几天，取值 0 每天,1 周日,2 周一,3 周二,4 周三,5 周四,6 周五,7 周六
	Time         string    `json:"time"`                      //每天时分秒,例如 14:09:08
	Ignore       bool      `json:"ignore" orm:"default(0)"`   //是否忽略定时任务 0 不忽略，1 忽略
}

type CronItemSlice []*CronItem

func (s CronItemSlice) Len() int      { return len(s) }
func (s CronItemSlice) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s CronItemSlice) Less(i, j int) bool {
	if s[i].WeekDay != s[j].WeekDay {
		return s[i].WeekDay < s[j].WeekDay
	}

	itokens := strings.Split(s[i].Time, ":")
	jtokens := strings.Split(s[j].Time, ":")
	ih, _ := strconv.Atoi(itokens[0])
	im, _ := strconv.Atoi(itokens[1])
	jh, _ := strconv.Atoi(jtokens[0])
	jm, _ := strconv.Atoi(jtokens[1])

	if ih == jh {
		return im < jm
	}
	return ih < jh
}

type DependItem struct {
	Id           int       `json:"id" orm:"pk;auto"`
	ExecTask     *ExecTask `json:"-" orm:"rel(fk)"`         //依赖任务
	Pool         *Pool     `json:"pool"  orm:"rel(fk)"`     //依赖服务池id
	Ratio        float64   `json:"ratio"`                   //依赖比例
	ElasticCount int       `json:"elastic_count"`           //冗余机器数量
	StepName     string    `json:"step_name"`               //依赖步骤名称
	Ignore       bool      `json:"ignore" orm:"default(0)"` //是否忽略依赖 0 不忽略，1 忽略
}

//任务流
type Flow struct {
	Id     int    `json:"id" orm:"pk;auto"`
	Name   string `json:"task_name" orm:"size(50);null"`
	Status int    `json:"state"`
	Pool   *Pool  `json:"pool_id" orm:"rel(fk);null;on_delete(set_null)"`
	//Params 		string 		`json:"params"  orm:"type(text)"`
	Options     string    `json:"options" orm:"type(text)"`
	Impl        *FlowImpl `json:"-" orm:"rel(fk);on_delete(cascade)"`
	StepLen     int       `json:"step_len"`
	OpUser      string    `json:"opr_user"`
	FlowType    string    `json:"flow_type" orm:"default(manual)"` //任务类型 manual手动 crontab定时 depend依赖
	CreatedTime time.Time `json:"created" orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time `json:"updated" orm:"auto_now_add;type(datetime)"`
}

type FlowBatch struct {
	Id          int       `json:"id" orm:"pk;auto"`
	Flow        *Flow     `json:"-" orm:"rel(fk);on_delete(cascade)"`
	Status      int       `json:"status"`
	Step        int       `json:"step"`
	Nodes       string    `json:"nodes" orm:"type(text)"`
	CreatedTime time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time `orm:"auto_now_add;type(datetime)"`
}

// Hold the status of one vm node
type NodeState struct {
	Id          int        `json:"id" orm:"pk;auto"`
	Ip          string     `json:"ip"`
	VmId        string     `json:"vm_id"`
	CorrId      string     `json:"corr_id" orm:"null"` // Correlation ID
	Node        *Node      `json:"-" orm:"rel(fk);null;on_delete(set_null)"`
	Pool        *Pool      `json:"-" orm:"rel(fk)"`
	Flow        *Flow      `json:"-" orm:"rel(fk);on_delete(cascade)"`
	Batch       *FlowBatch `json:"-" orm:"rel(fk);null;"`
	Status      int        `json:"state"`
	Steps       string     `json:"steps" orm:"type(text)"`
	StepNum     int        `json:"step_num"`
	Log         string     `json:"log" orm:"type(text)"`
	LastOp      string     `json:"last_op"`
	CreatedTime time.Time  `json:"created" orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time  `json:"updated" orm:"auto_now_add;type(datetime)"`
}
