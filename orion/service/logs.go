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

const (
	Debug_level = "Debug"
	Error_level = "Error"
	Info_level  = "Info"
	Warn_level  = "Warn"
)

type LogsService struct {
	BaseService
}

var (
	isPrintBeego = true
)

func (store *LogsService) Debug(Fid, Nid int, v ...interface{}) {
	if isPrintBeego {
		beego.Debug(v...)
	}

	go store.saveToDb(Fid, Nid, Debug_level, v...)
}

func (store *LogsService) Info(Fid, Nid int, v ...interface{}) {
	if isPrintBeego {
		beego.Info(v...)
	}

	go store.saveToDb(Fid, Nid, Info_level, v...)
}

func (store *LogsService) Warn(Fid, Nid int, v ...interface{}) {
	if isPrintBeego {
		beego.Warn(v...)
	}

	go store.saveToDb(Fid, Nid, Warn_level, v...)
}

func (store *LogsService) Error(Fid, Nid int, v ...interface{}) {
	if isPrintBeego {
		beego.Error(v...)
	}

	go store.saveToDb(Fid, Nid, Error_level, v...)
}

func (store *LogsService) saveToDb(Fid, Nid int, Level string, v ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			beego.Info("saveToDb is error !", r)
		}
	}()

	msg := strings.Repeat(" %v", len(v))
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}

	logs := models.NewLogsInit(Fid, Nid, Level, msg)
	o := orm.NewOrm()
	id, err := o.Insert(logs)

	if err != nil {
		beego.Error("[Store] LogMessage fail!", err)
	} else {
		fmt.Println(id)
	}
}
