## 集群 
# 创建集群
curl -H "Content-type: application/json" -X POST -d '{"name":"SamplePlatform", "desc":"Sample Platform", "biz": "平台"}' http://$HOST/cluster/create
# 获取集群
curl -H "Content-type: application/json" http://$HOST/cluster/:id
# 获取集群列表
curl -H "Content-type: application/json" http://$HOST/cluster/list
# 更改集群
curl -H "Content-type: application/json" -X POST -d '{"name":"SamplePlatform", "desc":"Sample Platform", "biz": "平台"}' http://$HOST/cluster/update/:id
# 删除集群
curl -H "Content-type: application/json" -X POST http://$HOST/cluster/delete/:id

## 服务 
# 创建服务
curl -H "Content-type: application/json" -X POST -d '{"name":"test","desc":"Test服务","service_type":"Java","docker_image":"sample/test_service:latest","cluster_id":14}' http://$HOST/service/create
# 更改服务
curl -H "Content-type: application/json" -X POST -d '{"name":"test","desc":"Test Service","service_type":"Java","docker_image":"sample/test_service:latest"}' http://$HOST/service/update/:id
# 获取集群下服务列表
curl -H "Content-type: application/json" http://$HOST/cluster/:cluster_id/list_services
# 获取服务
curl -H "Content-type: application/json" http://$HOST/service/:id
# 删除服务
curl -H "Content-type: application/json" http://$HOST/service/delete/:id

## 服务池 
# Create
curl -H "Content-type: application/json" -X POST -d '{"name":"aliyun","desc":"Aliyun","vm_type":1,"sd_id": 1, "tasks":{"expand": 1, "shrink":2,"deploy":3},"service_id":2}' http://$HOST/pool/create
# Update
curl -H "Content-type: application/json" -X POST -d '{"name":"aliyun","desc":"Aliyun","vm_type":1,"sd_id": 1, "tasks":{"expand": 1, "shrink":2,"deploy":3},"service_id":2}' http://$HOST/pool/update/:id
# Get
curl -H "Content-type: application/json" http://$HOST/pool/:id
# List pools of a service
curl -H "Content-type: application/json" http://$HOST/service/:service_id/list_pools
# Delete
curl -H "Content-type: application/json" -X POST http://$HOST/pool/delete/:id

## 远程命令 
# List
curl -H "Content-type: application/json" http://$HOST/action/list
# Get
curl -H "Content-type: application/json" http://$HOST/action/:id
# Create
curl -H "Content-type: application/json" -X POST -d '{"name":"command","desc":"Command", "params":{"time": "integer"}}' http://$HOST/action/create
# Update
curl -H "Content-type: application/json" -X POST -d '{"desc":"Command xx", "params":{"time": "integer"}}' http://$HOST/action/update/:id
# Delete
curl -H "Content-type: application/json" -X POST http://$HOST/action/delete/:id

## 远程命令实现 
# Create
curl -H "Content-type: application/json" -X POST -d '{"action_id":14,"type":"ansible","template":{"action":{"module":"shell","args":"which {{exec}}"}}}' http://$HOST/actimpl/create
# Update
curl -H "Content-type: application/json" -X POST -d '{"action_id":14,"type":"ansible","template":{"action":{"module":"shell","args":"which {{exec}}"}}}' http://$HOST/actimpl/update/:id
# Get
curl -H "Content-type: application/json" http://$HOST/actimpl/:id
# Delete
curl -H "Content-type: application/json" -X POST http://$HOST/actimpl/delete/:id

## 远程步骤 
# Create
curl -H "Content-type: application/json" -X POST -d '{"name":"step","desc":"Step", "actions":["action1", "action2"]}' http://$HOST/remote_step/create
# Update
curl -H "Content-type: application/json" -X POST -d '{"name":"step","desc":"Step", "actions":["action1", "action2", "action3"]}' http://$HOST/remote_step/update/:id
# Get
curl -H "Content-type: application/json" http://$HOST/remote_step/:id
# List
curl -H "Content-type: application/json" http://$HOST/remote_step/list
# Delete
curl -H "Content-type: application/json" -X POST http://$HOST/remote_step/delete/:id

## 任务模板 
# Create
curl -H "Content-type: application/json" -X POST -d '{
            "name":"tpl1","desc":"Template 1",
            "steps":[
                {"name":"sleep","param_values":{"time":1},"retry":{"retry_times":2,"ignore_error":false}},
                {"name":"echo_step","param_values":{"name":"aaa"},"retry":{"retry_times":1,"ignore_error":true}}
             ]
         }' http://$HOST/task_tpl/create
# Get
curl -H "Content-type: application/json" http://$HOST/task_tpl/:id
# List
curl -H "Content-type: application/json" http://$HOST/task_tpl/list
# Delete
curl -H "Content-type: application/json" -X POST http://$HOST/task_tpl/delete/:id
