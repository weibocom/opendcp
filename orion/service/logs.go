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


package service

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	"weibo.com/opendcp/orion/models"
)

type LogsService struct {
	BaseService
}

var (
	isPrintBeego = true
)


func (store *LogsService) Debug(Fid int,BatchId int, correlationId string, Message string, v ...interface{}) {
	if isPrintBeego {
		beego.Debug(Message,v)
	}

	go store.saveToDb(Fid,BatchId,correlationId,Message,v ...)
}


func (store *LogsService) Info(Fid int,BatchId int, correlationId string, Message string, v ...interface{}) {
	if isPrintBeego {
		beego.Info(Message,v)
	}

	go store.saveToDb(Fid,BatchId,correlationId,Message,v ...)
}

func (store *LogsService) Warn(Fid int,BatchId int, correlationId string, Message string, v ...interface{}) {
	if isPrintBeego {
		beego.Warn(Message,v)
	}

	go store.saveToDb(Fid,BatchId,correlationId,Message,v ...)
}

func (store *LogsService) Error(Fid int,BatchId int, correlationId string, Message string, v ...interface{}) {
	if isPrintBeego {
		beego.Error(Message,v)
	}

	go store.saveToDb(Fid,BatchId,correlationId,Message,v ...)
}


func (store *LogsService) saveToDb(Fid int,BatchId int, correlationId string, Message string,v ...interface{}){
	defer func() {
		if r := recover(); r != nil {
			beego.Info("saveToDb is error !", r)
		}
	}()

	if len(v) > 0 {
		msg := strings.Repeat(" %v", len(v))
		Message += fmt.Sprintf(msg, v...)
	}

	logs := models.NewLogsInit(Fid,BatchId,correlationId,Message)
	orm := orm.NewOrm()
	id, err := orm.Insert(logs)

	if err != nil {
		beego.Error("[Store] LogMessage fail!", err)
	} else {
		fmt.Println(id)
	}
}