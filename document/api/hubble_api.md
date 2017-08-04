# 服务发现API

## 服务注册

### 1.获取服务发现类型

获取用户创建的所有服务发现类型列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/balance.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| page | int  | 否    | 请求的列表页数量，为空时默认是1    |
| fIdx      | string  | 否    | 服务发现类型名称，为空时查询所有 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |
| title        | array | ["#","名称","类型","用户","时间","#"]  | 返回列表头   |
| content        | object   | [{},{}]        | 服务发现类型详细信息    |
| count   | string   | "2"       | 当前结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 参数中object对象说明:

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "1"    | id号 |
| name | string | "default_service_name"    | 服务发现类型名称 |
| type  | string | "NGINX"   | 类型  |
| content   | string    | "{}"    | 详细信息 |
| create_time | string | "2016-11-15 22:16:50"    | 创建时间 |
| update_time | string | "2016-11-15 22:16:50"    | 更新时间 |
| opr_user  | string | "system"    | 创建者/更新者  |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/balance.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
```
#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "#",
        "名称",
        "类型",
        "用户",
        "时间",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "id": "1",
            "name": "default_service_name",
            "type": "NGINX",
            "content": "{\"group_id\":\"1\",\"name\":\"default.upstream\",\"port\":\"8080\",\"weight\":\"20\",\"script_id\":\"2\"}",
            "create_time": "2016-11-15 22:16:50",
            "update_time": "2016-11-15 22:16:50",
            "opr_user": "system"
        }
    ],
    "count": "1",
    "pageCount": 1,
    "page": 1,
    "pageSize": 20
}
```

异常返回结果：

```json
{
    "code": 1,
    "msg": "Illegal request, please login at first."
}
```
#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |




### 2.新增/修改服务发现类型

创建新的服务发现类型

#### 请求地址
| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/balance.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为insert或者为update |
| data      | object  | 是    | 详细信息(目前为aliyun或nginx的信息) |

##### Nginx的 data 参数中object对象说明：

示例：

```object
{"name":"21212","type":"NGINX","id":"","content":"{\"group_id\":\"1\",\"name\":\"default.upstream\",\"port\":\"8080\",\"weight\":\"20\",\"script_id\":\"2\"}"}
```

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| name      | string  | 是    | 服务发现名称 |
| type      | string  | 是    | 服务发现类型:NGINX |
| id      | string  | 否    | id |
| content      | object  |   是  | 服务发现类型的关联选项，包含下面几列 |
| group_id      | string  | 是    | 服务发现分组id |
| name      | string  | 是    | upstream名 |
| port      | string  | 是    | 端口 |
| weight      | string  |   是  | 权重 |
| script_id      | string  |   是  | 发布脚本id |

##### Aliyun 的 data 参数中object对象说明：

示例：

```json
{"name":"11212","type":"SLB","id":"","content":"{\"region\":\"cn-qingdao\",\"slb_id\":\"lbm5e7198depwn40w6kz\",\"weight\":\"100\"}"}
```

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| name      | string  | 是    | 服务发现名称 |
| type      | string  | 是    | 服务发现类型:SLB |
| id      | string  | 否    | id |
| content      | object  |   是  | 服务发现类型的关联选项，包含下面几列 |
| region      | string  | 是    | 地域 |
| slb_id      | string  | 是    | 阿里云SLB的id |
| weight      | string  | 是    | 权重 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |
| title        | array | ["#","名称","类型","用户","时间","#"]  | 返回列表头   |
| content        | object   | [{},{}]        | 服务发现类型详细信息    |
| count   | string   | "2"       | 当前结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 参数中object对象说明

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "1"    | id号 |
| name | string | "default_service_name"    | 服务发现类型名称 |
| type  | string | "NGINX"   | 类型  |
| content   | string    | "{}"    | 详细信息 |
| create_time | string | "2016-11-15 22:16:50"    | 创建时间 |
| update_time | string | "2016-11-15 22:16:50"    | 更新时间 |
| opr_user  | string | "system"    | 创建者/更新者  |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/balance.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
-d "data={'name':'11212','type':'SLB','id":'','content':'{\'region\':\'cn-qingdao\',\'slb_id\':\'lbm5e7198depwn40w6kz\',\'weight\':\'100\'}'}"
```
#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```
异常返回结果：
```json
{
    "code": 1,
    "msg": "Illegal request, please login at first."
}
```
#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 3.删除服务发现类型

删除用户创建的服务发现类型

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/balance.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为delete |
| data      | object  | 是    | 服务发现类型的详细信息 |

data 参数中object对象说明:

示例：
```json
{"id":"3"}
```

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| id      | string  | 是    | 服务发现类型id |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回状态码    |
| action         | strng | "delete" | action信息，执行失败时出现该参数 |
| msg        | string | "success"  | 返回执行结果   |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/balance.php"  \
-H "Content-type: application/json"  \
-d "action=delete" \
-d "{'id':'3'}"
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "delete",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |


## 七层Nginx

### 1.获取Nginx分组列表

获取用户创建的所有Nginx分组列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_group.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| page | int  | 否    | 列表页数量，默认是1    |
| fIdx      | string  | 否    | Nginx分组的名称 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "lis"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |
| title | array | ["#","名称","单元","用户","时间","#"] | 返回列表头 |
| content        | array   | [{},{}]        | 类型详细信息    |
| count   | string   | "2"       | 结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 的参数中object对象数组说明：

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "1"    | id号 |
| name | string | "default_group"    | 服务发现类型名称 |
| opr_user  | string | ""    | 创建者  |
| create_time  | string | "1970-01-01 00:00:00"    | 创建/更新时间  |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_group.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "#",
        "名称",
        "单元",
        "用户",
        "时间",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "id": "1",
            "name": "default_group",
            "opr_user": "",
            "create_time": "1970-01-01 00:00:00"
        }
    ],
    "count": "1",
    "pageCount": 1,
    "page": 1,
    "pageSize": 20
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "lis",
    "msg": "Param Error!"
}
```
#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |


### 2.创建Nginx分组

创建新的Nginx分组

#### 请求地址
| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_group.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为insert |
| data | object  | 是    | 新的Nginx分组详细信息   |
data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| name      | string  | 是    |  新的Nginx分组名称 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "insertsss"  | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |

#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_group.php"  \
-H "Content-type: application/json"  \
-d "action=insert" \
-d "data={'name':'New Nginx Group'}" 
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "inserts",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 3.更新/删除Nginx分组

更新/删除Nginx分组

#### 请求地址
| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_group.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为update或delete |
| data | object  | 是    | Nginx分组详细信息   |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| id      | int  | 是    |  Nginx分组的id |
| name      | string  | 是    |  Nginx分组的新名称，删除时不需要此项 |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "upd"  | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |

#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_group.php"  \
-H "Content-type: application/json"  \
-d "action=insert" \
-d "data={'name':'New Nginx Group','id':'1'}" 
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "upd",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 4.获取Nginx单元列表

获取用户创建的所有Nginx单元列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_unit.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| page | int  | 否    | 列表页数量，默认是1    |
| fGroup | int  | 是    | Nginx分组的id    |
| fIdx      | string  | 否    | Nginx单元的名称，默认全部 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "lis"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |
| title | array | ["#","名称","单元","用户","时间","#"] | 返回列表头 |
| content        | array   | [{},{}]        | 类型详细信息    |
| count   | string   | "2"       | 结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 的参数中object对象数组说明：

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "1"    | id号 |
| name | string | "default_unit"    | 服务发现类型名称 |
|  group_id    |  string | "1"    | 隶属Nginx分组的id |
| opr_user  | string | "system"    | 创建者  |
| create_time  | string | "2016-11-15 22:09:44"    | 创建/更新时间  |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_unit.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
-d "fGroup=1" 
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "#",
        "单元名称",
        "隶属分组",
        "下属节点",
        "配置",
        "用户",
        "时间",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "id": "1",
            "name": "default_unit",
            "group_id": "1",
            "opr_user": "system",
            "create_time": "2016-11-15 22:09:44"
        }
    ],
    "count": "1",
    "pageCount": 1,
    "page": 1,
    "pageSize": 20
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "list33",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 5.更新/删除Nginx单元列表

更新/删除Nginx单元列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_unit.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为update或delete |
| data | object  | 是    | Nginx分组的id    |

data 参数中object对象说明:

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| group_id   | int    | 1    | 隶属Nginx分组的id，删除时不需要该选项 |
| id   | int    | 1    | Nginx单元的id |
| name | string | "1"    | Nginx单元的名称，删除时不需要该选项 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "updatess" | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_unit.php"  \
-H "Content-type: application/json"  \
-d "action=update" \
-d "data={'group_id':1,'name':'大苏打大苏打撒旦','id':1}" 
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "updatewwwwww",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 6.获取Nginx单元的下属节点列表

获取Nginx单元的下属节点列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_node.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| page | int  | 否    | 列表页数量，默认是1    |
| fUnit   | int    | 是    | Nginx单元的id |
| fIdx   | int    | 否    | Nginx单元的节点ip，默认查询所有 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "lis"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |
| title| array| ["*","#","单元","IP","创建时间","用户","#"] |返回列表头 |
| content        | array   | [{},{}]        | 类型详细信息    |
| count   | string   | "2"       | 结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 的参数中object对象数组说明：

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "1"    | 节点id |
| ip | string | "default_unit"    | 节点IP |
|  unit_id    |  string | "1"    | 隶属Nginx的单元id |
| opr_user  | string | "system"    | 创建者  |
| create_time  | string | "2016-11-15 22:09:44"    | 创建/更新时间  |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_node.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
-d "fUnit=1" 
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "<input type=\"checkbox\" id=\"selectAll\" onclick=\"checkAll(this)\"/>",
        "#",
        "单元",
        "IP",
        "创建时间",
        "用户",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "id": "3",
            "ip": "101.1.1.1",
            "unit_id": "1",
            "opr_user": "root",
            "create_time": "2017-08-03 15:29:08"
        }
    ],
    "count": "1",
    "pageCount": 1,
    "page": 1,
    "pageSize": 20
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "listd",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 7.新增Nginx单元的下属节点

新增Nginx单元的下属节点

#### 请求地址
| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_node.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为insert |
| sid | int  | 是    | Nginx的分组id    |
| nodeips   | string | 是 | 增加节点的ip,支持逗号,分号,空格间隔添加多个 |


#### 返回参数 

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "insertsss" | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_node.php"  \
-H "Content-type: application/json"  \
-d "action=insert" \
-d "sid=1" \
-d "nodeips=1.4.5.6"
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "msg": "exits:12.5.6.7"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 8.删除Nginx单元的下属节点

删除Nginx单元的下属节点

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_node.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为delete |
| nodeips   | string | 是 | 节点的ip,支持逗号,分号,空格间隔添加多个 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "deleteeee" | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_node.php"  \
-H "Content-type: application/json"  \
-d "action=delete" \
-d "nodeips=1.4.5.6"
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "msg": "unit_id error"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |


### 9.获取Upstream列表

获取Upstream列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_upstream.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| page | int  | 否    | 请求的列表页数量，为空时默认是1    |
| fGroup      | int  | 是    | Nginx分组id |
| fIdx   | string | 否 | Upstream文件名称 |


#### 返回参数 

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "lis"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |
| title | array | ["#","文件名称","隶属分组","用户","最近更新","发布","#"] | 返回列表头 |
| content        | array   | [{},{}]        | 类型详细信息    |
| count   | string   | "2"       | 结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 的参数中object对象数组说明：

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "1"    | id号 |
| name | string | "default_unit"    | 服务发现类型名称 |
|  group_id    |  string | "1"    | 隶属Nginx分组的id |
|  is_consul    |  string | "0"    | 已Consul化(0:否) |
|  deprecated    |  string | "0"    | 已废弃(0:否) |
|  release_id    |  string | "0"    | 发布序号 |
| create_time  | string | "2016-11-15 22:09:44"    | 创建时间  |
| update_time  | string | "2017-08-02 12:45:42"    | 更新时间  |
| opr_user  | string | "system"    | 操作用户  |



#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_upstream.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
-d "fGroup=1"
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "#",
        "文件名称",
        "隶属分组",
        "用户",
        "最近更新",
        "发布",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "id": "1",
            "name": "default.upstream",
            "group_id": "1",
            "is_consul": "0",
            "deprecated": "0",
            "release_id": "0",
            "create_time": "2016-11-15 22:11:23",
            "update_time": "2017-08-02 12:45:42",
            "opr_user": "system"
        }
    ],
    "count": 1,
    "pageCount": 1,
    "page": 1,
    "pageSize": 20
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "list1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 10.新建Upstream文件

新建Upstream文件

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_upstream.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为insert |
| data | object  | 是    | upstream的详细信息    |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| group_id      | string  | 是    |  Nginx分组id |
| name      | string  | 是    |  upstream的文件名 |
| content      | string  | 是    |  upstream的文件内容 |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "inse"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |

#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_upstream.php"  \
-H "Content-type: application/json"  \
-d "action=insert" \
-d "data={'group_id':1,'name':'dadadad','content':'dadadadad'}"
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "inserta",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |


### 11.修改/删除Upstream文件

修改/删除Upstream文件

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_upstream.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为update或delete |
| data | object  | 是    | upstream的详细信息    |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| group_id      | string  | 是    |  Nginx分组id，删除时不需要此字段 |
| id      | string  | 是    |  upstream的id |
| content      | string  | 是    |  upstream的文件内容，删除时不需要此字段 |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "upda"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |

#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_upstream.php"  \
-H "Content-type: application/json"  \
-d "action=insert" \
-d "data={'group_id':1,'id':'4','content':'dadadadad'}"
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```


异常返回结果：

```json
{
    "code": 1,
    "action": "update1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |


### 12.发布Upstream文件

发布Upstream文件

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/nginx_upstream.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为publish |
| data | object  | 是    | upstream的详细信息    |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| upstream_id      | string  | 是    | upstream的id |
| unit_ids      | string  | 是    |  目标单元id，多个以逗号分隔 |
| tunnel      | string  | 是    |  发布方式，目前只有ANSIBLE |
| script_id      | string  | 是 |  upstream的发布脚本 |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "pub"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |

#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/nginx_upstream.php"  \
-H "Content-type: application/json"  \
-d "action=publish" \
-d 'data={"upstream_id": "1", "tunnel": "ANSIBLE", "script_id": "1", "unit_ids": "1"}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "pub",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

## 阿里云SLB

### 1.获取SLB列表

获取SLB列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST//api/for_cloud/slb.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| pagesize | int  | 否    | 页大小，设置为1000    |
| fRegion      | string  | 是    | SLB地域 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |
| title        | array | ["#","名称","服务地址","服务端口","后端端口","Backend","Listener","实例状态","类型","创建时间","#"]  | 返回列表头   |
| content        | object   | [{},{}]        | 服务发现类型详细信息    |
| count   | string   | "2"       | 当前结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 参数中object对象说明:

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| LoadBalancerId | string | "lbm5e7198depwn40w6kz"    | id号 |
| LoadBalancerName | string | ""    | SLB_Id |
| LoadBalancerStatus  | string | "active"   | 状态  |
| Address   | string    | "139.129.85.40"    | IP地址 |
| RegionId | string | "cn-qingdao"    | 地域ID |
| RegionIdAlias | string | "cn-qingdao"    | 地域别名 |
| AddressType  | string | "internet"    | 地址类型  |
| VSwitchId  | string | ""    | 子网  |
| VpcId  | string | ""    | VPC-ID  |
| NetworkType  | string | "classic"    | 网络类型  |
| Bandwidth  | int | 0   | 带宽大小  |
| InternetChargeType  | int | 0    | 计费类型  |
| CreateTime  | string | "2016-07-28T11:55Z"    | 创建时间  |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_cloud/slb.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
-d "fRegion=cn-qingdao"
```
#### 响应示例

正常返回结果：
```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "#",
        "名称",
        "服务地址",
        "服务端口",
        "后端端口",
        "Backend",
        "Listener",
        "实例状态",
        "类型",
        "创建时间",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "LoadBalancerId": "lbm5e7198depwn40w6kz",
            "LoadBalancerName": "",
            "LoadBalancerStatus": "active",
            "Address": "139.129.85.40",
            "RegionId": "cn-qingdao",
            "RegionIdAlias": "cn-qingdao",
            "AddressType": "internet",
            "VSwitchId": "",
            "VpcId": "",
            "NetworkType": "classic",
            "Bandwidth": 0,
            "InternetChargeType": 0,
            "CreateTime": "2016-07-28T11:55Z"
        }
    ],
    "count": 0,
    "pageCount": 1,
    "page": 1,
    "pageSize": 20
}
```

异常返回结果：

```json
{
    "code": 1,
    "msg": "Aliyun SDK Request Error: URL: https://slb.aliyuncs.com/?RegionId=cn-qingda&Format=JSON&Version=2014-05-15&Action=DescribeLoadBalancers&Signature=xa4exVwobjtLzTrs%2BtrC3CGfg7A%3D&HttpMethod=GET&AccessKeyId=zFaewKYsce0TpQvv&SignatureMethod=HMAC-SHA1&SignatureVersion=1.0&TimeStamp=2017-08-03T10:37:57Z&SignatureNonce=5575b9d97c4b2bba97759314df0dba53 RequestId: 48077CA3-2A1B-49E3-9D95-56FE5656CC47 HostId: slb.aliyuncs.com Code: InvalidParameter Message: The specified region is not exist.",
    "remote": "{\"code\":1,\"http_code\":500,\"url\":\"http://jupiter:8080/v1/slb/list/cn-qingda?page=1&pageSize=20\",\"msg\":\"Aliyun SDK Request Error: URL: https://slb.aliyuncs.com/?RegionId=cn-qingda&Format=JSON&Version=2014-05-15&Action=DescribeLoadBalancers&Signature=xa4exVwobjtLzTrs%2BtrC3CGfg7A%3D&HttpMethod=GET&AccessKeyId=zFaewKYsce0TpQvv&SignatureMethod=HMAC-SHA1&SignatureVersion=1.0&TimeStamp=2017-08-03T10:37:57Z&SignatureNonce=5575b9d97c4b2bba97759314df0dba53 RequestId: 48077CA3-2A1B-49E3-9D95-56FE5656CC47 HostId: slb.aliyuncs.com Code: InvalidParameter Message: The specified region is not exist.\"}",
    "page": 1,
    "pageSize": 20
}
```
#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 2.新建SLB列表

获取SLB列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_cloud/slb.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为insert |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| AddressType      | string  | 是    | 网络类型 |
| Bandwidth      | int  | 是    |  带宽大小 |
| InternetChargeType      | string  | 是    |  计费类型 |
| LoadBalancerName      | string  | 是    |  SLB名称 |
| RegionId      | string  | 是    |  区域ID |




#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |



#### 请求示例

```php
curl -X POST "http://HOST/api/for_cloud/slb.php"  \
-H "Content-type: application/json"  \
-d "action=insert" \
-d 'data={"AddressType": "internet", "Bandwidth": 1, "InternetChargeType": "paybytraffic", "LoadBalancerName": "鹅","RegionId":"cn-qingdao"}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "insert1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 3.删除SLB

删除SLB

#### 请求地址


| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_cloud/slb.php |
#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为delete |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| LoadBalancerId      | string  | 是    | id号 |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_cloud/slb.php"  \
-H "Content-type: application/json"  \
-d "action=delete" \
-d 'data={"LoadBalancerId": "lbm5e7198depwn40w6kz"}'
```
#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "delete1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 4.启用关闭SLB

启用关闭SLB

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_cloud/slb.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是   | 请求操作的类型，参数为inactive或active |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| LoadBalancerId      | string  | 是    | id号 |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_cloud/slb.php"  \
-H "Content-type: application/json"  \
-d "action=inactive" \
-d 'data={"LoadBalancerId": "lb-m5ep7k40dwmv1oyxm2mnl"}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "inactive1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

## 脚本管理

### 1.获取脚本列表

获取脚本列表


#### 请求地址

| POST/GET方法  |
| ------------------------ |
| http://HOST/api/for_hubble/shell.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| page | int  | 否    | 列表页数量，默认是1    |
| fIdx      | string  | 否    | 脚本名称，默认查询全部 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "lis"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |
| title | array | ["#","名称","描述","用户","时间","#"] | 返回列表头 |
| content        | array   | [{},{}]        | 类型详细信息    |
| count   | string   | "2"       | 结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 的参数中object对象数组说明：

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "1"    | 脚本id |
| name | string | "default_group"    | 脚本名称 |
| desc | string | "更新主配置脚本"    | 脚本描述 |
| content | string | ""    | 脚本详细信息 |
| opr_user  | string | "system"    | 创建者  |
| create_time  | string | "1970-01-01 00:00:00"    | 创建时间  |
| update_time  | string | "1970-01-01 00:00:00"    | 更新时间  |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/shell.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "#",
        "名称",
        "描述",
        "用户",
        "时间",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "id": "1",
            "name": "updateMainConf.sh",
            "desc": "更新主配置脚本",
            "content": "#!/bin/bash\nHUBBLE_UNIT_ID=\"{{HUBBLE_UNIT_ID}}\"\nHUBBLE_GROUP_ID=\"{{HUBBLE_GROUP_ID}}\"\nHUBBLE_RSYNC_HOST=\"{{HUBBLE_RSYNC_HOST}}\"\n\nHUBBLE_OUTFILE=\"/tmp/alterationMainConf.out\"\ncat > $HUBBLE_OUTFILE<< EOF\nEOF\n\nexec 1> $HUBBLE_OUTFILE  2> $HUBBLE_OUTFILE\n\nconfdir=\"/usr/local/nginx_conf\"\ntime_echo(){\n    echo `date +%F\"-\"%T`\" \"$*\n}\n#RSYNC MAIN CONFIG FILE nginx.conf FROM HUBBLE SERVER\n\ntime_echo rsync -argtv \"${HUBBLE_RSYNC_HOST}/group_${HUBBLE_GROUP_ID}/unit_${HUBBLE_UNIT_ID}/current/main/\" $confdir/\nrsync -argtv \"${HUBBLE_RSYNC_HOST}/group_${HUBBLE_GROUP_ID}/unit_${HUBBLE_UNIT_ID}/current/main/\" $confdir/ \n\nif [ $? -ne 0 ]; then\n    time_echo \"rsync file fail\" && exit 1\nfi\n\ntime_echo rsync nginx_conf successful...\n#CONFIGRATION CHECK AND RELOAD\ndocker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -t\ntime_echo docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -t\nif [ $? -eq 0 ];then\n    docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -s reload\n    if [ $? -ne 0 ]; then\n        time_echo \"reload nginx failed\" && exit 1\n    fi\n    time_echo docker exec opendcp_lb_ngx_ctn /usr/local/nginx/sbin/nginx -s reload\n    exit 0\nfi\ntime_echo \"check nginx_conf failed…\" && exit 1",
            "create_time": "2016-11-28 12:52:17",
            "update_time": "2016-11-28 12:52:17",
            "opr_user": "system"
        }
    ],
    "count": 1,
    "pageCount": 1,
    "page": 1,
    "pageSize": 20
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "list1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |


### 2.新建脚本

新建脚本

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_cloud/slb.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为insert |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| name      | string  | 是    | 脚本名称 |
| desc      | string  | 是    |  脚本描述 |
| content      | string  | 是    |  脚本信息 |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |



#### 请求示例

```php
curl -X POST "http://HOST/api/for_cloud/slb.php"  \
-H "Content-type: application/json"  \
-d "action=insert" \
-d 'data={"name": "12121","desc":"test","content":"dadadadad"}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "insert1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 3.更新脚本

更新脚本

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_cloud/slb.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为update |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| name      | string  | 是    | 脚本名称 |
| desc      | string  | 是    |  脚本描述 |
| content      | string  | 是    |  脚本信息 |
| id      | int  | 是    |  脚本id |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |



#### 请求示例

```php
curl -X POST "http://HOST/api/for_cloud/slb.php"  \
-H "Content-type: application/json"  \
-d "action=update" \
-d 'data={"name": "12121","desc":"test","content":"dadadadad","id":3}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "update111",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 4.删除脚本

删除脚本

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_cloud/slb.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为delete |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| id      | int  | 是    |  脚本id |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |



#### 请求示例

```php
curl -X POST "http://HOST/api/for_cloud/slb.php"  \
-H "Content-type: application/json"  \
-d "action=delete" \
-d 'data={"id":3}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "delete11",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

## 授权管理

### 1.获取AppKey列表

获取AppKey列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/appkey.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| page | int  | 否    | 列表页数量，默认是1    |
| fIdx      | string  | 否    | appkey名称，默认查询全部 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "lis"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |
| title | array | ["#","AppKey","名称","描述","#"] | 返回列表头 |
| content        | array   | [{},{}]        | 类型详细信息    |
| count   | string   | "2"       | 结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 的参数中object对象数组说明：

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "1"    | appkey的id |
| appkey | string | "048EA63AFD38D629993819EF980DE5AE" | appkey的内容 |
| name | string | "default_appkey"    | 名称 |
| describe | string | "默认配置，建议删除"    | 描述 |



#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/appkey.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "#",
        "AppKey",
        "名称",
        "描述",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "id": "1",
            "appkey": "048EA63AFD38D629993819EF980DE5AE",
            "name": "default_appkey",
            "describe": "默认配置，建议删除"
        }
    ],
    "page": 1,
    "pageSize": 20,
    "pageCount": 1,
    "count": 1
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "list1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 2.新建Appkey

新建Appkey

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/appkey.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为insert |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| name      | string  | 是    | 脚本名称 |
| desc      | string  | 是    |  脚本描述 |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |



#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/appkey.php"  \
-H "Content-type: application/json"  \
-d "action=insert" \
-d 'data={"name":"12121","desc":"dadad"}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "insert1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 3.删除Appkey

删除Appkey

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/appkey.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为delete |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| key      | string  | 是    | Appkey的标识号 |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/appkey.php"  \
-H "Content-type: application/json"  \
-d "action=delete" \
-d 'data={"key": "048EA63AFD38D629993819EF980DE5AE"}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "delete1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 4.获取接口列表

获取接口列表

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/interface.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| page | int  | 否    | 列表页数量，默认是1    |
| fIdx      | string  | 否    | 接口地址，默认查询全部 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "lis"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |
| title | array | ["#","接口","描述","方法","#"] | 返回列表头 |
| content        | array   | [{},{}]        | 类型详细信息    |
| count   | string   | "2"       | 结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 的参数中object对象数组说明：

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "1"    | 接口id |
| addr | string | "/v1/secure/appkey/list_interface/" | 接口地址 |
| desc | string | "接口列表"    | 接口描述 |
| method | string | "GET"    | 接口类型 |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/interface.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "#",
        "接口",
        "描述",
        "方法",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "id": "1",
            "addr": "/v1/secure/appkey/list_interface/",
            "desc": "接口列表",
            "method": "GET"
        }
    ],
    "count": 1,
    "pageCount": 1,
    "page": 1,
    "pageSize": 1
}
```


异常返回结果：

```json
{
    "code": 1,
    "action": "list1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 5.新建接口

新建接口

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/interface.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为insert |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| addr      | string  | 是    | 接口地址 |
| method      | string  | 是    |  接口方式 |
| desc      | string  | 是    |  接口描述 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |



#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/interface.php"  \
-H "Content-type: application/json"  \
-d "action=insert" \
-d 'data={"addr":"qqwqwqwqw","desc":"qwqwqwqw","method":"POST"}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "insert1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

### 6.删除接口

删除接口

#### 请求地址

| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/interface.php |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为delete |
| data      | object  | 是    | 详细信息 |

data 参数中object对象说明:

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| id      | int  | 是    | 接口的id |

#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| msg         | strng | "sucess" | 接口返回结果 |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/interface.php"  \
-H "Content-type: application/json"  \
-d "action=delete" \
-d 'data={"id": "5"}'
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success"
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "delete1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |

## 操作日志

### 1.获取操作日志列表

获取操作日志列表

#### 请求地址
| POST/GET方法                    |
| ------------------------ |
| http://HOST/api/for_hubble/oprlog.php |
#### 请求参数

| 名称        | 类型   | 是否必须 | 描述                  |
| --------- | ---- | ---- | ----------------------------------- |
| action      | string  | 是    | 请求操作的类型，为list |
| page | int  | 否    | 列表页数量，默认是1    |
| fIdx      | string  | 否    | 操作日志地址，默认查询全部 |


#### 返回参数

| 名称          | 类型    | 示例值      | 描述     |
| ----------- | ----- | -------- | ------ |
| code        | int   | 0        | 返回码    |
| action        | string   | "lis"        | action参数不正确时出现  |
| msg         | strng | "sucess" | 接口返回结果 |
| title | array | ["#","模块","行为","用户","时间","#"] | 返回列表头 |
| content        | array   | [{},{}]        | 类型详细信息    |
| count   | string   | "2"       | 结果数量  |
| pageCount | int   | 1        | 结果页数量    |
| page        | int   | 1        | 当前页    |
| pageSize   | int   | 20       | 当前页大小  |

content 的参数中object对象数组说明：

| 名称   | 类型     | 示例值  | 描述   |
| ---- | ------ | ---- | ---- |
| i   | int    | 1    | 序号 |
| id | string | "50"    | 操作id |
| module | string | "Secure" | 模块 |
| operation | string | "Del Interface"    | 行为 |
| opr_time | string | "2017-08-03 21:25:57"    | 操作时间 |
| appkey | string | "6741bc42-9e21-4763-977c-ac3a1fc0bdd8" | AppKey |
| user | string | "root"    | 操作用户 |
| args | string | "5 : Array"    | 请求参数 |


#### 请求示例

```php
curl -X POST "http://HOST/api/for_hubble/oprlog.php"  \
-H "Content-type: application/json"  \
-d "action=list" \
```

#### 响应示例

正常返回结果：

```json
{
    "code": 0,
    "msg": "success",
    "title": [
        "#",
        "模块",
        "行为",
        "用户",
        "时间",
        "#"
    ],
    "content": [
        {
            "i": 1,
            "id": "50",
            "module": "Secure",
            "operation": "Del Interface",
            "opr_time": "2017-08-03 21:25:57",
            "appkey": "6741bc42-9e21-4763-977c-ac3a1fc0bdd8",
            "user": "root",
            "args": "5 : Array"
        }
    ],
    "count": "50",
    "pageCount": 3,
    "page": 1,
    "pageSize": 20
}
```

异常返回结果：

```json
{
    "code": 1,
    "action": "list1",
    "msg": "Param Error!"
}
```

#### 返回码解释

| 返回码  | 状态   | 描述                        |
| ---- | ---- | ------------------------- |
| 0    | 成功 |           执行成功               |
| 1    | 失败 |      执行失败，详细信息请查看返回错误信息       |


