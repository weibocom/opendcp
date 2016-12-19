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



package project

/**
项目信息
 */
type ProjectInfo struct {
	Creator              string `json:"creator"`
	Name                 string `json:"name"`
	CreateTime           string `json:"createTime"`
	LastModifyTime       string `json:"lastModifyTime"`
	LastModifyOperator   string `json:"lastModifyOperator"`
	Cluster              string `json:"Cluster"`
	DefineDockerFileType string `json:"DefineDockerFileType"`
}

func (p *ProjectInfo) Info() ProjectInfo {
	return ProjectInfo{
		Creator:              p.Creator,
		Name:                 p.Name,
		CreateTime:           p.CreateTime,
		LastModifyTime:       p.LastModifyTime,
		LastModifyOperator:   p.LastModifyOperator,
		Cluster:              p.Cluster,
		DefineDockerFileType: p.DefineDockerFileType}
}

func BuildEmptyProjectInfo() ProjectInfo {
	return ProjectInfo{
		Creator:              "",
		Name:                 "",
		CreateTime:           "",
		LastModifyTime:       "",
		LastModifyOperator:   "",
		Cluster:              "",
		DefineDockerFileType: ""}
}

// 排序使用
type ProjectInfoList []ProjectInfo

func (c ProjectInfoList) Len() int {
	return len(c)
}

func (c ProjectInfoList) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ProjectInfoList) Less(i, j int) bool {
	return c[i].Name < c[j].Name
}
