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

		//beego.NSRouter("/run", &api.FlowApi{}, "*:RunFlow"),
		beego.NSRouter("/list", &api.FlowApi{}, "*:ListFlow"),
		beego.NSRouter("/:id:int", &api.FlowApi{}, "*:GetFlow"),
		beego.NSRouter("/:id:int/detail", &api.FlowApi{}, "*:GetNodeStates"),

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
