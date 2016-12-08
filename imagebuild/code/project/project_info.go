/*
 *  Copyright 2009-2016 Weibo, Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
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
