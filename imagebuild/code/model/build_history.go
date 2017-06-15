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



package model

import (
	"time"
)

/**
build 历史记录
 */
type BuildHistory struct {
	project  string
	operator string
	time     time.Time
	state    int
	logs     string
}

func (h *BuildHistory) State() int {
	return h.state
}

func (h *BuildHistory) Logs() string {
	return h.logs
}

func GetBuildHistory(project string, operator string, time time.Time, state int, logs string) *BuildHistory {
	return &BuildHistory{
		project:  project,
		operator: operator,
		time:     time,
		state:    state,
		logs:     logs,
	}
}
