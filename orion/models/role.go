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

type RoleResource struct {
	Id                int       `json:"id" orm:"pk;auto"`
	Name              string    `json:"name" orm:"size(50);unique"`
	Desc              string    `json:"desc" orm:"size(200)"`
	ResourceType      string    `json:"resource_type" orm:"size(200)"`
	ResourceContent   string    `json:"resource_content" orm:"type(text)"`
	User              string    `json:"user" orm:"size(200)"`
	State             int16     `json:"state" orm:"size(1)"`
	CreateTime        time.Time `json:"create_time" orm:"type(datetime);auto_now_add"`
	UpdateTime        time.Time `json:"update_time" orm:"type(datetime);auto_now"`
	Type              int16     `json:"type" orm:"size(3)"`
	Hidden            int16     `json:"hidden" orm:"size(3)"`
	TemplateFilePath  string    `json:"template_file_path" orm:"size(512)"`
	TemplateFilePerm  string    `json:"template_file_perm" orm:"size(32)"`
	TemplateFileOwner string    `json:"template_file_owner" orm:"size(32)"`
	AssociateRole     string    `json:"associate_role" orm:"size(512)"`
}

type Role struct {
	Id           int       `json:"id" orm:"pk;auto"`
	Name         string    `json:"name" orm:"size(50);unique"`
	Desc         string    `json:"desc" orm:"size(200)"`
	RoleFilePath string    `json:"role_file_path" orm:"size(512)"`
	Files        string    `json:"files" orm:"size(512)"`
	Handles      string    `json:"handles" orm:"size(512)"`
	Meta         string    `json:"meta" orm:"size(512)"`
	Tasks        string    `json:"tasks" orm:"size(512)"`
	Templates    string    `json:"templates" orm:"size(512)"`
	Vars         string    `json:"vars" orm:"size(512)"`
	User         string    `json:"user" orm:"size(200)"`
	State        int       `json:"state" orm:"size(1);default(0)"`
	CreateTime   time.Time `json:"create_time" orm:"type(datetime);auto_now_add"`
	UpdateTime   time.Time `json:"update_time" orm:"type(datetime);auto_now"`
}
