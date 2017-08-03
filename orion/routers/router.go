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

package routers

import (
	"github.com/astaxie/beego"

	"weibo.com/opendcp/orion/api"
	"weibo.com/opendcp/orion/controllers"
)

func init() {
	cluster := beego.NewNamespace("/cluster",
		//beego.NSRouter("/", &api.ClusterApi{}, "*:ClusterList"),
		beego.NSRouter("/?:id", &api.ClusterApi{}, "*:ClusterInfo"),
		beego.NSRouter("/list", &api.ClusterApi{}, "*:ClusterList"),
		beego.NSRouter("/create", &api.ClusterApi{}, "*:ClusterAppend"),
		beego.NSRouter("/delete/:id", &api.ClusterApi{}, "*:ClusterDelete"),
		beego.NSRouter("/update/:id", &api.ClusterApi{}, "*:ClusterUpdate"),
		beego.NSRouter("/:id/list_services", &api.ClusterApi{}, "*:ServiceList"),
	)

	service := beego.NewNamespace("/service",
		beego.NSRouter("/:id", &api.ClusterApi{}, "*:ServiceInfo"),
		beego.NSRouter("/create", &api.ClusterApi{}, "*:ServiceAppend"),
		beego.NSRouter("/delete/:id", &api.ClusterApi{}, "*:ServiceDelete"),
		beego.NSRouter("/update/:id", &api.ClusterApi{}, "*:ServiceUpdate"),
		beego.NSRouter("/:id/list_pools", &api.ClusterApi{}, "*:PoolList"),
	)
	pool := beego.NewNamespace("/pool",
		beego.NSRouter("/:id:int", &api.ClusterApi{}, "*:PoolInfo"),

		beego.NSRouter("/create", &api.ClusterApi{}, "*:PoolAppend"),
		beego.NSRouter("/update/:id:int", &api.ClusterApi{}, "*:PoolUpdate"),
		beego.NSRouter("/delete/:id:int", &api.ClusterApi{}, "*:PoolDelete"),

		beego.NSRouter("/:id:int/list_nodes", &api.ClusterApi{}, "*:NodeList"),
		beego.NSRouter("/:id:int/add_nodes", &api.ClusterApi{}, "*:NodeAppend"),
		beego.NSRouter("/:id:int/remove_nodes", &api.ClusterApi{}, "*:NodeDelete"),
		beego.NSRouter("/list", &api.ClusterApi{}, "*:AllPoolList"),

		// search ip
		beego.NSRouter("/search_by_ip/:iplist", &api.ClusterApi{}, "*:SearchPoolByIP"),

		// TODO move this to exec
		beego.NSRouter("/expand/:id:int", &api.ExecApi{}, "*:ExpandPool"),
		beego.NSRouter("/shrink/:id:int", &api.ExecApi{}, "*:ShrinkPool"),
		beego.NSRouter("/deploy/:id:int", &api.ExecApi{}, "*:DeployPool"),
	)

	task := beego.NewNamespace("/task",
		beego.NSNamespace("/impl",
			beego.NSRouter("/create", &api.FlowApi{}, "*:AppendFlowImpl"),
			beego.NSRouter("/delete", &api.FlowApi{}, "*:DeleteFlowImpl"),
			beego.NSRouter("/list", &api.FlowApi{}, "*:ListFlowImpl"),
		),

		beego.NSRouter("/create", &api.FlowApi{}, "*:RunFlow"),
		beego.NSRouter("/start/:id:int", &api.FlowApi{}, "*:StartFlow"),
		beego.NSRouter("/stop/:id:int", &api.FlowApi{}, "*:StopFlow"),
		beego.NSRouter("/pause/:id:int", &api.FlowApi{}, "*:PauseFlow"),
		beego.NSRouter("/flow/:id:int/log", &api.FlowApi{}, "*:GetFlowLogById"),

		//beego.NSRouter("/run", &api.FlowApi{}, "*:RunFlow"),
		beego.NSRouter("/list", &api.FlowApi{}, "*:ListFlow"),
		beego.NSRouter("/:id:int", &api.FlowApi{}, "*:GetFlow"),
		beego.NSRouter("/:id:int/detail", &api.FlowApi{}, "*:GetNodeStates"),

		////获取该pool下的依赖任务列表
		//beego.NSRouter("/expandList/:poolId:int", &api.TaskApi{}, "*:GetExpandList"),
		////获取该pool下的定时任务列表
		//beego.NSRouter("/uploadList/:poolId:int", &api.TaskApi{}, "*:GetUploadList"),
		////增加Task
		//beego.NSRouter("/saveTask", &api.TaskApi{}, "*:SaveTask"),

		beego.NSRouter("/node/:nsid:int/log", &api.FlowApi{}, "*:GetLog"),
	)

	tasktpl := beego.NewNamespace("/task_tpl",
		beego.NSRouter("/create", &api.FlowApi{}, "*:AppendFlowImpl"),
		beego.NSRouter("/list", &api.FlowApi{}, "*:ListFlowImpl"),
		beego.NSRouter("/delete/:id:int", &api.FlowApi{}, "*:DeleteFlowImpl"),
		beego.NSRouter("/:id:int", &api.FlowApi{}, "*:GetFlowImpl"),
		beego.NSRouter("/update/:id:int", &api.FlowApi{}, "*:FlowImplUpdate"),
	)

	taskstep := beego.NewNamespace("/task_step",
		beego.NSRouter("/list", &api.FlowApi{}, "*:ListTaskStep"),
	)

	action := beego.NewNamespace("/action",
		beego.NSNamespace("/impl",
			beego.NSRouter("/create", &api.FlowApi{}, "*:AppendFlowImpl"),
			beego.NSRouter("/delete", &api.FlowApi{}, "*:DeleteFlowImpl"),
			beego.NSRouter("/list", &api.FlowApi{}, "*:ListFlowImpl"),
		),
		beego.NSRouter("/create", &api.RemoteApi{}, "*:ActionAppend"),
		beego.NSRouter("/list", &api.RemoteApi{}, "*:ActionList"),
		beego.NSRouter("/delete/:id:int", &api.RemoteApi{}, "*:ActionDelete"),
		beego.NSRouter("/:id:int", &api.RemoteApi{}, "*:GetAction"),
		beego.NSRouter("/update/:id:int", &api.RemoteApi{}, "*:ActionUpdate"),
	)

	remotestep := beego.NewNamespace("/remote_step",

		beego.NSRouter("/create", &api.RemoteApi{}, "*:RemoteStepAppend"),
		beego.NSRouter("/list", &api.RemoteApi{}, "*:RemoteStepList"),
		beego.NSRouter("/delete/:id:int", &api.RemoteApi{}, "*:RemoteStepDelete"),
		beego.NSRouter("/:id:int", &api.RemoteApi{}, "*:GetRemoteStep"),
		beego.NSRouter("/update/:id:int", &api.RemoteApi{}, "*:RemoteStepUpdate"),
	)

	remoteactionimpl := beego.NewNamespace("/actimpl",

		beego.NSRouter("/create", &api.RemoteApi{}, "*:RemoteActionImplAppend"),
		beego.NSRouter("/:actionId:int", &api.RemoteApi{}, "*:GetRemoteActionImpl"),
		beego.NSRouter("/delete/:id:int", &api.RemoteApi{}, "*:RemoteActionImplDelete"),
		beego.NSRouter("/update/:id:int", &api.RemoteApi{}, "*:RemoteActionImplUpdate"),
		beego.NSRouter("/list", &api.RemoteApi{}, "*:RemoteActionImplList"),
	)

	beego.Router("/", &controllers.MainController{}, "*:Index")

	beego.AddNamespace(cluster)
	beego.AddNamespace(service)
	beego.AddNamespace(pool)
	beego.AddNamespace(task)
	beego.AddNamespace(action)
	beego.AddNamespace(remotestep)
	beego.AddNamespace(tasktpl)
	beego.AddNamespace(taskstep)
	beego.AddNamespace(remoteactionimpl)

}
