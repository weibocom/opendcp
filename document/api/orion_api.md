# 服务编排API

## 集群接口

### 1.集群列表接口

获取用户创建的所有服务集群列表

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/cluster/list?page=1&page_size=10 |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                                  |
| --------- | ---- | ---- | ----------------------------------- |
| page      | int  | 否    | 当前页码数，即本次API调用是获得结果的第几页，从1开始计数，默认为1 |
| page_size | int  | 否    | 当前页包含的结果数，默认结果数为10                  |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | string | "sucess" | 接口返回信息 |
| data        | array | [{},{}]  | 数据结果   |
| page        | int   | 1        | 当前页    |
| page_size   | int   | 10       | 当前页大小  |
| query_count | int   | 2        | 结果数    |

data 参数中object对象说明

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| id   | int    | 0    | 集群id |
| name | string | 0    | 集群名称 |
| desc | string | 0    | 集群描述 |
| biz  | string | 0    | 产品线  |

#### 请求示例

```php
curl -X GET "http://HOST/cluster/list?page=1&page_size=10"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
            "id": 2,
            "name": "openstack_cluster",
            "desc": "虚拟化集群",
            "biz": "1"
        }, {
            "id": 2,
            "name": "openstack_cluster",
            "desc": "虚拟化集群",
            "biz": "1"
        },
    ],
    "page": 1,
    "page_size": 10,
    "query_count": 1
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": [],
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 执行成功 |                           |
| 400  | 执行失败 | 查询数据库失败，详细信息请查看返回结果的msg字段 |

----------

### 2.创建集群接口

创建服务集群

#### 请求地址

| POST方法                     |
| -------------------------- |
| http://HOST/cluster/create |

#### 请求参数

| 名称   | 类型     | 是否必须 | 描述       |
| ---- | ------ | ---- | -------- |
| name | string | 是    | 要增加的集群名称 |
| desc | string | 是    | 集群描述     |
| biz  | string | 是    | 产品线，通常为1 |

#### 返回参数

| 名称   | 类型    | 示例值      | 描述      |
| ---- | ----- | -------- | ------- |
| code | int   | 0        | 返回码     |
| msg  | string| "sucess" | 接口返回信息  |
| data | int   |          | 添加的集群id |

#### 请求示例

```php
curl -X POST 'http://HOST/cluster/create' \
-H "Content-type: application/json" \
-d '{"name":"SamplePlatform", "desc":"Sample Platform", "biz": "平台"}' 
```

#### 响应示例

正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": 1,
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 400  | 执行失败 | 添加失败，详细信息请查看返回结果的msg字段 |

----------

### 3.修改集群接口

修改集群

#### 请求地址

| POST方法                         |
| ------------------------------ |
| http://HOST/cluster/update/:id |

#### 请求参数
| 名称   | 类型     | 是否必须 | 描述      |
| ---- | ------ | ---- | ------- |
| name | string | 是    | 集群修改名称  |
| desc | string | 是    | 集群为修改描述 |
| biz  | string | 是    | 产品线     |

#### 返回参数
| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | string | ""       |        |

#### 请求示例
```php
curl -X POST 'http://HOST/cluster/update/1' \
-H "Content-type: application/json" \
 -d '{"name":"SamplePlatform", "desc":"Sample Platform", "biz": "平台"}' 
```
#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": "",
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": null,
}
```
#### 返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 400  | 执行失败 | 修改失败，详细信息请查看返回结果的msg字段 |

----------

### 4.删除集群接口

删除集群

#### 请求地址

| POST方法                         |
| ------------------------------ |
| http://HOST/cluster/delete/:id |

#### 请求参数

| 名称   | 类型   | 示例值  | 描述       |
| ---- | ---- | ---- | -------- |
| id   | int  | 0    | 要删除的集群id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | null     |        |

#### 请求示例

```php
curl -X POST 'http://$HOST/cluster/delete/2' \
-H "Content-type: application/json" 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": "",
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": null,
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 400  | 执行失败 | 删除失败，详细信息请查看返回结果的msg字段 |

----------

### 5.获取集群详情

获取集群详情

#### 请求地址

| POST方法                  |
| ----------------------- |
| http://HOST/cluster/:id |

#### 请求参数
| 名称   | 类型   | 是否必须 | 描述       |
| ---- | ---- | ---- | -------- |
| id   | int  | 是    | 要增加的集群名称 |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data参数

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| id   | int    | 0    | 服务Id |
| name | string |      | 服务名称 |
| desc | string |      | 服务描述 |
| biz  | string |      | 产品线  |

#### 请求示例

```php
curl 
-H "Content-type: application/json" \
-X POST 'http://$HOST/cluster/2' \
```

#### 响应示例

正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": {
      "id": 1,
      "name":"default cluster",
      "desc": "默认集群",
      "biz":"1"
    },
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "id is error...",
    "data": {},
}
```

```json
{
    "code": 404,
    "msg": "cluster in db not found",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述          |
| ---- | ---- | ----------- |
| 0    | 执行成功 |             |
| 400  | 执行失败 | 传入的id有误     |
| 404  | 执行失败 | 数据库中查询不到该集群 |

----------

### 6.获取集群中服务列表接口

获取该集群中包含的所有服务

#### 请求地址

| POST方法                                |
| ------------------------------------- |
| http://HOST/cluster/:id/list_services |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述            |
| --------- | ---- | ---- | ------------- |
| id        | int  | 是    | 要增加的集群名称      |
| page      | int  | 否    | 当前页，从1开始，默认为1 |
| page_size | int  | 否    | 当前页大小，默认为10   |

#### 返回参数

| 名称          | 类型     | 示例值      | 描述     |
| ----------- | ------ | -------- | ------ |
| code        | int    | 0        | 返回码    |
| msg         | string | "sucess" | 接口返回信息 |
| data        | object | {}       | 返回结果   |
| page        | int    |          | 当前页    |
| page_size   | int    |          | 当前页大小  |
| query_count | int    |          | 结果个数   |

data参数

| 名称           | 类型     | 示例值  | 描述   |
| ------------ | ------ | ---- | ---- |
| id           | int    | 0    | 服务Id |
| name         | string |      | 服务名称 |
| desc         | string |      | 集群描述 |
| docker_image | string |      | 镜像地址 |
| cluster_id   | int    |      | 集群id |

#### 请求示例

```php
curl 
-H "Content-type: application/json" \
-X POST 'http://HOST/cluster/:id/list_services ' \
-d "page=1" \
-d "page_size=10"
```

#### 响应示例

正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": [{
      "id": 1,
      "name":"service 1",
      "desc": "服务1",
      "docker_image": "registry.cn-beijing.aliyuncs.com/opendcp/java-web:latest",
      "cluster_id":1
    },{
      "id": 1,
      "name":"sevice 2",
      "desc": "服务2",
      "docker_image": "-",
      "cluster_id":1
    }
    ],
    "page":1,
    "page_size":10,
    "query_count":2
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "error...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述             |
| ---- | ---- | -------------- |
| 0    | 执行成功 |                |
| 400  | 执行失败 | 集群id有误，查询数据库失败 |

----------

## 服务接口

### 1.创建服务接口

服务创建

#### 请求地址

| POST方法                     |
| -------------------------- |
| http://HOST/service/create |

#### 请求参数

| 名称           | 类型     | 是否必须 | 描述       |
| ------------ | ------ | ---- | -------- |
| name         | string | 是    | 服务名称     |
| desc         | string | 是    | 服务描述     |
| service_type | string | 是    | 服务类型     |
| docker_image | string | 是    | 服务镜像     |
| cluster_id   | string | 是    | 服务对应集群id |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data 参数中object对象说明

| 名称   | 类型   | 示例值  | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 1    | 服务id |


#### 请求示例

```php
curl -X GET "http://HOST/service/create"  \
-H "Content-type: application/json" \
-d '{"name":"web_service","desc":"web服务","service_type":"Java","docker_image":"myregistry/myweb:1.0","cluster_id":1}  
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": {
            "id": 1
        }
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "请求参数解析错误",
    "data": []
}
```

#### 返回码解释

| 返回码  | 状态   | 描述       |
| ---- | ---- | -------- |
| 0    | 执行成功 |          |
| 400  | 执行失败 | 请求参数解析错误 |

----------

### 2.删除服务接口

服务删除

#### 请求地址

| DELETE方法                       |
| ------------------------------ |
| http://HOST/service/delete/:id |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 是    | 服务id |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |


#### 请求示例

```php
curl -X DELETE "http://HOST/service/delete/1"  \
-H "Content-type: application/json"   
```

#### 响应示例

正常返回结果：
```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```
异常返回结果：

```json
{
    "code": 400,
    "msg": "id is error",
    "data": []
}
```

```json
{
    "code": 404,
    "msg": "数据库操作失败",
    "data": []
}
```

#### 返回码解释

| 返回码  | 状态   | 描述       |
| ---- | ---- | -------- |
| 0    | 执行成功 |          |
| 400  | 执行失败 | 请求参数解析错误 |
| 404  | 执行失败 | 数据库操作失败  |

----------

### 3.获取服务信息接口

根据服务id获取服务具体信息

#### 请求地址

| GET方法                    |
| ------------------------ |
| http://HOST/service/{id} |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述    |
| ---- | ---- | ---- | ----- |
| id   | int  | 是    | 服务池id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data 参数中object对象说明

| 名称           | 类型     | 示例值           | 描述       |
| ------------ | ------ | ------------- | -------- |
| id           | int    | 1             | 服务id     |
| name         | string | web_service   | 服务名称     |
| desc         | string | web服务         | 服务描述     |
| service_type | string | Java          | 服务类型     |
| docker_image | string | XXX/myweb:1.0 | 服务镜像     |
| cluster_id   | int    | 1             | 服务对应集群id |


#### 请求示例

```php
curl -X GET "http://HOST/service/1"  \
-H "Content-type: application/json" 
```
#### 响应示例

正常返回结果：
```json
{
    "code": 0,
    "msg": "success",
    "data": {
            "id": 1,
            "name": "web_service",
            "desc": "web服务",
            "service_type": "Java",
            "docker_image": "myregistry/myweb:1.0",
            "cluster_id": 1
        }
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "id is error !",
    "data": []
}
```
```json
{
    "code": 404,
    "msg": "db error",
    "data": []
}
```

#### 返回码解释

| 返回码  | 状态   | 描述       |
| ---- | ---- | -------- |
| 0    | 执行成功 |          |
| 404  | 执行失败 | 数据库记录不存在 |
| 400  | 执行失败 | id校验错误   |

----------

### 3.获取服务中服务池列表接口

获取该服务中包含的所有服务池

#### 请求地址

| POST方法                             |
| ---------------------------------- |
| http://HOST/service/:id/list_pools |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述            |
| --------- | ---- | ---- | ------------- |
| id        | int  | 是    | 服务id          |
| page      | int  | 否    | 当前页，从1开始，默认为1 |
| page_size | int  | 否    | 当前页大小，默认为10   |

#### 返回参数

| 名称          | 类型     | 示例值      | 描述     |
| ----------- | ------ | -------- | ------ |
| code        | int    | 0        | 返回码    |
| msg         | string | "sucess" | 接口返回信息 |
| data        | object | {}       | 返回结果   |
| page        | int    |          | 当前页    |
| page_size   | int    |          | 当前页大小  |
| query_count | int    |          | 结果个数   |

data参数

| 名称         | 类型     | 示例值  | 描述       |
| ---------- | ------ | ---- | -------- |
| id         | int    | 0    | 服务池Id    |
| name       | string |      | 服务池名称    |
| desc       | string |      | 服务池描述    |
| vm_type    | string |      | 虚拟机模板id  |
| sd_id      | int    |      | 服务发现id   |
| service_id | int    |      | 服务id     |
| node_count | int    |      | 服务池中节点数目 |

#### 请求示例

```php
curl 
-H "Content-type: application/json" \
-X POST 'http://HOST/service/:id/list_pools' \
-d "page=1" \
-d "page_size=10"
```
#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
      "id": 1,
      "name":"web_pool",
      "desc":"web服务池",
      "vm_type":1,
      "sd_id":1,
      "service_id":1,
      "node_count":2,
    }],
    "page":1,
    "page_size":10,
    "query_count":1
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "error...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述             |
| ---- | ---- | -------------- |
| 0    | 执行成功 |                |
| 400  | 执行失败 | 集群id有误，查询数据库失败 |

----------

## 服务池接口

### 1.创建服务池接口

服务池创建

#### 请求地址

| POST方法                  |
| ----------------------- |
| http://HOST/pool/create |

#### 请求参数

| 名称         | 类型     | 是否必须 | 描述        |
| ---------- | ------ | ---- | --------- |
| name       | string | 是    | 服务池名称     |
| desc       | string | 是    | 服务池描述     |
| vm_type    | int    | 是    | 虚拟机模板id   |
| sd_id      | int    | 是    | 服务发现id    |
| service_id | int    | 是    | 服务池所在服务id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data 参数中object对象说明

| 名称   | 类型   | 示例值  | 描述    |
| ---- | ---- | ---- | ----- |
| id   | int  | 1    | 服务池id |


#### 请求示例

```php
curl -X GET "http://HOST/service/create"  \
-H "Content-type: application/json" \
-d '{"name":"web_pool","desc":"web服务池","vm_type":1,"sd_id":1,"service_id":1}  
```

#### 响应示例

正常返回结果：
```json
{
    "code": 0,
    "msg": "success",
    "data": {
            "id": 1
        }
}
```
异常返回结果：
```json
{
    "code": 400,
    "msg": "请求参数解析错误",
    "data": []
}
```

#### 返回码解释

| 返回码  | 状态   | 描述       |
| ---- | ---- | -------- |
| 0    | 执行成功 |          |
| 400  | 执行失败 | 请求参数解析错误 |

----------

### 2.删除服务池接口

服务池删除

#### 请求地址

| DELETE方法                    |
| --------------------------- |
| http://HOST/pool/delete/:id |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述    |
| ---- | ---- | ---- | ----- |
| id   | int  | 是    | 服务池id |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

#### 请求示例

```php
curl -X DELETE "http://HOST/pool/delete/1"  \
-H "Content-type: application/json"   
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "data not found",
    "data": []
}
```
```json
{
    "code": 400,
    "msg": "node exists in this pool",
    "data": []
}
```
```json
{
    "code": 400,
    "msg": "fail to delete pool",
    "data": []
}
```

#### 返回码解释

| 返回码  | 状态   | 描述   |
| ---- | ---- | ---- |
| 0    | 执行成功 |      |
| 400  | 执行失败 |      |

----------

### 3.修改服务池接口

服务修改

#### 请求地址

| POST方法                      |
| --------------------------- |
| http://HOST/pool/update/:id |

#### 请求参数

| 名称      | 类型     | 是否必须 | 描述      |
| ------- | ------ | ---- | ------- |
| id      | int    | 否    | 服务池id   |
| desc    | string | 否    | 服务池描述   |
| vm_type | int    | 否    | 虚拟机模板id |
| sd_id   | int    | 否    | 服务发现id  |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |


#### 请求示例

```php
curl -X DELETE "http://HOST/pool/update/1"  \
-H "Content-type: application/json" \
-d '{"name":"web_pool","desc":"web服务池","vm_type":1,"sd_id":1}  
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "参数解析错误",
    "data": []
}
```

```json
{
    "code": 404,
    "msg": "数据库操作失败",
    "data": []
}
```

#### 返回码解释

| 返回码  | 状态   | 描述       |
| ---- | ---- | -------- |
| 0    | 执行成功 |          |
| 400  | 执行失败 | 请求参数解析错误 |
| 404  | 执行失败 | 数据库操作失败  |

----------

### 4.获取服务池信息接口 

根据服务池id获取服务池具体信息

#### 请求地址

| GET方法                 |
| --------------------- |
| http://HOST/pool/{id} |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述    |
| ---- | ---- | ---- | ----- |
| id   | int  | 是    | 服务池id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data 参数中object对象说明

| 名称         | 类型     | 示例值      | 描述        |
| ---------- | ------ | -------- | --------- |
| name       | string | web_pool | 服务池名称     |
| desc       | string | web服务池   | 服务池描述     |
| vm_type    | int    | 1        | 虚拟机模板id   |
| sd_id      | int    | 1        | 服务发现id    |
| service_id | int    | 1        | 服务池所在服务id |

#### 请求示例

```php
curl -X GET "http://HOST/pool/1"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": {
            "id": 1,
            "name": "web_pool",
            "desc": "web服务池",
            "vm_type": 1,
            "sd_id": 1,
            "service_id": 1
        }
}
```

异常返回结果：

```json
{
    "code": 404,
    "msg": "db error",
    "data": []
}
```

#### 返回码解释

| 返回码  | 状态   | 描述       |
| ---- | ---- | -------- |
| 0    | 执行成功 |          |
| 404  | 执行失败 | 数据库记录不存在 |

-----

### 5.获取服务池中所有节点接口

获取该服务池中所有节点信息

#### 请求地址

| POST方法                          |
| ------------------------------- |
| http://HOST/pool/:id/list_nodes |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述            |
| --------- | ---- | ---- | ------------- |
| id        | int  | 是    | 服务id          |
| page      | int  | 否    | 当前页，从1开始，默认为1 |
| page_size | int  | 否    | 当前页大小，默认为10   |

#### 返回参数

| 名称          | 类型     | 示例值      | 描述     |
| ----------- | ------ | -------- | ------ |
| code        | int    | 0        | 返回码    |
| msg         | string | "sucess" | 接口返回信息 |
| data        | object | {}       | 返回结果   |
| page        | int    |          | 当前页    |
| page_size   | int    |          | 当前页大小  |
| query_count | int    |          | 结果个数   |

data参数

| 名称        | 类型     | 示例值  | 描述                         |
| --------- | ------ | ---- | -------------------------- |
| id        | int    | 0    | 节点Id                       |
| ip        | string |      | 节点ip地址                     |
| vm_id     | string |      | 节点实例编号                     |
| status    | int    |      | 节点状态（1 成功/3 失败）            |
| node_type | int    |      | 节点类型（手动 manual/定时 crontab） |

#### 请求示例

```php
curl 
-H "Content-type: application/json" \
-X POST 'http://HOST/pool/:id/list_nodes' \
-d "page=1" \
-d "page_size=10"
```
#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
      "id": 1,
      "ip":"10.77.9.75",
      "vm_id":"i-2zeen6mal4s9qvpqb4iq",
      "status":1,
      "vnode_type":"manual"
    }],
    "page":1,
    "page_size":10,
    "query_count":1
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "error...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述      |
| ---- | ---- | ------- |
| 0    | 执行成功 |         |
| 400  | 执行失败 | 查询数据库失败 |

-----

### 6.服务池录入节点接口

服务池录入节点

#### 请求地址

| POST方法                         |
| ------------------------------ |
| http://HOST/pool/:id/add_nodes |

#### 请求参数

| 名称    | 类型       | 是否必须 | 描述     |
| ----- | -------- | ---- | ------ |
| id    | int      | 是    | 服务池id  |
| nodes | []string | 是    | 节点ip列表 |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data 参数中object对象说明

| 名称   | 类型    | 示例值   | 描述          |
| ---- | ----- | ----- | ----------- |
|      | []int | 1,2,3 | 插入node的id数组 |

#### 请求示例

```php
curl -X GET "http://HOST/pool/1/add_nodes"  \
-H "Content-type: application/json" \
-d '{"nodes":"["10.85.41.166","10.85.41.167","10.85.41.168"]"}  
```

#### 响应示例

正常返回结果：
```json
{
    "code": 0,
    "msg": "success",
    "data": [1,2,3]
}
```
异常返回结果：
```json
{
    "code": 404,
    "msg": "ip exists already",
    "data": []
}
```
```json
{
    "code": 404,
    "msg": "pool_id is not vaild",
    "data": []
}
```
```json
{
    "code": 404,
    "msg": "ip is empty",
    "data": []
}
```

#### 返回码解释

| 返回码  | 状态   | 描述       |
| ---- | ---- | -------- |
| 0    | 执行成功 |          |
| 404  | 执行失败 | 请求参数解析错误 |

------

### 7.服务池删除节点接口

服务池删除节点

#### 请求地址

| DELETE方法                          |
| --------------------------------- |
| http://HOST/pool/:id/remove_nodes |

#### 请求参数

| 名称    | 类型       | 是否必须 | 描述     |
| ----- | -------- | ---- | ------ |
| id    | int      | 是    | 服务池id  |
| nodes | []string | 是    | 节点ip列表 |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

#### 请求示例

```php
curl -X GET "http://HOST/pool/1/add_nodes"  \
-H "Content-type: application/json" \
-d '{"nodes":"["10.85.41.166","10.85.41.167","10.85.41.168"]"}
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "参数解析出错",
    "data": []
}
```
```json
{
    "code": 404,
    "msg": "error when delete id...",
    "data": []
}
```

#### 返回码解释

| 返回码  | 状态   | 描述       |
| ---- | ---- | -------- |
| 0    | 执行成功 |          |
| 400  | 执行失败 | 请求参数解析错误 |

------

### 8.根据node节点ip地址获取所在服务池接口

根据node节点ip地址获取所在服务池

#### 请求地址

| POST方法                                |
| ------------------------------------- |
| http://HOST/pool/search_by_ip/:iplist |

#### 请求参数

| 名称     | 类型     | 是否必须 | 描述       |
| ------ | ------ | ---- | -------- |
| iplist | string | 是    | 节点ip地址集合 |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data 参数中object对象说明

| 名称   | 类型             | 示例值                                 | 描述           |
| ---- | -------------- | ----------------------------------- | ------------ |
|      | map[string]int | {"10.85.41.160":1,"10.85.41.161":2} | 每个ip所在的服务池id |

#### 请求示例

```php
curl -X POST "http://HOST/pool/search_by_ip/10.85.41.160,10.85.41.161"  \
-H "Content-type: application/json" 
```

#### 响应示例

正常返回结果：
```json
{
    "code": 0,
    "msg": "success",
    "data": {"10.85.41.160":1,
    "10.85.41.161":2
    }
}
```

#### 返回码解释

| 返回码  | 状态   | 描述   |
| ---- | ---- | ---- |
| 0    | 执行成功 |      |

------

### 9.服务池扩容接口

服务池扩容

#### 请求地址

| POST方法                      |
| --------------------------- |
| http://HOST/pool/expand/:id |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述    |
| ---- | ---- | ---- | ----- |
| id   | int  | 是    | 服务池id |
| num  | int  | 是    | 扩容数量  |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

#### 请求示例

```php
curl -X POST "http://HOST/pool/expand/1  \
-H "Content-type: application/json" \
-H "Authorization: root" \
-d '{"num":2}'  
```
#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "扩容失败原因",
    "data": {}
}
```

#### 返回码解释

| 返回码  | 状态   | 描述     |
| ---- | ---- | ------ |
| 0    | 执行成功 |        |
| 400  | 执行失败 | 扩容失败原因 |

-------

### 10.服务池缩容接口

服务池缩容

#### 请求地址

| POST方法                      |
| --------------------------- |
| http://HOST/pool/shrink/:id |

#### 请求参数

| 名称    | 类型       | 是否必须 | 描述     |
| ----- | -------- | ---- | ------ |
| id    | int      | 是    | 服务池id  |
| nodes | []string | 是    | 节点ip列表 |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |


#### 请求示例

```php
curl -X POST "http://HOST/pool/shrink/1  \
-H "Content-type: application/json" \
-H "Authorization: root" \
-d '{"nodes":["10.85.41.160","10.85.41.161"]}  
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "缩容失败原因",
   "data": {}
}
```

#### 返回码解释

| 返回码  | 状态   | 描述     |
| ---- | ---- | ------ |
| 0    | 执行成功 |        |
| 400  | 执行失败 | 缩容失败原因 |

-------

### 11.服务池上线接口

服务池上线

#### 请求地址

| POST方法                      |
| --------------------------- |
| http://HOST/pool/deploy/:id |

#### 请求参数

| 名称      | 类型     | 是否必须 | 描述       |
| ------- | ------ | ---- | -------- |
| id      | int    | 是    | 服务池id    |
| max_num | int    | 是    | 每次最大上线数量 |
| tag     | string | 是    | 上线服务镜像名称 |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |


#### 请求示例

```php
curl -X POST "http://HOST/pool/deploy/1  \
-H "Content-type: application/json" \
-H "Authorization: root" \
-d '{"max_num":5,"tag":"myregistry/myweb:1.0"}  
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": {}
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "上线失败原因",
    "data": {}
}
```
#### 返回码解释

| 返回码  | 状态   | 描述     |
| ---- | ---- | ------ |
| 0    | 执行成功 |        |
| 400  | 执行失败 | 上线失败原因 |

-------

## 任务接口

### 1.创建任务执行步骤接口

创建任务执行步骤以及每步配置

#### 请求地址

| POST方法                       |
| ---------------------------- |
| http://HOST/task/impl/create |

#### 请求参数

| 名称    | 类型    | 是否必须 | 描述   |
| ----- | ----- | ---- | ---- |
| id    | int   | 否    | 任务id |
| name  | int   | 否    | 任务名称 |
| desc  | int   | 否    | 任务描述 |
| steps | array | 是    | 步骤   |

steps array中每个元素参数

| 名称           | 类型     | 是否必须 | 描述     |
| ------------ | ------ | ---- | ------ |
| name         | string | 是    | 步骤名称   |
| param_values | map    | 是    | 配置参数   |
| retry_times  | array  | 是    | 尝试次数   |
| ignore_error | array  | 是    | 是否忽略错误 |

####  返回参数

| 名称   | 类型    | 示例值      | 描述       |
| ---- | ----- | -------- | -------- |
| code | int   | 0        | 返回码      |
| msg  | string| "sucess" | 接口返回信息   |
| data | int   |          | 创建的任务流id |

#### 请求示例

```php
curl -X GET "http://HOST/task/impl/create"  \
-H "Content-type: application/json"  \
-d '{
    "name":"tpl1",
    "desc":"Template 1",
    "steps":[{ "name":"sleep",
                "param_values":{"time":1},
                "retry":{"retry_times":2,"ignore_error":false}
               },
               {"name":"echo_step",
                 "param_values":{"name":"aaa"},
                 "retry":{"retry_times":1,"ignore_error":true
               }
              }]

   }'
```

#### 响应示例

正常返回结果

```json
{
    "code": 0,
    "msg": "db server lost...",
    "data": 2,
}
```

异常返回结果

```json
{
    "code": 400,
    "msg": "json can not convert ",
    "data": {},
}
```
```json
{
    "code": 404,
    "msg": "step name not found",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                             |
| ---- | ---- | ------------------------------ |
| 0    | 执行成功 |                                |
| 400  | 执行失败 | 传入json参数无法转化                   |
| 404  | 执行失败 | 查询数据库失败或者改传入的step name找不到在数据库中 |

----------

### 2.删除任务执行步骤

删除任务执行步骤

#### 请求地址

| POST方法                       |
| ---------------------------- |
| http://HOST/task/impl/delete |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述       |
| ---- | ---- | ---- | -------- |
| id   | int  | 是    | 要删除的任务id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task/impl/delete' \
-H "Content-type: application/json" \
-d '{"id":1}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": null,
}
```

异常响应结果

```json
{
    "code": 404,
    "msg": "task impl not found",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述             |
| ---- | ---- | -------------- |
| 0    | 执行成功 |                |
| 404  | 执行失败 | 数据库中找不到该任务执行步骤 |

-------

### 3.获取任务执行步骤列表

获取任务执行步骤列表

#### 请求地址

| POST方法                     |
| -------------------------- |
| http://HOST/task/impl/list |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                                  |
| --------- | ---- | ---- | ----------------------------------- |
| page      | int  | 否    | 当前页码数，即本次API调用是获得结果的第几页，从1开始计数，默认为1 |
| page_size | int  | 否    | 当前页包含的结果数，默认结果数为10                  |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | string| "sucess" | 接口返回信息 |
| data        | array | [{},{}]  | 数据结果   |
| page        | int   | 1        | 当前页    |
| page_size   | int   | 10       | 当前页大小  |
| query_count | int   | 2        | 结果数    |

data 参数中object对象说明

| 名称    | 类型     | 示例值  | 描述       |
| ----- | ------ | ---- | -------- |
| id    | int    | 0    | 任务步骤id   |
| name  | string | 0    | 任务执行步骤名称 |
| desc  | string | 0    | 任务执行步骤描述 |
| Steps | string | 0    | 任务步骤     |

#### 请求示例

```php
curl -X GET "http://HOST/task/impl/list"  \
-H "Content-type: application/json"  \
-d "page=1" \
-d "page_size=10" 
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
            "id": 2,
            "name": "expand_nginx",
            "desc": "扩容nginx服务",
            "steps": [{"name":"return_vm","param_values":{"vm_type_id":1},"retry":{"retry_times":2,"ignore_error":false}}]
        }, {
            "id": 2,
            "name": "undeploy_nginx",
            "desc": "缩容nginx服务",
            "steps": [{"name":"return_vm","param_values":{"vm_type_id":1},"retry":{"retry_times":2,"ignore_error":false}}]
        },
    ],
    "page": 1,
    "page_size": 10,
    "query_count": 1
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": [],
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 执行成功 |                           |
| 400  | 执行失败 | 查询数据库失败，详细信息请查看返回结果的msg字段 |

-------

### 4.创建任务接口

创建任务

#### 请求地址

| POST方法                  |
| ----------------------- |
| http://HOST/task/create |

#### 请求参数

| 名称          | 类型     | 是否必须 | 描述        |
| ----------- | ------ | ---- | --------- |
| template_id | int    | 是    | 任务模板id    |
| task_name   | string | 是    | 任务名称      |
| timeout     | int    | 否    | 超时时间      |
| auto        | int    | 否    | 自动        |
| max_ratio   | int    | 否    | 最大比率      |
| max_num     | int    | 否    | 最大个数      |
| opr_user    | 用户名    | 是    | 操作用户      |
| nodes       | map    | 是    | 任务中含有的节点数 |
| params      | map    | 是    | 任务配置      |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          | 数据结果   |

#### 请求示例

```php
curl -X GET "http://HOST/cluster/list"  \
-H "Content-type: application/json"  \
-d ' {
		"template_id":1,
		"task_name":"expand_nignix",
		"timeout":10,
		"auto":1,
		"max_ratio":20,
		"max_num":10,
		"opr_user":"root",
		"nodes":[{"":object}],
		"params":[{"":""}]
	}' 
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": null,
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "error...",
    "data": [],
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                |
| ---- | ---- | ----------------- |
| 0    | 执行成功 |                   |
| 400  | 执行失败 | 详细信息请查看返回结果的msg字段 |

-------

### 5.开始任务

开始执行任务

#### 请求地址

| POST方法                           |
| -------------------------------- |
| http://HOST/task/start :"id":int |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 是    | 任务id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | ""       |        |

#### 请求示例

```php
curl -X POST ' http://HOST/task/start :"id":int' \
-H "Content-type: application/json" \
 -d '{"id":1}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": null
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "flow not found id 1",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 开启任务失败，详细信息请查看返回结果的msg字段 |

-------

### 6.暂停任务

暂停执行的任务

#### 请求地址

| POST方法                          |
| ------------------------------- |
| http://HOST/task/pause:"id":int |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 是    | 任务id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | ""       |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task/pause:"id":int ' \
-H "Content-type: application/json" \
 -d '{"id":1}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": null
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "pause flow 1 fails:",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 暂停任务失败，详细信息请查看返回结果的msg字段 |

-------

### 7.终止任务

#### 请求地址

终止任务

| POST方法                         |
| ------------------------------ |
| http://HOST/task/stop:"id":int |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 是    | 任务id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | ""       |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task/stop:"id":int' \
-H "Content-type: application/json" \
 -d '{"id":1}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": null,
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "stop flow 1 fails:",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 终止任务失败，详细信息请查看返回结果的msg字段 |

-------

### 8.获取任务日志

获取该任务执行的日志

#### 请求地址

| GET/ POST方法                    |
| ------------------------------ |
| http://HOST/task/:"id":int/log |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 是    | 任务名称 |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述      |
| ---- | ------ | -------- | ------- |
| code | int    | 0        | 返回码     |
| msg  | string | "sucess" | 接口返回信息  |
| data | object | Logs     | 获取的日志信息 |


Logs参数说明

| 名称             | 类型     | 示例值  | 描述     |
| -------------- | ------ | ---- | ------ |
| id             | int    |      | 日志id   |
| fid            | int    | 1    | 任务id   |
| batch_id       | int    | Logs | 任务池id  |
| correlation_id | int    | Logs | 任务关联id |
| message        | string | Logs | 日志信息   |
| ctime          | int    | Logs | 时间     |

####  请求示例

```php
curl -X POST 'http://HOST/task/2/log' \
-H "Content-type: application/json" 
```

####  响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
      "id":1,
      "fid":2,
      "batch_id":2,
      "correlation_id":2-2,
      "message":"run flow...",
      "ctime":1501162794
    }],
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "Flow not found",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                         |
| ---- | ---- | -------------------------- |
| 0    | 执行成功 |                            |
| 400  | 执行失败 | 获取任务日志失败，详细信息请查看返回结果的msg字段 |

-------

###  9.获取任务列表

获取所有的任务列表

####  请求地址

| POST方法                |
| --------------------- |
| http://HOST/task/list |


#### 请求参数

| 名称        | 类型   | 是否必须 | 描述    |
| --------- | ---- | ---- | ----- |
| page      | int  | 是    | 当前页   |
| page_size | int  | 是    | 当前页大小 |

#### 返回参数

| 名称          | 类型     | 示例值      | 描述     |
| ----------- | ------ | -------- | ------ |
| code        | int    | 0        | 返回码    |
| msg         | string | "sucess" | 接口返回信息 |
| data        | object |          | 数据结果   |
| page        | int    | 1        | 当前页    |
| page_size   | int    | 10       | 当前页大小  |
| query_count | int    | 2        | 返回结果数  |

#### 请求示例

```php
curl -X POST 'http://HOST/task/list' \
-H "Content-type: application/json" \
 -d '{"page":1, "page_size":10}' 
```

#### 响应示例

正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": [{
    "id":1,
    "template_id":2,
    "template_name":"aaa",
    "task_name":"bb",
    "pool_name":"pool",
    "Status":1,
    "options":[],
    "step_len":2,
    "opr_user":"root",
    "created":"2017-07-27 13:39:55",
    "updated":"2017-07-27 13:41:00",
    "stat":[1,1,1,1,1]
    }],
    "page":1,
    "page_size":10,
    "query_count":1
}
```

异常响应结果
```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 400  | 执行失败 | 获取列表，详细信息请查看返回结果的msg字段 |

-------

### 10.获取任务

#### 请求地址

| POST方法                     |
| -------------------------- |
| http://HOST/task/:"id":int |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 是    | 任务id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          | 数据结果   |

data 中对象参数说明

| 名称            | 类型     | 示例值      | 描述          |
| ------------- | ------ | -------- | ----------- |
| id            | int    | 0        | 任务id        |
| template_id   | int    | "sucess" | 模板id        |
| template_name | string |          | 模板名称        |
| pool_name     | string |          | 服务池名称       |
| Status        | int    |          | 状态          |
| options       | array  |          | 执行步骤        |
| step_len      | int    |          | 执行步长        |
| opr_user      | string |          | 操作用户        |
| created       | time   |          | 创建时间        |
| updated       | time   |          | 更新时间        |
| stat          | array  |          | 给任务个状态节点的个数 |

#### 请求示例

```php
curl -X POST ' http://HOST/task/2' \
-H "Content-type: application/json" \ 
```
#### 响应示例
正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": 
          [{"id":1,
          "template_id":2,
          "template_name":"aaa",
          "task_name":"bb",
          "pool_name":"pool",
          "Status":1,
          "options":[],
          "step_len":2,
          "opr_user":"root",
          "created":"2017-07-27 13:39:55",
          "updated":"2017-07-27 13:41:00",
          "stat":[1,1,1,1,1],
          }],
    "page":1,
    "page_size":10,
    "query_count":1
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": null,
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 获取任务失败，详细信息请查看返回结果的msg字段 |

-------

### 11.获取服务扩容任务列表

#### 请求地址
| POST方法                                  |
| --------------------------------------- |
| http://HOST/task/expandList/:poolId:int |


#### 请求参数

| 名称     | 类型   | 是否必须 | 描述    |
| ------ | ---- | ---- | ----- |
| poolId | int  | 是    | 服务池id |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | strng  | "sucess" | 接口返回信息 |
| data | object |          |        |


#### 请求示例

```php
curl -X POST 'http://HOST/task/expandList/2' \
-H "Content-type: application/json" \
```


#### 响应示例

正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": [{
      "id":1,
      "pool":{},
      "cron_items":[],
      "depend_items":[],
      "type":"expand",
      "exec_type:""crontab"
    }],
}
```
异常响应结果
```json
{
    "code": 400,
    "msg": "Bad pool id...",
    "data": {},
}
```


#### 返回码解释

| 返回码  | 状态   | 描述                             |
| ---- | ---- | ------------------------------ |
| 0    | 执行成功 |                                |
| 400  | 执行失败 | 获取扩容定时任务列表失败，详细信息请查看返回结果的msg字段 |

-------

### 12.获取服务上线上线列表

#### 请求地址

| POST方法                                  |
| --------------------------------------- |
| http://HOST/task/uploadList/:poolId:int |


#### 请求参数

| 名称   | 类型   | 是否必须 | 描述    |
| ---- | ---- | ---- | ----- |
| id   | int  | 是    | 服务池id |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |


#### 请求示例

```php
curl -X POST 'http://HOST/task/uploadList/1' \
-H "Content-type: application/json" \
```
#### 响应示例
正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": [{
        "id":1,
        "pool":{},
        "cron_items":[],
        "depend_items":[],
        "type":"expand",
        "exec_type":"crontab"
    }],
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "Bad pool id...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                                |
| ---- | ---- | --------------------------------- |
| 0    | 执行成功 |                                   |
| 400  | 执行失败 | 获取服务池上线定时任务列表失败，详细信息请查看返回结果的msg字段 |

-------

### 13.获取任务包含的节点

#### 请求地址

| POST方法                            |
| --------------------------------- |
| http://HOST/task/:"id":int/detail |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 是    | 任务id |


#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task/:"id":int/detail ' \
-H "Content-type: application/json"  
```

#### 响应示例

正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": [{
      "id":1,
      "ip":"127.0.0.1",
      "state":1,
      "steps":{},
      "pool_name":"pool",
      "vm_id":"3",
      "created":"2017-07-27 13:39:55"
    }],
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                         |
| ---- | ---- | -------------------------- |
| 0    | 执行成功 |                            |
| 400  | 执行失败 | 获取任务节点失败，详细信息请查看返回结果的msg字段 |

-------

### 14.保存定时任务

#### 请求地址

| POST方法                    |
| ------------------------- |
| http://HOST/task/saveTask |

####  请求参数

| 名称           | 类型     | 是否必须 | 描述                |
| ------------ | ------ | ---- | ----------------- |
| id           | int    | 是    | 任务id,如果该任务不存在id置0 |
| pool_id      | int    | 是    | 服务池id             |
| cron_itmes   | array  | 是    | 定时任务列表            |
| depend_itmes | array  | 是    | 依赖任务列表            |
| type         | string | 是    | 定时任务类型            |
| exec_type    | string | 是    | 定时任务执行类型          |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task/saveTask ' \
-H "Content-type: application/json" \
 -d '{
 "id":1, 
 "pool_id":2, 
 "cron_itmes": [],
 "depend_itmes": [],
 "type": " expand/upload",
 "exec_type":"crontab/depend"
 }' 
```
#### 响应示例
正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": true,
}
```

异常响应结果
```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                         |
| ---- | ---- | -------------------------- |
| 0    | 执行成功 |                            |
| 400  | 执行失败 | 保存定时任务失败，详细信息请查看返回结果的msg字段 |

-------

### 15.获取每个节点日志

#### 请求地址

| POST方法                         |
| ------------------------------ |
| http://HOST/task/;nsid"int/log |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| nsid | int  | 是    | 节点id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |

#### 请求示例
```php
curl -X POST 'http://HOST/task/2/log' \
-H "Content-type: application/json" 
```

#### 响应示例

正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": [{
          "id":1,
          "template_id":2,
          "template_name":"aaa",
          "task_name":"bb",
          "pool_name":"pool",
          "Status":1,
          "options":[],
          "step_len":2,
          "opr_user":"root",
          "created":"2017-07-27 13:39:55",
          "updated":"2017-07-27 13:41:00",
          "stat":[1,1,1,1,1],
    }],
    "page":1,
    "page_size":10,
    "query_count":1
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": null,
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                         |
| ---- | ---- | -------------------------- |
| 0    | 执行成功 |                            |
| 400  | 执行失败 | 获取节点日志出错，详细信息请查看返回结果的msg字段 |

-------

## 任务模板接口

###  1.创建任务模板

创建任务执行模板

####  请求地址

| POST方法                      |
| --------------------------- |
| http://HOST/task_tpl/create |

####  请求参数

| 名称    | 类型     | 是否必须 | 描述     |
| ----- | ------ | ---- | ------ |
| id    | string | 否    | 任务模板id |
| name  | string | 是    | 模板名称   |
| desc  | string | 是    | 模板描述   |
| steps | array  | 是    | 模板执行步骤 |

####   返回参数

| 名称   | 类型    | 示例值      | 描述      |
| ---- | ----- | -------- | ------- |
| code | int   | 0        | 返回码     |
| msg  | string| "sucess" | 接口返回信息  |
| data | int   |          | 创建的模板id |

#### 请求示例

```php
curl -X POST 'http://HOST/task_tpl/create' \
-H "Content-type: application/json" \
 -d '{
 "id":1, 
 "name":"Sample Platform",
 "desc": "平台",
 "steps":[],
 }' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": 2,
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 创建模板失败，详细信息请查看返回结果的msg字段 |

-------

### 2.获取任务模板列表

获取任务模板列表

#### 请求地址

| POST方法                    |
| ------------------------- |
| http://HOST/task_tpl/list |

#### 请求参数

| 名称        | 类型     | 是否必须 | 描述             |
| --------- | ------ | ---- | -------------- |
| page      | string | 是    | 当前页 ，从1开始，默认为1 |
| page_size | string | 是    | 当前页大小 ，默认为10   |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | ""       |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task_tpl/list \
-H "Content-type: application/json" \
 -d '{"page":1, "page_size":10}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
      "id":1,
      "name":"pool",
      "desc":"pool name",
      "steps":{}
    }],
    "page":1,
    "page_size":10,
    "query_count":2
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}

```

#### 返回码解释

| 返回码  | 状态   | 描述                           |
| ---- | ---- | ---------------------------- |
| 0    | 执行成功 |                              |
| 400  | 执行失败 | 获取任务模板列表失败，详细信息请查看返回结果的msg字段 |

-------

### 3.删除任务模板

删除任务模板

#### 请求地址

| POST方法                                |
| ------------------------------------- |
| http://HOST/task_tpl/delete/:"id":int |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 是    | 任务id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task_tpl/update/2' \
-H "Content-type: application/json" 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": null
}
```

异常响应结果

```json
{
    "code": 404,
    "msg": "data not found",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述               |
| ---- | ---- | ---------------- |
| 0    | 执行成功 |                  |
| 404  | 执行失败 | 删除任务模板失败，任务模板未找到 |

-------

### 4.获取任务模板

获取任务模板

#### 请求地址
| POST方法                         |
| ------------------------------ |
| http://HOST/task_tpl/:"id":int |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述     |
| ---- | ---- | ---- | ------ |
| id   | int  | 是    | 获取任务模板 |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task_tpl/2' \
-H "Content-type: application/json" \
```

####  响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": {
     "id":1,
     "name":"task impl",
     "desc":"impl",
     "steps":[{
       "name":"create_vm",
       "param_values":{},
       "retry":{}
     }]
    },
}
```

异常响应结果

```json
{
    "code": 404,
    "msg": "error...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述               |
| ---- | ---- | ---------------- |
| 0    | 执行成功 |                  |
| 404  | 执行失败 | 获取任务模板失败，任务模板未找到 |

-------

### 5.更新任务模板

更新任务模板

####  请求地址

| POST方法                                |
| ------------------------------------- |
| http://HOST/task_tpl/update/:"id":int |

####  请求参数

| 名称    | 类型     | 是否必须 | 描述     |
| ----- | ------ | ---- | ------ |
| id    | int    | 是    | 任务模板id |
| name  | string | 是    | 任务模板名称 |
| desc  | string | 是    | 任务模板描述 |
| steps | array  | 是    | 任务模板步骤 |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| id   | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task_tpl/update/2' \
-H "Content-type: application/json" \
 -d '{"id":1, 
 "name":"SamplePlatform",
 "desc":"Sample Platform", 
 "steps": [
   {
       "name":"create_vm",
       "param_values":{},
       "retry":{}
   }
 ]
 }' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": "",
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 更新模板失败，详细信息请查看返回结果的msg字段 |

-------

## 任务执行步骤接口

### 1.获取任务执行步骤列表

获取任务执行步骤列表

#### 请求地址

| POST方法                     |
| -------------------------- |
| http://HOST/task_step/list |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述         |
| --------- | ---- | ---- | ---------- |
| page      | int  | 是    | 当前页，默认从1开始 |
| page_size | int  | 是    | 当前页大小      |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |

#### 请求示例

```php
curl -X POST 'http://HOST/task_step/update/2' \
-H "Content-type: application/json" \
 -d '{"page":1, "page_size":10}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
      "id":1,
      "name":"action impl",
      "desc":"action descible",
      "type":"manual",
      "params":{}
    }],
    "page":1,
    "page_size":10,
    "query_count":1
}
```

异常响应结果
```json

```
#### 返回码解释

| 返回码  | 状态   | 描述   |
| ---- | ---- | ---- |
| 0    | 执行成功 |      |

-------

## 远程步骤接口

###  1.远程步骤列表接口

获取所有远程步骤列表

####   请求地址

| GET方法                        |
| ---------------------------- |
| http://HOST/remote_step/list |

####  请求参数

| 名称        | 类型   | 是否必须 | 描述                                  |
| --------- | ---- | ---- | ----------------------------------- |
| page      | int  | 否    | 当前页码数，即本次API调用是获得结果的第几页，从1开始计数，默认为1 |
| page_size | int  | 否    | 当前页包含的结果数，默认结果数为10                  |

####  返回参数

| 名称          | 类型    | 示例值     | 描述    |
| ----------- | ----- | ------- | ----- |
| code        | int   | 0       | 返回码   |
| data        | array | [{},{}] | 数据结果  |
| page        | int   | 1       | 当前页   |
| page_size   | int   | 10      | 当前页大小 |
| query_count | int   | 2       | 结果数   |

data 参数中object对象说明

| 名称      | 类型       | 示例值  | 描述     |
| ------- | -------- | ---- | ------ |
| id      | int      | 0    | 远程步骤id |
| name    | string   | 0    | 远程步骤名称 |
| desc    | string   | 0    | 远程步骤描述 |
| actions | []string | 0    | 命令顺序   |

####  请求示例

```php
curl -X GET "http://HOST/remote_step/list"  \
-H "Content-type: application/json"  \
-d "page=1" \
-d "page_size=10" 
```

####   响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
            "id": 2,
            "name": "add-default-image",
            "desc": "添加openstack缺省镜像",
            "action":["add-default-image"]
        }, {
            "id": 1,
            "name": "init_compute",
            "desc": "init_compute",
            "action":["init_compute"]
        },
    ],
    "page": 1,
    "page_size": 10,
    "query_count": 2
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": [],
}
```

####  返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 执行成功 |                           |
| 400  | 执行失败 | 查询数据库失败，详细信息请查看返回结果的msg字段 |

-------

###   2.创建步骤接口

创建步骤

#### 请求地址

| POST方法                         |
| ------------------------------ |
| http://HOST/remote_step/create |

#### 请求参数

| 名称      | 类型       | 是否必须 | 描述       |
| ------- | -------- | ---- | -------- |
| id      | int      | 是    | id       |
| name    | string   | 是    | 要增加的步骤名称 |
| desc    | string   | 是    | 步骤描述     |
| actions | []string | 是    | 命令       |

#### 返回参数

| 名称   | 类型    | 示例值      | 描述      |
| ---- | ----- | -------- | ------- |
| code | int   | 0        | 返回码     |
| msg  | string| "sucess" | 接口返回信息  |
| data | int   |          | 添加的步骤id |

#### 请求示例

```php
curl -X POST 'http://HOST/remote_step/create' \
-H "Content-type: application/json" \
-d '{"name":"step","desc":"Step", "actions":["action1", "action2"]}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": 1,
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

####  返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 400  | 执行失败 | 添加失败，详细信息请查看返回结果的msg字段 |

-------

###  3.修改远程步骤接口

修改步骤

#### 请求地址

| POST方法                             |
| ---------------------------------- |
| http://HOST/remote_step/update/:id |

#### 请求参数

| 名称      | 类型       | 是否必须 | 描述      |
| ------- | -------- | ---- | ------- |
| name    | string   | 是    | 步骤修改名称  |
| desc    | string   | 是    | 步骤修改描述  |
| actions | []string | 是    | 步骤修改的命令 |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | string | ""       |        |

#### 请求示例

```php
curl -X POST 'http://HOST/remote_step/update/2' \
-H "Content-type: application/json" \
 -d '{"name":"step","desc":"Step", "actions":["action1", "action2", "action3"]}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": "",
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": null,
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 读取参数失败，详细信息请查看返回结果的msg字段 |
| 404  | 执行失败 | 修改失败，详细信息请查看返回结果的msg字段   |

-------

###  4.删除步骤接口

删除步骤

#### 请求地址

| POST方法                             |
| ---------------------------------- |
| http://HOST/remote_step/delete/:id |

####  请求参数

| 名称   | 类型   | 示例值  | 描述       |
| ---- | ---- | ---- | -------- |
| id   | int  | 0    | 要删除的步骤id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | null     |        |

#### 请求示例

```php
curl -X POST 'http://$HOST/remote_step/delete/2' \
-H "Content-type: application/json" 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": "",
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": null,
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 404  | 执行失败 | 删除失败，详细信息请查看返回结果的msg字段 |

-------

###  5.获取步骤详情

获取步骤详情

#### 请求地址

| POST方法                      |
| --------------------------- |
| http://HOST/remote_step/:id |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述    |
| ---- | ---- | ---- | ----- |
| id   | int  | 是    | 步骤的id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data参数

| 名称      | 类型       | 示例值  | 描述   |
| ------- | -------- | ---- | ---- |
| id      | int      | 0    | 步骤Id |
| name    | string   |      | 步骤名称 |
| desc    | string   |      | 步骤描述 |
| actions | []string |      | 命令顺序 |

#### 请求示例

```php
curl 
-H "Content-type: application/json" \
-X POST 'http://$HOST/remote_step/3' \
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": {
      "id": 1,
      "name":"add-default-image",
      "desc": "添加openstack缺省镜像",
      "actions":{"add-default-image"}
    },
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "id is error...",
    "data": {},
}
```

```json
{
    "code": 404,
    "msg": "remote_step in db not found",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述          |
| ---- | ---- | ----------- |
| 0    | 执行成功 |             |
| 404  | 执行失败 | 数据库中查询不到该步骤 |

-------

## 远程命令接口

### 1.远程命令列表接口

获取所有远程命令列表

####  请求地址

| GET方法                   |
| ----------------------- |
| http://HOST/action/list |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                                  |
| --------- | ---- | ---- | ----------------------------------- |
| page      | int  | 否    | 当前页码数，即本次API调用是获得结果的第几页，从1开始计数，默认为1 |
| page_size | int  | 否    | 当前页包含的结果数，默认结果数为10                  |

#### 返回参数

| 名称          | 类型    | 示例值     | 描述    |
| ----------- | ----- | ------- | ----- |
| code        | int   | 0       | 返回码   |
| data        | array | [{},{}] | 数据结果  |
| page        | int   | 1       | 当前页   |
| page_size   | int   | 10      | 当前页大小 |
| query_count | int   | 2       | 结果数   |

data 参数中object对象说明

| 名称     | 类型                     | 示例值  | 描述     |
| ------ | ---------------------- | ---- | ------ |
| id     | int                    | 0    | 远程命令id |
| name   | string                 | 0    | 远程命令名称 |
| desc   | string                 | 0    | 远程命令描述 |
| params | map[string]interface{} | 0    | 参数列表   |

#### 请求示例

```php
curl -X GET "http://HOST/action/list"  \
-H "Content-type: application/json"  \
-d "page=1" \
-d "page_size=10" 
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
            "id": 2,
            "name": "start_docker",
            "desc": "启动docker",
            "params":{"host":"string","name":"string","tag":"string"},
        }, {
            "id": 1,
            "name": "check_port",
            "desc": "检查端口",
            "params":{"check_port":"integer","check_times":"integer"},
     
        },
    ],
    "page": 1,
    "page_size": 10,
    "query_count": 2
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": [],
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 执行成功 |                           |
| 400  | 执行失败 | 查询数据库失败，详细信息请查看返回结果的msg字段 |

-------

### 2.创建远程命令接口

创建远程命令

#### 请求地址

| POST方法                    |
| ------------------------- |
| http://HOST/action/create |


#### 请求参数

| 名称     | 类型                     | 是否必须 | 描述   |
| ------ | ---------------------- | ---- | ---- |
| id     | int                    | 是    | 命令id |
| name   | string                 | 是    | 命令名字 |
| desc   | string                 | 是    | 命令描述 |
| params | map[string]interface{} | 是    | 参数列表 |

#### 返回参数

| 名称   | 类型    | 示例值     | 描述     |
| ---- | ----- | ------- | ------ |
| code | int   | 0       | 返回码    |
| err  | string| "error" | 接口返回信息 |
| data | int   |         | 命令id   |

#### 请求示例

```php
curl -X POST 'http://HOST/action/create' \
-H "Content-type: application/json" \
-d '{"name":"command","desc":"Command", "params":{"time": "integer"}}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": 1,
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 400  | 执行失败 | 创建失败，详细信息请查看返回结果的msg字段 |

-------

### 3.修改远程命令接口

修改远程命令

#### 请求地址

| POST方法                        |
| ----------------------------- |
| http://HOST/action/update/:id |

#### 请求参数

| 名称     | 类型                     | 是否必须 | 描述          |
| ------ | ---------------------- | ---- | ----------- |
| name   | string                 | 是    | 远程命令修改名称    |
| desc   | string                 | 是    | 远程命令修改描述    |
| params | map[string]interface{} | 是    | 远程命令修改的参数列表 |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | string | ""       |        |

#### 请求示例

```php
curl -X POST 'http://HOST/action/update/2' \
-H "Content-type: application/json" \
 -d '{"desc":"Command xx", "params":{"time": "integer"}}", "action3"]}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": "",
}
```

异常响应结果

```json
{
    "code": 404,
    "msg": "db server lost...",
    "data": null,
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 读取参数失败，详细信息请查看返回结果的msg字段 |
| 404  | 执行失败 | 修改失败，详细信息请查看返回结果的msg字段   |

-------

###  4.删除远程命令接口

删除远程命令

#### 请求地址

| POST方法                        |
| ----------------------------- |
| http://HOST/action/delete/:id |

#### 请求参数

| 名称   | 类型   | 示例值  | 描述         |
| ---- | ---- | ---- | ---------- |
| id   | int  | 0    | 要删除的远程命令id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | null     |        |

#### 请求示例

```php
curl -X POST 'http://$HOST/action/delete/:id' \
-H "Content-type: application/json" 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": "",
}
```

异常响应结果

```json
{
    "code": 404,
    "msg": "db server lost...",
    "data": null,
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 404  | 执行失败 | 删除失败，详细信息请查看返回结果的msg字段 |

-------

###   5.获取远程命令详情

获取远程命令详情

#### 请求地址

| POST方法                 |
| ---------------------- |
| http://HOST/action/:id |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述      |
| ---- | ---- | ---- | ------- |
| id   | int  | 是    | 远程命令的id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data参数

| 名称     | 类型                     | 示例值  | 描述     |
| ------ | ---------------------- | ---- | ------ |
| id     | int                    | 0    | 远程命令Id |
| name   | string                 |      | 远程命令名称 |
| desc   | string                 |      | 远程命令描述 |
| params | map[string]interface{} |      | 参数列表   |

####  请求示例

```php
curl 
-H "Content-type: application/json" \
-X POST 'http://$HOST/action/:id' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": {
     	     "id": 2,
            "name": "start_docker",
            "desc": "启动docker",
            "params":{"host":"string","name":"string","tag":"string"},
    },
}
```

异常响应结果

```json
{
    "code": 404,
    "msg": "action in db not found",
    "data": {},
}
```



#### 返回码解释

| 返回码  | 状态   | 描述            |
| ---- | ---- | ------------- |
| 0    | 执行成功 |               |
| 404  | 执行失败 | 数据库中查询不到该远程命令 |

-------

###  6.创建远程命令任务模板实现接口

创建任务执行模板

#### 请求地址

| POST方法                         |
| ------------------------------ |
| http://HOST/action/impl/create |

#### 请求参数

| 名称    | 类型     | 是否必须 | 描述     |
| ----- | ------ | ---- | ------ |
| id    | string | 否    | 任务模板id |
| name  | string | 是    | 模板名称   |
| desc  | string | 是    | 模板描述   |
| steps | array  | 是    | 模板执行步骤 |

#### 返回参数

| 名称   | 类型    | 示例值      | 描述      |
| ---- | ----- | -------- | ------- |
| code | int   | 0        | 返回码     |
| msg  | string| "sucess" | 接口返回信息  |
| data | int   |          | 创建的模板id |

#### 请求示例

```php
curl -X POST 'http://HOST/action/impl/create ' \
-H "Content-type: application/json" \
 -d '{
 "id":1, 
 "name":"Sample Platform",
 "desc": "平台",
 "steps":[].
 }' 
```

#### 响应示例

正常响应结果
```json
{
    "code": 0,
    "msg": "success",
    "data": 2,
}
```
异常响应结果
```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 创建模板失败，详细信息请查看返回结果的msg字段 |

-------

### 7.获取远程命令任务模板实现列表接口

获取任务模板列表

#### 请求地址

| POST方法                       |
| ---------------------------- |
| http://HOST/action/impl/list |

#### 请求参数

| 名称        | 类型     | 是否必须 | 描述             |
| --------- | ------ | ---- | -------------- |
| page      | string | 是    | 当前页 ，从1开始，默认为1 |
| page_size | string | 是    | 当前页大小 ，默认为10   |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | ""       |        |

#### 请求示例

```php
curl -X POST 'http://HOST/action/impl/list\
-H "Content-type: application/json" \
 -d '{"page":1, "page_size":10}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": [{
      "id":1,
      "name":"pool",
      "desc":"pool name",
      "steps":{}
    }],
    "page":1,
    "page_size":10,
    "query_count":2
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                           |
| ---- | ---- | ---------------------------- |
| 0    | 执行成功 |                              |
| 400  | 执行失败 | 获取任务模板列表失败，详细信息请查看返回结果的msg字段 |

-------

### 8.删除远程命令任务模板实现接口

删除任务模板

#### 请求地址

| POST方法                                   |
| ---------------------------------------- |
| http://HOST/action/impl/delete/:"id":int |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| id   | int  | 是    | 任务id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object |          |        |

#### 请求示例

```php
curl -X POST 'http://HOST/action/impl/delete/2' \
-H "Content-type: application/json" 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": null
}
```

异常响应结果

```json
{
    "code": 404,
    "msg": "data not found",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述               |
| ---- | ---- | ---------------- |
| 0    | 执行成功 |                  |
| 404  | 执行失败 | 删除任务模板失败，任务模板未找到 |

-------

## 远程命令实现接口

### 1.获取所有远程命令列表

获取用户自定义的所有远程命令列表

#### 请求地址

| GET方法                    |
| ------------------------ |
| http://HOST/actimpl/list |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                                  |
| --------- | ---- | ---- | ----------------------------------- |
| page      | int  | 否    | 当前页码数，即本次API调用是获得结果的第几页，从1开始计数，默认为1 |
| page_size | int  | 否    | 当前页包含的结果数，默认结果数为10                  |

#### 返回参数

| 名称          | 类型    | 示例值     | 描述    |
| ----------- | ----- | ------- | ----- |
| code        | int   | 0       | 返回码   |
| data        | array | [{},{}] | 数据结果  |
| page        | int   | 1       | 当前页   |
| page_size   | int   | 10      | 当前页大小 |
| query_count | int   | 2       | 结果数   |

data 参数中object对象说明

| 名称        | 类型                     | 示例值  | 描述     |
| --------- | ---------------------- | ---- | ------ |
| id        | int                    | 0    | 远程命令id |
| type      | string                 | 0    | 远程命令类型 |
| action_id | int                    | 0    | 命令组id  |
| template  | map[string]interface{} | 0    | 模板     |

#### 请求示例

```php
curl -X GET "http://HOST/actimpl/list"  \
-H "Content-type: application/json"  \
-d "page=1" \
-d "page_size=10" 
```

#### 响应示例

正常返回结果：
```json
{
    "code": 0,
    "msg": "success",
    "data": [{
            "id": 2,
            "template": {"action":{"content":"docker run -d --net=\"{{host}}\" --name {{name}} {{tag}} ","module":"longscript"}},
            "type": "ansible",
            "action_id":2,
        }, {
            "id": 1,
            "template": {"action":{"args":"echo {{echo_word}} ","module":"shell"}},
            "type": "ansible",
            "action_id":4,
        },
    ],
    "page": 1,
    "page_size": 10,
    "query_count": 2,
}
```

异常返回结果：

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": [],
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 执行成功 |                           |
| 400  | 执行失败 | 查询数据库失败，详细信息请查看返回结果的msg字段 |

-------

###  2.创建远程命令实现接口

创建远程命令实现

#### 请求地址

| POST方法                     |
| -------------------------- |
| http://HOST/actimpl/create |

#### 请求参数

| 名称        | 类型                     | 是否必须 | 描述       |
| --------- | ---------------------- | ---- | -------- |
| id        | int                    | 是    | id       |
| type      | string                 | 是    | 要增加的命令类型 |
| action_id | int                    | 是    | 命令id     |
| template  | map[string]interface{} | 是    | 模板       |

#### 返回参数

| 名称   | 类型    | 示例值      | 描述        |
| ---- | ----- | -------- | --------- |
| code | int   | 0        | 返回码       |
| msg  | string| "sucess" | 接口返回信息    |
| data | int   |          | 添加的命令实现id |

#### 请求示例

```php
curl -X POST 'http://HOST/actimpl/create' \
-H "Content-type: application/json" \
-d '{"action_id":14,"type":"ansible","template":{"action":{"module":"shell","args":"which {{exec}}"}}}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": 1,
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 400  | 执行失败 | 添加失败，详细信息请查看返回结果的msg字段 |

-------

### 3.修改远程命令实现接口

修改远程命令实现

#### 请求地址

| POST方法                         |
| ------------------------------ |
| http://HOST/actimpl/update/:id |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述     |
| ---- | ---- | ---- | ------ |
| id   | int  | 是    | 命令实现id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | string | ""       |        |

#### 请求示例

```php
curl -X POST 'http://HOST/actimpl/update/:id' \
-H "Content-type: application/json" \
 -d '{"action_id":14,"type":"ansible","template":{"action":{"module":"shell","args":"which {{exec}}"}}}' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": "",
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": null,
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                       |
| ---- | ---- | ------------------------ |
| 0    | 执行成功 |                          |
| 400  | 执行失败 | 读取参数失败，详细信息请查看返回结果的msg字段 |

-------

###   4.删除远程命令实现接口

删除远程命令实现

#### 请求地址

| POST方法                         |
| ------------------------------ |
| http://HOST/actimpl/delete/:id |

#### 请求参数

| 名称   | 类型   | 示例值  | 描述         |
| ---- | ---- | ---- | ---------- |
| id   | int  | 0    | 要删除的命令实现id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | null     |        |

#### 请求示例

```php
curl -X POST 'http://$HOST/actimpl/delete/:id' \
-H "Content-type: application/json" 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": "",
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "db server lost...",
    "data": null,
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                     |
| ---- | ---- | ---------------------- |
| 0    | 执行成功 |                        |
| 400  | 执行失败 | 删除失败，详细信息请查看返回结果的msg字段 |

-------

###  5.获取远程命令实现详情

获取远程命令实现详情

#### 请求地址

| POST方法                  |
| ----------------------- |
| http://HOST/actimpl/:id |

#### 请求参数

| 名称       | 类型   | 是否必须 | 描述   |
| -------- | ---- | ---- | ---- |
| actionId | int  | 是    | 命令id |

#### 返回参数

| 名称   | 类型     | 示例值      | 描述     |
| ---- | ------ | -------- | ------ |
| code | int    | 0        | 返回码    |
| msg  | string | "sucess" | 接口返回信息 |
| data | object | {}       | 返回结果   |

data参数

| 名称        | 类型                     | 示例值  | 描述   |
| --------- | ---------------------- | ---- | ---- |
| id        | int                    | 0    | 步骤Id |
| type      | string                 |      | 步骤名称 |
| action_id | int                    |      | 命令id |
| template  | map[string]interface{} | 模板   |      |

#### 请求示例

```php
curl 
-H "Content-type: application/json" \
-X POST 'http://$HOST/actimpl/:id' 
```

#### 响应示例

正常响应结果

```json
{
    "code": 0,
    "msg": "success",
    "data": {
      "id": 1,
      "action_id":2,
      "type": "ansible",
      "template":{"action":{"args":"echo {{echo_word}} ","module":"shell"}}
    },
}
```

异常响应结果

```json
{
    "code": 400,
    "msg": "id is error...",
    "data": {},
}
```

```json
{
    "code": 404,
    "msg": "actimpl in db not found",
    "data": {},
}
```

#### 返回码解释

| 返回码  | 状态   | 描述          |
| ---- | ---- | ----------- |
| 0    | 执行成功 |             |
| 400  | 执行失败 | 传入参数id有误    |
| 404  | 执行失败 | 数据库中查询不到该步骤 |
