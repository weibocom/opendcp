/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 */
package models

import "time"

const (
	STATUS_INIT = iota
	STATUS_RUNNING
	STATUS_SUCCESS
	STATUS_FAILED
	STATUS_STOPPED
)

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
	Id          int       `json:"id" orm:"pk;auto"`
	Ip          string    `json:"ip"`
	VmId        string    `json:"vm_id"`
	CorrId      string    `json:"corr_id" orm:"null"` // Correlation ID
	Node        *Node     `json:"-" orm:"rel(fk);null;on_delete(set_null)"`
	Pool        *Pool     `json:"-" orm:"rel(fk)"`
	Flow        *Flow     `json:"-" orm:"rel(fk);on_delete(cascade)"`
	Batch	    *FlowBatch`json:"-" orm:"rel(fk);null;"`
	Status      int       `json:"state"`
	Steps       string    `json:"steps" orm:"type(text)"`
	Log         string    `orm:"type(text)"`
	CreatedTime time.Time `json:"created" orm:"auto_now_add;type(datetime)"`
	UpdatedTime time.Time `json:"updated" orm:"auto_now_add;type(datetime)"`
}
