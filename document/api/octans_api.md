# 下发通道接口 

### 1.任务执行 

执行中心节点下发的任务

#### 请求地址

| POST方法              |
| ------------------- |
| http://HOST/api/run |

#### 请求参数

| 名称               | 类型       | 是否必须 | 描述          |
| ---------------- | -------- | ---- | ----------- |
| X-CORRELATION-ID | string   | 是    | 关系ID，用于日志入库 |
| X-RESOURCE       | string   | 是    | 请求来源        |
| nodes            | string[] | 是    | 节点IP        |
| user             | string   | 是    | 用户名         |
| name             | string   | 是    | 任务名         |
| tasks            | string   | 是    | 任务内容        |
| params           | string   | 是    | Ansible参数   |
| fork_num         | string   | 是    | Ansible参数   |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| content | object | {}   | 正常信息 |
| message | object | {}   | 异常信息 |

#### 请求示例

```bash
curl -X POST "http://HOST/api/run"  \
-H "Content-type: application/json"  -H "X-CORRELATION-ID:1-1" -H "X-RESOURCE:orion"\
-d '{
	"nodes":["192.201.11.31",],
	"user":"root",
	"name":"192.201.11.31_install_nginx_1501169594802406967",
	"tasks":"{具体任务}",
	"params":"{任务参数}"
	"fork_num":5,
}'
```

#### 响应示例

正常返回结果：

```json
{	
    "code":0,
    "content":{
        "id":2
    }    
}
```

异常返回结果：

任务重复

```json
{
    "code":-1
    "message":"task name is duplicate",
}
```

解析错误

```json
{   
    "code":-1,
    "message":"json encode error" 
}
```

参数错误

```json
{    
    "code":-1,
    "message":"param error, error:"
}
```

其它错误

```json
{  
    "code":-1,
    "message":"{other_err}" 
}
```



#### 返回码解释

| 返回码  | 状态      | 描述   |
| ---- | ------- | ---- |
| 0    | success | 执行成功 |
| -1   | failed  | 异常   |

### 2.任务停止

停止中心节点下发的任务  

#### 请求地址

| POST方法               |
| -------------------- |
| http://HOST/api/stop |

#### 请求参数

| 名称               | 类型     | 是否必须 | 描述          |
| ---------------- | ------ | ---- | ----------- |
| X-CORRELATION-ID | string | 是    | 关系ID，用于日志入库 |
| X-RESOURCE       | string | 是    | 请求来源        |
| id               | int    | 是    | 任务ID        |
| name             | string | 是    | 任务名称        |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| content | object | {}   | 正常信息 |
| message | object | {}   | 异常信息 |
#### 请求示例

```bash
curl -X POST "http://HOST/api/stop"  \
-H "Content-type: application/json"  -H "X-CORRELATION-ID:1-1" -H "X-RESOURCE:orion"\
-d '{
    "id":1,
    "name":"task1"
}'
```

#### 响应示例

正常返回结果：

```json
 {
    "code":0,
    "content":{}
}
```

异常返回结果：

查询任务失败

```json
{
    "code":-1,
    "message":"task not found..."
}
```

解析错误

```json
{   
    "code":-1,
    "message":"json encode error" 
}
```

其它错误

```json
{  
    "code":-1,
    "message":"{other_err}" 
}
```



#### 返回码解释

| 返回码  | 状态      | 描述   |
| ---- | ------- | ---- |
| 0    | success | 执行成功 |
| -1   | failed  | 异常   |

### 3.任务状态

检查任务的执行状态

#### 请求地址

| POST方法                |
| --------------------- |
| http://HOST/api/check |

#### 请求参数

| 名称               | 类型     | 是否必须 | 描述          |
| ---------------- | ------ | ---- | ----------- |
| X-CORRELATION-ID | string | 是    | 关系ID，用于日志入库 |
| X-RESOURCE       | string | 是    | 请求来源        |
| id               | int    | 是    | 任务ID        |
| name             | string | 是    | 任务名称        |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| content | object | {}   | 正常信息 |
| message | object | {}   | 异常信息 |

#### 请求示例

```bash
curl -X POST "http://HOST/api/check"  \
-H "Content-type: application/json"  -H "X-CORRELATION-ID:1-1" -H "X-RESOURCE:orion"\
-d '{
	 "id":1,
	 "name":"task1"
}'
```

#### 响应示例

正常返回结果：

```json
{
  	"code": 0,
  	"content": {
    	"create_time": "Wed, 17 Aug 2016 18:38:17 GMT",
    	"err": null,
    	"id": 7,
    	"log": "[{...},{...}]",
    	"nodes": "[\"123.207.136.186\"]",
    	"status": 4,
    	"update_time": "Wed, 17 Aug 2016 18:38:22 GMT"
  	}
}
```

异常返回结果：

查询任务失败

```json
{
	"code":-1,
	"message":"task not found..."
}
```

解析错误

```json
{   
    "code":-1,
    "message":"json encode error" 
}
```

其它错误

```json
{  
  	"code":-1,
    "message":"{other_err}" 
}
```



#### 返回码解释

| 返回码  | 状态      | 描述   |
| ---- | ------- | ---- |
| 0    | success | 执行成功 |
| -1   | failed  | 异常   |

### 4.任务日志

获取任务的的日志信息

#### 请求地址

| POST方法                 |
| ---------------------- |
| http://HOST/api/getlog |

#### 请求参数

| 名称               | 类型     | 是否必须 | 描述          |
| ---------------- | ------ | ---- | ----------- |
| X-CORRELATION-ID | string | 是    | 关系ID，用于日志入库 |
| X-RESOURCE       | string | 是    | 请求来源        |
| host             | string | 是    | 主机          |
| source           | string | 是    | 请求来源        |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| content | object | {}   | 正常信息 |
| message | object | {}   | 异常信息 |

#### 请求示例

```bash
curl -X POST "http://HOST/api/getlog"  \
-H "Content-type: application/json"  -H "X-CORRELATION-ID:1-1" -H "X-RESOURCE:orion"\
-d '{
	 "host":"127.0.0.1",
	 "source":"hubble"
}'
```

#### 响应示例

正常返回结果：

```json
["global_id = 16-16, source = hubble, create_time = 2017-07-27T16:37:04Z, end_time = 2017-07-27T16:37:05Z, host = 127.0.0.1, reuslt = ok",]
```

异常返回结果：

查询任务失败

```json
{
	"code":-1,
	"message":"task not found..."
}
```

解析错误

```json
{   
    "code":-1,
    "message":"json encode error" 
}
```

其它错误

```json
{  
  	"code":-1,
    "message":"{other_err}" 
}
```



#### 返回码解释

| 返回码  | 状态      | 描述   |
| ---- | ------- | ---- |
| 无    | success | 执行成功 |
| -1   | failed  | 异常   |

### 5.任务分发

中心节点分发任务给目标节点

#### 请求地址

| POST方法                       |
| ---------------------------- |
| http://HOST/api/parallel_run |

#### 请求参数

| 名称               | 类型       | 是否必须 | 描述          |
| ---------------- | -------- | ---- | ----------- |
| X-CORRELATION-ID | string   | 是    | 关系ID，用于日志入库 |
| X-RESOURCE       | string   | 是    | 请求来源        |
| nodes            | string[] | 是    | 节点IP        |
| user             | string   | 是    | 用户名         |
| name             | string   | 是    | 任务名         |
| tasks            | string   | 是    | 任务内容        |
| params           | string   | 是    | Ansible参数   |
| fork_num         | string   | 是    | Ansible参数   |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| content | object | {}   | 正常信息 |
| message | object | {}   | 异常信息 |

#### 请求示例


```bash
curl -X POST "http://HOST/api/run"  \
-H "Content-type: application/json"  -H "X-CORRELATION-ID:1-1" -H "X-RESOURCE:orion"\
-d '{
	"nodes":["192.201.11.31"],
	"user":"root",
	"name":"192.201.11.31_install_nginx_1501169594802406967",
	"content":"{具体任务}",
	"params":"{任务参数}"
	"fork_num":5,
}'
```

#### 响应示例

正常返回结果：

```json
["global_id = 16-16, source = hubble, create_time = 2017-07-27T16:37:04Z, end_time = 2017-07-27T16:37:05Z, host = 127.0.0.1, reuslt = ok",]
```

异常返回结果：

任务重复

```json
{
    "code":-1
    "message":"task name is duplicate",
}
```

解析错误

```json
{   
    "code":-1,
    "message":"json encode error" 
}
```

参数错误

```json
{    
    "code":-1,
    "message":"param error, error:"
}
```

其它错误

```json
{  
    "code":-1,
    "message":"{other_err}" 
}
```

#### 返回码解释

| 返回码  | 状态      | 描述   |
| ---- | ------- | ---- |
| 0    | success | 执行成功 |
| -1   | failed  | 异常   |
