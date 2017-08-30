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
	"strconv"

	h "weibo.com/opendcp/orion/helper"
)

type ExecApi struct {
	baseAPI
}

func (e *ExecApi) URLMapping() {
	e.Mapping("Expand", e.ExpandPool)
	e.Mapping("Shrink", e.ShrinkPool)
	e.Mapping("Deploy", e.DeployPool)
}

func (c *ExecApi) ExpandPool() {

	opUser := c.Ctx.Input.Header("Authorization")
	id := c.Ctx.Input.Param(":id")
	idInt, _ := strconv.Atoi(id)

	req := struct {
		Num int `json:"num"`
	}{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	num := req.Num
	if num < 1 || num > 500 {
		c.ReturnFailed("Bad num: "+strconv.Itoa(num), 400)
		return
	}

	err = h.Expand(idInt, num, opUser)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnSuccess(nil)
}

func (c *ExecApi) ShrinkPool() {

	opUser := c.Ctx.Input.Header("Authorization")
	id := c.Ctx.Input.Param(":id")
	poolId, _ := strconv.Atoi(id)

	req := struct {
		Nodes []string `json:"nodes"`
	}{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	nodes := req.Nodes
	err = h.Shrink(poolId, nodes, opUser)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnSuccess(nil)
}

func (c *ExecApi) DeployPool() {

	opUser := c.Ctx.Input.Header("Authorization")
	id := c.Ctx.Input.Param(":id")
	poolId, _ := strconv.Atoi(id)

	req := struct {
		MaxNum int    `json:"max_num"`
		Tag    string `json:"tag"`
	}{}

	err := c.Body2Json(&req)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	err = h.Deploy(poolId, req.Tag, req.MaxNum, opUser)
	if err != nil {
		c.ReturnFailed(err.Error(), 400)
		return
	}

	c.ReturnSuccess(nil)
}
