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


package api

import (
	//"fmt"
	"strconv"

	h "weibo.com/opendcp/orion/helper"
	//. "weibo.com/opendcp/orion/models"
	//s "weibo.com/opendcp/orion/service"
	//u "weibo.com/opendcp/orion/utils"
)

const ()

type ExecApi struct {
	baseAPI
}

func (e *ExecApi) URLMapping() {
	e.Mapping("Expand", e.ExpandPool)
	e.Mapping("Shrink", e.ShrinkPool)
	e.Mapping("Deploy", e.DeployPool)
}

func (c *ExecApi) ExpandPool() {
	biz := c.Ctx.Input.Header("X-Biz-ID")
	biz_id,err := strconv.Atoi(biz)
	if err !=nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	opUser := c.Ctx.Input.Header("Authorization")
	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)


	req := struct {
		Num int `json:"num"`
	}{}

	err = c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	num := req.Num
	if num < 1 || num > 100 {
		c.ReturnFailed("Bad num: "+strconv.Itoa(num), 400)
		return
	}

	err = h.Expand(idInt, num, opUser, biz_id)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnSuccess(nil)
}

func (c *ExecApi) ShrinkPool() {
	biz := c.Ctx.Input.Header("X-Biz-ID")
	biz_id,err := strconv.Atoi(biz)
	if err !=nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	opUser := c.Ctx.Input.Header("Authorization")
	id := c.Ctx.Input.Param(":id")
	poolId, _ := strconv.Atoi(id)

	req := struct {
		Nodes []string `json:"nodes"`
	}{}

	err = c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	nodes := req.Nodes
	err = h.Shrink(poolId, nodes, opUser, biz_id)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnSuccess(nil)
}

func (c *ExecApi) DeployPool() {
	biz := c.Ctx.Input.Header("X-Biz-ID")
	biz_id,err := strconv.Atoi(biz)
	if err !=nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}
	opUser := c.Ctx.Input.Header("Authorization")
	id := c.Ctx.Input.Param(":id")
	poolId, _ := strconv.Atoi(id)

	req := struct {
		MaxNum int    `json:"max_num"`
		Tag    string `json:"tag"`
	}{}

	err = c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	err = h.Deploy(poolId, req.Tag, req.MaxNum, opUser, biz_id)
	//err = h.Deploy(poolId, req.MaxNum)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnSuccess(nil)
}
