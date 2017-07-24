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
	"errors"

	"github.com/astaxie/beego"
)

// Worker keep a queue to hold tasks, and run them one by one.
type Worker struct {
	workQueue chan Job
	key       int
}

// Job represents a task to be run by worker.
type Job func() error

// NewWorker creates a new worker.
func NewWorker(key int) *Worker {
	return &Worker{workQueue: make(chan Job, 5), key: key}
}

func (w *Worker) loop() {
	for {
		job, ok := <-w.workQueue
		if !ok {
			// queue closed
			break
		}
		//修改使任务并行运行
		go w.safeRun(job)
	}
}

func (w *Worker) safeRun(job Job) {
	defer func() {
		if r := recover(); r != nil {
			beego.Info("Recovered from err:", r)
		}
	}()
	key := w.key

	beego.Info("Start running job on pool [", key, "] ... ")
	err := job()
	if err == nil {
		beego.Info("Finish running job on pool [", key, "]")
	} else {
		beego.Error("Error running job on pool [", key, "], err: ", err)
	}
}

// Submit submits new job into queue of this worker, and return error
// if the queue if full.
func (w *Worker) Submit(f Job) error {
	select {
	case w.workQueue <- f:
		beego.Info("Worker[", w.key, "] got new job")
		return nil
	default:
		beego.Error("Too many jobs in queue of worker[", w.key, "]")
		return errors.New("Too many jobs in queue")
	}
}

// Start starts the loop of the worker.
func (w *Worker) Start() {
	go w.loop()
}

// Stop stops the loop of the worker.
func (w *Worker) Stop() {
	close(w.workQueue)
}
