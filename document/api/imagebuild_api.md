# 镜像市场API
## 打包系统

### 1.创建项目接口
创建项目
#### 请求地址
| POST方法                      |
| --------------------------- |
| http://HOST/api/project/new |
#### 请求参数
| 名称          | 类型     | 是否必须 | 描述        |
| ----------- | ------ | ---- | --------- |
| projectName | string | 是    | 当前创建的项目名称 |
| operator    | string | 是    | 操作用户      |
#### 返回参数
| 名称      | 类型     | 示例值  | 描述     |
| ------- | ------ | ---- | ------ |
| code    | int    | 1000 | 返回码    |
| errMsg  | string | ""   | 接口返回信息 |
| content | object | ""   | 数据结果   |
#### 请求示例
```php
curl -X POST "http://HOST/api/project/new "  \
-H "Content-type: application/json"  \
-d '{"projectName":"myweb","operator":"root"}'
```
#### 响应示例
正常返回结果：
```json
{
    "code": 10000,
    "errMsg": "ok",
    "content": "",
}
```
异常返回结果：
```json
{
    "code": 10006,
    "errMsg": "parameter invalid",
    "content": -1,
}
```
```json
{
    "code": 10100,
    "errMsg": "server internal error",
    "content": -1,
}
```
#### 返回码解释
| 返回码   | 状态   | 描述          |
| ----- | ---- | ----------- |
| 10000 | 执行成功 |             |
| 10006 | 执行失败 | 参数非法，参数传入有误 |
| 10100 | 执行失败 | 接口内部出错      |
### 2.删除项目接口
删除项目
#### 请求地址
| POST方法                         |
| ------------------------------ |
| http://HOST/api/project/delete |
#### 请求参数
| 名称          | 类型     | 是否必须 | 描述   |
| ----------- | ------ | ---- | ---- |
| projectName | string | 是    | 项目名称 |
| operator    | string | 是    | 创建用户 |

#### 返回参数
| 名称      | 类型     | 示例值      | 描述     |
| ------- | ------ | -------- | ------ |
| code    | int    | 0        | 返回码    |
| errMsg  | strng  | "sucess" | 接口返回信息 |
| content | string |          | 数据结果   |
#### 请求示例
```php
curl -X POST 'http://HOST/api/project/delete' \
-H "Content-type: application/json" \
-d '{"projectName":"myweb","operator":"root"}' 
```
#### 响应示例
正常响应结果
```json
{
    "code": 10000,
    "errmsg": "ok",
    "content": "",
}
```
异常响应结果
```json
{
     "code": 10006,
     "errmsg": "parameter invalid",
     "content": -1,
}
```
#### 返回码解释
| 返回码   | 状态   | 描述                      |
| ----- | ---- | ----------------------- |
| 10000 | 执行成功 |                         |
| 10006 | 执行失败 | 参数不合法，详细信息请查看返回结果的msg字段 |
| 10100 | 执行失败 | 接口内部出错                  |

### 3.克隆项目接口
项目克隆
#### 请求地址
| POST方法                        |
| ----------------------------- |
| http://HOST/api/project/clone |

#### 请求参数
| 名称             | 类型     | 是否必须 | 描述      |
| -------------- | ------ | ---- | ------- |
| srcProjectName | string | 是    | 克隆源项目名称 |
| dstProjectName | string | 是    | 克隆项目名称  |
| operator       | string | 是    | 操作者     |
#### 返回参数
| 名称      | 类型     | 示例值   | 描述     |
| ------- | ------ | ----- | ------ |
| code    | int    | 10000 | 返回码    |
| errMsg  | strng  | ""    | 接口返回信息 |
| content | object |       |        |
#### 请求示例
```php
curl -X POST 'http://HOST/api/project/clone ' \
-H "Content-type: application/json" \
 -d '{"srcProjectName":"source", 
 "dstProjectName":"target", 
 "operator": "root"}' 
```
#### 响应示例
正常响应结果
```json
{
    "code": 10000,
    "errMsg": "success",
    "content": "",
}
```
异常响应结果
```json
{
    "code": 10006,
    "msg": "parameter invalid",
    "content": -1,
}
```
```json
{
    "code": 10003,
    "msg": "project to delete not exist",
    "content": -1,
}
```
```json
{
    "code": 10002,
    "msg": "src project not exist",
    "content": -1,
}
```
#### 返回码解释
| 返回码   | 状态   | 描述          |
| ----- | ---- | ----------- |
| 10000 | 执行成功 |             |
| 10006 | 执行失败 | 参数非法，参数传入有误 |
| 10100 | 执行失败 | 接口内部出错      |
| 10002 | 执行失败 | 克隆源项目不存在    |
| 10003 | 执行失败 | 项目删除时，项目不存在 |
### 4.项目详情接口

获取项目详情

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/api/project/info?projectName=myweb&operator=admin |
#### 请求参数
| 名称          | 类型     | 示例值  | 描述   |
| ----------- | ------ | ---- | ---- |
| projectName | string | 是    | 项目名称 |
| operator    | string | 是    | 创建用户 |
#### 返回参数
| 名称      | 类型     | 示例值      | 描述     |
| ------- | ------ | -------- | ------ |
| code    | int    | 10000    | 返回码    |
| errMsg  | strng  | "sucess" | 接口返回信息 |
| content | object | null     |        |

content参数

| 名称                   | 类型     | 示例值  | 描述              |
| -------------------- | ------ | ---- | --------------- |
| creator              | int    | 0    | 项目创建者           |
| name                 | strng  |      | 项目名称            |
| createTime           | string |      | 项目创建时间          |
| lastModifyTime       | string |      | 最近修改时间          |
| lastModifyOperator   | string |      | 最近修改用户          |
| Cluster              | string |      | 集群              |
| DefineDockerFileType | string |      | 定义的DockerFile类型 |

#### 请求示例

```php
curl -X GET 'http://$HOST/api/project/info' \
-H "Content-type: application/json" 
```
#### 响应示例
正常响应结果
```json
{
    "code": 10000,
    "errMsg": "success",
    "data": "",
}
```
异常响应结果
```json
{
    "code": 10008,
    "errMsg": "Project: xxx  to query info not exist",
    "data": null,
}
```
#### 返回码解释
| 返回码   | 状态   | 描述          |
| ----- | ---- | ----------- |
| 0     | 执行成功 |             |
| 10008 | 执行失败 | 项目不存在       |
| 10006 | 执行失败 | 参数非法，参数传入有误 |
| 10100 | 执行失败 | 接口内部出错      |


### 5.获取项目列表接口
获取项目列表

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/api/project/list?page=1&page_size=10&projectName=myweb |

#### 请求参数
| 名称          | 类型   | 是否必须 | 描述    |
| ----------- | ---- | ---- | ----- |
| projectName | int  | 是    | 项目名称  |
| page        | int  | 是    | 页码    |
| page_size   | int  | 是    | 当前页大小 |
#### 返回参数
| 名称          | 类型     | 示例值   | 描述            |
| ----------- | ------ | ----- | ------------- |
| code        | int    | 10000 | 返回码           |
| errMsg      | strng  | ""    | 接口返回信息        |
| content     | object |       | 返回结果          |
| page        | int    |       | 当前页，从1开始，默认为1 |
| page_size   | int    |       | 当前页大小，默认为10   |
| total_count | int    |       | 结果总数          |

data参数

| 名称                   | 类型     | 示例值  | 描述             |
| -------------------- | ------ | ---- | -------------- |
| creator              | string |      | 创建者            |
| name                 | strng  |      | 项目名称           |
| createTime           | string |      | 创建时间           |
| lastModifyTime       | string |      | 最后一次修改时间       |
| lastModifyOperator   | string |      | 最后一次修改者        |
| Cluster              | string |      | 隶属镜像集群         |
| DefineDockerFileType | string |      | dockerfile定义类型 |
#### 请求示例
```php
curl 
-H "Content-type: application/json" \
-X POST 'http://$HOST/api/project/list' 
```
#### 响应示例
正常响应结果
```json
{
    "code": 10000,
    "errMsg": "success",
    "content": "",
}
```
异常响应结果
```json
{
    "code": 10006,
    "msg": "parameter invalid",
    "content": -1,
}
```
#### 返回码解释
| 返回码   | 状态   | 描述          |
| ----- | ---- | ----------- |
| 10000 | 执行成功 |             |
| 10006 | 执行失败 | 参数非法，参数传入有误 |

### 6.项目保存接口
保存项目
#### 请求地址
| POST方法                       |
| ---------------------------- |
| http://HOST/api/project/save |
#### 请求参数
| 名称                   | 类型     | 是否必须 | 描述           |
| -------------------- | ------ | ---- | ------------ |
| project              | string | 是    | 项目名称         |
| Cluster              | string | 是    | 隶属镜像集群       |
| DefineDockerFileType | string | 是    | dockerfile类型 |
| addOrUpdate          | string | 是    | 操作用户         |
| $$plugin             | array  | 是    | 项目所用插件       |
#### 返回参数
| 名称      | 类型     | 示例值   | 描述     |
| ------- | ------ | ----- | ------ |
| code    | int    | 10000 | 返回码    |
| errMsg  | strng  | ""    | 接口返回信息 |
| content | object |       | 返回结果   |
#### 请求示例
```php
curl 
-H "Content-type: application/json" \
-X POST 'http://HOST/api/project/save' \
-d "{
  project:"aa"
  Cluster:"default"
  DefineDockerFileType:"define"
  addOrUpdate:"add"
  $$plugin:[{
  }]
}"
```
#### 响应示例
正常响应结果
```json
{
    "code": 10000,
    "errMsg": "success",
    "content": "",
}
```
异常响应结果
```json
{
    "code": 10006,
    "errMsg": "project already exist",
    "content": "project: aaa already exist ""
}
```
```json
{
    "code": 10001,
    "errMsg": "parameter invalid",
    "content": "projectName contains special char: *"
}
```
```json
{
    "code": 10100,
    "errMsg": "server internal error",
    "content": ""
}
```
#### 返回码解释
| 返回码   | 状态   | 描述             |
| ----- | ---- | -------------- |
| 10000 | 执行成功 |                |
| 10006 | 执行失败 | 参数非法，项目已存在不能添加 |
| 10100 | 执行失败 | 接口内部出错         |

### 7.构建项目接口
构建项目
#### 请求地址
| POST方法                        |
| ----------------------------- |
| http://HOST/api/project/build |
#### 请求参数
| 名称          | 类型     | 示例值  | 描述   |
| ----------- | ------ | ---- | ---- |
| projectName | string | 是    | 项目名称 |
| operator    | string | 是    | 创建者  |
| tag         | string | 是    | 镜像版本 |
#### 返回参数
| 名称      | 类型     | 示例值      | 描述     |
| ------- | ------ | -------- | ------ |
| code    | int    | 10000    | 返回码    |
| errmsg  | strng  | "sucess" | 接口返回信息 |
| content | object | null     | 返回结果   |
content参数
| 名称    | 类型     | 示例值  | 描述   |
| ----- | ------ | ---- | ---- |
| idStr | string | 0    | id   |
#### 请求示例
```php
curl -X POST 'http://$HOST/api/project/build' \
-H "Content-type: application/json" 
```
#### 响应示例
正常响应结果

```json
{
    "code": 10000,
    "msg": "success",
    "data": "",
}
```
异常响应结果

```json
{
    "code": 10008,
    "msg": "Project: xxx  to query info not exist",
    "data": null,
}
```

#### 返回码解释
| 返回码   | 状态   | 描述          |
| ----- | ---- | ----------- |
| 0     | 执行成功 |             |
| 10007 | 执行失败 | 构建项目不存在     |
| 10006 | 执行失败 | 参数非法，参数传入有误 |
| 10100 | 执行失败 | 接口内部出错      |

### 8.获取项目构建结果接口
获取项目构建结果
#### 请求地址
| GET方法                                    |
| ---------------------------------------- |
| http://HOST/api/project/buildHistory?projectName=myweb&operator=admin |
#### 请求参数
| 名称          | 类型     | 是否必须 | 描述   |
| ----------- | ------ | ---- | ---- |
| projectName | string | 是    | 项目名称 |
| operator    | string | 是    | 操作者  |
#### 返回参数
| 名称      | 类型     | 示例值   | 描述     |
| ------- | ------ | ----- | ------ |
| code    | int    | 10000 | 返回码    |
| errMsg  | string | ""    | 接口返回信息 |
| content | object |       |        |
#### 请求示例
```php
curl 
-H "Content-type: application/json" \
-X POST 'http://HOST/api/project/buildHistory?projectName=myweb&operator=admin' 
```
#### 响应示例
正常响应结果
```json
{
    "code": 10000,
    "errMsg": "success",
    "content": "",
}
```
异常响应结果
```json
{
    "code": 10006,
    "errMsg": "project already exist",
    "content": "project: aaa already exist ",
}
```
```json
{
   "code": 10001,
   "errMsg": "parameter invalid",
    "content": "projectName contains special char: *",
}
```
```json
{
    "code": 10100,
    "errMsg": "server internal error",
    "content": "",
}
```
#### 返回码解释
| 返回码   | 状态   | 描述             |
| ----- | ---- | -------------- |
| 10000 | 执行成功 |                |
| 10006 | 执行失败 | 参数非法，项目已存在不能添加 |
| 10100 | 执行失败 | 接口内部出错         |

### 9.项目插件扩展接口
插件扩展接口
#### 请求地址
| POST方法                                   |
| ---------------------------------------- |
| http://HOST/api/plugin/extension/interface |
#### 请求参数
| 名称     | 类型     | 是否必须 | 描述   |
| ------ | ------ | ---- | ---- |
| plugin | string | 是    |      |
| method | string | 是    | 操作者  |
#### 返回参数
| 名称      | 类型     | 示例值   | 描述     |
| ------- | ------ | ----- | ------ |
| code    | int    | 10000 | 返回码    |
| errMsg  | strng  | ""    | 接口返回信息 |
| content | object |       |        |
#### 请求示例
```php
curl -H "Content-type: application/json" \
-X POST 'http://HOST/api/plugin/extension/interface' \
-d 'plugin="aa"' \
-d 'method="download file"' \
-d '{  
}'
```
#### 响应示例
正常响应结果
```json
{
    "code": 10000,
    "errMsg": "ok",
    "content": null
}
```
异常响应结果
```json
{
    "code": 10004,
    "errMsg": "param error",
    "content": {},
}
```
```json
{
    "code": 10100,
    "errMsg": "server internal error",
    "content": {},
}
```
#### 返回码解释
| 返回码   | 状态   | 描述     |
| ----- | ---- | ------ |
| 10000 | 执行成功 |        |
| 10004 | 执行失败 | 参数非法   |
| 10100 | 执行失败 | 接口内部出错 |

### 10.获取项目配置页面
获取项目构建镜像配置信息
#### 请求地址
| GET方法                                    |
| ---------------------------------------- |
| http://HOST/view/config/independent?projectName=aaa |
#### 请求参数
| 名称          | 类型     | 是否必须 | 描述   |
| ----------- | ------ | ---- | ---- |
| projectName | string | 是    | 项目名称 |
#### 返回参数
| 名称   | 类型     | 示例值  | 描述data     |
| ---- | ------ | ---- | ---------- |
| data | string |      | 返回配置页面HTML |

#### 请求示例
```php
curl 
-H "Content-type: application/json" \
-X GET 'http://HOST/view/config/independent'
```
#### 响应示例
正常响应结果
```json
{
  "data":"<html>....</html>"
}
```
#### 返回码解释


## 镜像仓库接口

该部分调用Harbor镜像存储接口，默认用户名admin，密码为Harbor12345

### 1.获取项目列表接口
获取项目集群列表
#### 请求地址
| GET方法                                    |
| ---------------------------------------- |
| http://HarborHOST/projects?page=1&page_size=10&project_name=myweb |
#### 请求参数
| 名称           | 类型     | 是否必须 | 描述        |
| ------------ | ------ | ---- | --------- |
| page         | string | 否    | 当前页       |
| pagesize     | string | 否    | 当前页大小     |
| project_name | string | 是    | 为空字符串表示全部 |

#### 返回参数
| 名称   | 类型     | 示例值  | 描述data |
| ---- | ------ | ---- | ------ |
| data | object |      | 返回项目列表 |

#### 请求示例
```php
curl 
-H "Content-type: application/json" \
-X GET 'http://HarborHOST/projects?page=1&page_size=10&project_name=yweb' 
```
#### 响应示例
正常响应结果
```json
[{
  "Togglable":true,
  "creation_time":"2017-07-26T08:37:13Z",
  "creation_time_str":"",
  "current_user_role_id":1,
  "deleted":0,
  "name":"base",
  "owner_id":1,
  "owner_name":"",
  "project_id":3,
  "public":1,
  "repo_count":0,
  "update_time":"2017-07-26T08:37:13Z"
},
]
```
#### 返回码解释

### 2.获取镜像列表
获取镜像列表
#### 请求地址
| GET方法                                    |
| ---------------------------------------- |
| http://HOST/repositories?page=1&page_size=10&project_id=1&q= |
#### 请求参数
| 名称         | 类型     | 是否必须 | 描述   |
| ---------- | ------ | ---- | ---- |
| page       | string | 是    | 项目名称 |
| page_size  | string | 是    | 项目名称 |
| project_id | string | 是    | 项目名称 |
| q          | string | 是    | 项目名称 |
#### 返回参数
| 名称   | 类型    | 示例值  | 描述data |
| ---- | ----- | ---- | ------ |
| data | array |      | 返回镜像结果 |

#### 请求示例
```php
curl 
-H "Content-type: application/json" \
-X GET 'http://HOST/repositories?page=1&page_size=10&project_id=1&q=' 
```
#### 响应示例
正常响应结果
```json
[{
  "name": "default_cluster/myweb"
}]
```
#### 返回码解释

### 3.获取镜像标签列表
获取镜像标签列表
#### 请求地址
| GET方法                                    |
| ---------------------------------------- |
| http://HOST/repositories/tags?page=1&page_size=10&project_id=1&repo_name=abc |
#### 请求参数
| 名称        | 类型     | 是否必须 | 描述   |
| --------- | ------ | ---- | ---- |
| page      | int    | 是    | 项目名称 |
| pagesize  | int    | 是    | 项目名称 |
| repo_name | string | 是    | 项目名称 |
#### 返回参数
| 名称   | 类型    | 示例值  | 描述data |
| ---- | ----- | ---- | ------ |
| data | array |      | 返回镜像结果 |

#### 请求示例
```php
curl 
-H "Content-type: application/json" \
-X POST 'http://HOST/repositories/tags?page=1&page_size=10&project_id=1&repo_name=abc'
```
#### 响应示例
正常响应结果
```json
[{
  "name": "1.0"
},
{
  "name": "1.2"
}]
```
#### 返回码解释

### 4.获取镜像详细信息
获取镜像详细信息
#### 请求地址
| GET方法                                    |
| ---------------------------------------- |
| http://HOST/repositories/manifests?repo_name=myweb&tag=1.1 |
#### 请求参数
| 名称        | 类型     | 是否必须 | 描述   |
| --------- | ------ | ---- | ---- |
| repo_name | string | 是    | 项目名称 |
| tag       | string | 是    | 镜像标签 |

#### 返回参数
| 名称   | 类型     | 示例值  | 描述data     |
| ---- | ------ | ---- | ---------- |
| data | string |      | 返回配置页面HTML |

#### 请求示例
```php
curl 
-H "Content-type: application/json" \
-X POST 'http://HOST/repositories/manifests?repo_name=myweb&tag=1.1'
```
#### 响应示例
正常响应结果
```json
{
  "config":"{...}",
  "manifest":{}
}
```
#### 返回码解释

### 5.删除镜像标签
删除镜像标签
#### 请求地址
| DELETE方法                 |
| ------------------------ |
| http://HOST/repositories |
#### 请求参数
| 名称          | 类型     | 是否必须 | 描述   |
| ----------- | ------ | ---- | ---- |
| repoName    | string | 是    | 标签   |
| projectName | string | 是    | 项目名称 |
#### 返回参数
| 名称   | 类型     | 示例值  | 描述data |
| ---- | ------ | ---- | ------ |
| code | int    |      | 返回码    |
| msg  | string |      | 返回信息   |
#### 请求示例
```php
curl 
-H "Content-type: application/json" \
-X POST 'http://HOST/repositories' \
-d 'method=DELETE' \
-d '{"repoName":"1.1","projectName":"myweb"}'  
```
#### 响应示例
正常响应结果
```json
{
  "code":0,
  "msg":""
}
```
#### 返回码解释
