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

type RemoteStep struct {
	Id      int    `json:"id" orm:"pk;auto"`
	Name    string `json:"name" orm:"size(50);unique"`
	Desc    string `json:"desc" orm:"size(255);null"`
	Actions string `json:"actions" orm:"type(text)"`
}

type RemoteAction struct {
	Id     int    `json:"id" orm:"pk;auto"`
	Name   string `json:"name" orm:"size(50);unique"`
	Desc   string `json:"desc" orm:"size(255);null"`
	Params string `json:"params" orm:"type(text)"`
}

type RemoteActionImpl struct {
	Id       int    `json:"id" orm:"pk;auto"`
	Type     string `json:"type" orm:"size(50)"`
	Template string `json:"template" orm:"type(text)"`
	ActionId int    `json:"action_id"`
}
