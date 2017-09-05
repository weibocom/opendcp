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

package executor

import (
	"github.com/astaxie/beego"
	"time"
	"weibo.com/opendcp/orion/models"
)

const (
	runNodeInterval = time.Microsecond //1us
)

// Worker keep a queue to hold tasks, and run them one by one.
type QueueNode struct {
	workNodeQueue chan ToRunNodeState
}

type ToRunNodeState struct {
	resultChannel chan *models.NodeState
	flow          *models.Flow
	steps         []*models.ActionImpl
	stepOptions   []*models.StepOption
	nodeState     *models.NodeState
}

// NewWorker creates a new worker.
func NewQueueNode() *QueueNode {
	return &QueueNode{workNodeQueue: make(chan ToRunNodeState, 500)}
}

func (q *QueueNode) loop() {
	for {
		runNode, ok := <-q.workNodeQueue
		if !ok {
			beego.Error("workNodeQueue was closed!")
			// queue closed
			break
		}
		go q.safeRun(runNode)
		time.Sleep(runNodeInterval)
	}
}

func (q *QueueNode) safeRun(runNode ToRunNodeState) {
	defer func() {
		if r := recover(); r != nil {
			beego.Info("Recovered from err:", r)
		}
	}()

	beego.Info("Start running nodeState on flow [", runNode.flow.Id, "]")
	err := Executor.RunNodeState(runNode.flow, runNode.nodeState, runNode.steps,
		runNode.stepOptions, runNode.resultChannel)
	if err == nil {
		beego.Info("Finish running nodeState on flow [", runNode.flow.Id, "]")
	} else {
		beego.Error("Error running NodeState on flow [", runNode.flow.Id, "], err: ", err)
	}
}

// Submit submits new job into queue of this worker, and return error
func (q *QueueNode) Submit(runNode ToRunNodeState) {
	select {
	case q.workNodeQueue <- runNode:
		beego.Info("WorkerNodeState[", runNode.nodeState.Id, "] got new job")
	}
}

// Start starts the loop of the worker.
func (q *QueueNode) Start() {
	go q.loop()
}

// Stop stops the loop of the worker.
func (q *QueueNode) Stop() {
	close(q.workNodeQueue)
}
