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

type Organization struct {
	Id         int64 `orm:"pk;auto"`
	Name       string
	Desc       string
	CreateTime time.Time `orm:"auto_now_add;type(datetime)"`
}

type InstanceOrganization struct {
	Id           int           `orm:"pk;auto"`
	Organization *Organization `orm:"rel(fk)"`
	Instance     *Instance   `orm:"rel(fk)"`
}
