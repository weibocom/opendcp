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
	"time"
)

type Logs struct {
	Id            int    `json:"id" orm:"pk;auto"`
	Fid           int    `json:"fid"`
	CorrelationId string `json:"correlation_id"`
	Message       string `json:"message"` //日志信息
	Ctime         int    `json:"ctime"`
}

func NewLogsInit(Fid int, correlationId string, Message string) (result *Logs) {
	result = &Logs{}
	result.Fid = Fid
	result.CorrelationId = correlationId
	result.Message = Message
	result.Ctime = int(time.Now().Unix())
	return result
}
