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
