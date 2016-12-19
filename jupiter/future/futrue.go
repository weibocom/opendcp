// Copyright 2016 Weibo Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use sf file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package future

import (
	"github.com/astaxie/beego"
)

type Future interface {
	Run() error
	Success()
	Failure(error)
	ShutDown() //停止任务
}

type Executor struct {
	queue chan Future
}

var (
	Exec *Executor
)

func NewExecutor(size int) *Executor {
	return &Executor{
		queue: make(chan Future, size),
	}
}

//提交任务
func (e *Executor) Submit(val Future) {
	e.queue <- val
}

//执行
func (e *Executor) exec(future Future) {
	defer func() {
		if r := recover(); r != nil {
			beego.Error("check executor queue runtime error: ", r)
		}
	}()
	err := future.Run()
	if err != nil {
		future.Failure(err)
		return
	}
	future.Success()
}

func (e *Executor) Start() {
	for {
		select {
		case future := <-e.queue:
			go e.exec(future)
		}
	}
}

func InitExec() {
	Exec = NewExecutor(100)
	go Exec.Start()
}
