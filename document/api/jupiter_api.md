# 	多云对接API

## 机型模板接口 
### 1.获取模板列表

获取用户创建的所有模板列表

#### 请求地址

| GET方法                  |
| ---------------------- |
| http://HOST/v1/cluster |

#### 请求参数
| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| 无    | 无    | 否    | 无    |

#### 返回参数
| 名称      | 类型       | 示例值     | 描述   |
| :------ | :------- | :------ | :--- |
| code    | int      | 0       | 返回码  |
| msg     | string   | null    | 错误信息 |
| content | object[] | [{},{}] | 正常信息 |
| ext     | string   | null    | 额外信息 |
 content中数据的重要参数

| 名称                 | 类型     | 示例值                                     | 描述    |
| ------------------ | ------ | --------------------------------------- | ----- |
| Name               | string | "test"                                  | 模板名称  |
| Provider           | string | "aliyun"                                | 云厂商   |
| Cpu                | int    | 16                                      | CPU个数 |
| Ram                | int    | 64                                      | 内存容量  |
| InstanceType       | string | ecs.n2.3xlarge                          | 实例类型  |
| ImageId            | string | centos7u2_64_40G_cloudinit_20160728.raw | 镜像Id  |
| Network            | object | {}                                      | 网络接口  |
| Zone               | object | {}                                      | 区域    |
| Replication        | string | null                                    | 副本    |
| SystemDiskCategory | string | cloud_efficiency                        | 系统盘类型 |
| DataDiskSize       | int    | 1024                                    | 数据盘容量 |
| DataDiskNum        | int    | 2                                       | 数据盘数量 |
| DataDiskCategory   | int    | cloud_efficiency                        | 数据盘类型 |

#### 请求示例

```php
curl -X GET "http://HOST/v1/cluster"  \
-H "Content-type: application/json"   
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": [
    {
      "Id": 16,
      "Name": "test2",
      "Provider": "aliyun",
      "LastestPartNum": 0,
      "Desc": "",
      "CreateTime": "2017-07-07T16:10:46+08:00",
      "DeleteTime": "0001-01-01T00:00:00Z",
      "Cpu": 16,
      "Ram": 64,
      "InstanceType": "ecs.n2.3xlarge",
      "ImageId": "centos7u2_64_40G_cloudinit_20160728.raw",
      "PostScript": "",
      "KeyName": "",
      "Network": {
        "Id": 3,
        "VpcId": "",
        "SubnetId": "",
        "SecurityGroup": "",
        "InternetChargeType": "PayByBandwidth",
        "InternetMaxBandwidthOut": 5
      },
      "Zone": {
        "Id": 3,
        "RegionName": "cn-beijing-test",
        "ZoneName": "cn-beijing-test-3"
      },
      "Replication": null,
      "SystemDiskCategory": "cloud_efficiency",
      "DataDiskSize": 1024,
      "DataDiskNum": 2,
      "DataDiskCategory": "cloud_efficiency",
      "FlavorId": ""
    },
  ],
  "ext": null
}
```

异常返回结果：

```json
{
  "code": 10999,
  "msg": "QuerySeter no row found",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述       |
| ----- | --------------- | -------- |
| 0     | SERVICE_SUCCESS | 执行成功     |
| 10999 | SERVICE_ERRROR  | 获取模板信息失败 |

### 2.获取模板信息

获取用户创建的某个模板的信息

#### 请求地址

| GET方法                             |
| --------------------------------- |
| http://HOST/v1/cluster/:clusterId |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述   |
| --------- | ---- | ---- | ---- |
| clusterId | int  | 是    | 模板ID |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

 content中数据的重要参数同**模板列表 **

#### 请求示例

```php
curl -X GET "http://HOST/v1/cluster/3"  \
-H "Content-type: application/json"   
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": {
    "Id": 3,
    "Name": "1Core1G经典网",
    "Provider": "aliyun",
    "LastestPartNum": 0,
    "Desc": "",
    "CreateTime": "2017-07-05T10:01:48+08:00",
    "DeleteTime": "0001-01-01T00:00:00Z",
    "Cpu": 1,
    "Ram": 1,
    "InstanceType": "ecs.n1.tiny",
    "ImageId": "centos7u2_64_40G_cloudinit_20160728.raw",
    "PostScript": "",
    "KeyName": "key",
    "Network": {
      "Id": 1,
      "VpcId": "",
      "SubnetId": "",
      "SecurityGroup": "",
      "InternetChargeType": "PayByBandwidth",
      "InternetMaxBandwidthOut": 5
    },
    "Zone": {
      "Id": 1,
      "RegionName": "cn-beijing",
      "ZoneName": "cn-beijing-c"
    },
    "Replication": null,
    "SystemDiskCategory": "cloud_efficiency",
    "DataDiskSize": 100,
    "DataDiskNum": 1,
    "DataDiskCategory": "cloud_efficiency",
    "FlavorId": ""
  },
  "ext": null
}
```

异常返回结果： 

  获取模板信息失败

```json
{
  "code": 10999,
  "msg": "QuerySeter no row found",
  "content": null,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述       |
| ----- | --------------- | -------- |
| 0     | SERVICE_SUCCESS | 执行成功     |
| 10999 | SERVICE_ERROR   | 获取模板信息失败 |
| 12966 | BAD_REQUEST     | 请求参数错误   |

### 3.创建模板

创建机型模板

#### 请求地址

| POST方法                 |
| ---------------------- |
| http://HOST/v1/cluster |

#### 请求参数

| 名称                 | 类型     | 是否必须 | 描述    |
| ------------------ | ------ | ---- | ----- |
| Name               | string | 是    | 模板名称  |
| Provider           | string | 是    | 云厂商   |
| InstanceType       | string | 是    | 实例类型  |
| SystemDiskCategory | string | 是    | 系统盘类型 |
| DataDiskCategory   | string | 是    | 数据盘类型 |
| Zone               | string | 是    | 区域    |
| Network            | string | 是    | 网络    |
| DataDiskNum        | int    | 是    | 数据盘容量 |
| DataDiskSize       | int    | 是    | 数据盘大小 |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |


#### 请求示例

```bash
curl -X POST "http://HOST/v1/cluster"  \
-H "Content-type: application/json"  \
-d '{         
  "Id": 0, 
  "Name": "1Core1G经典网", 
  "Provider": "aliyun",  
  "Cpu": 1,    
  "Ram": 2, 
  "InstanceType": "1核1G", 
  "ImageId": "centos7u2_64_40G_cloudinit_20160728.raw",  
  "Network": { 
    "Id": 0, 
    "VpcId": "", 
    "SubnetId": "", 
    "SecurityGroup": "", 
    "InternetChargeType": "PayByBandwidth", 
    "InternetMaxBandwidthOut": 5 
  }, 
  "Zone": { 
    "Id": 0, 
    "RegionName": "cn-beijing", 
    "ZoneName": "cn-beijing-c" 
  }, 
  "Replication": null, 
  "SystemDiskCategory": "cloud_efficiency", 
  "DataDiskSize": 100, 
  "DataDiskNum": 1, 
  "DataDiskCategory": "cloud_efficiency", 
  "FlavorId": "" 
}'                                                       
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": 19,
  "ext": null
}
```

异常返回结果： 

创建模板失败

```json
{
  "code": 10999,
  "msg": "Ceate cluster err",
  "content": null,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 11996,
  "msg": "The paramter (DataDiskNum) over size limited and the size range is larger than 0 and less equal 4.!",
  "content": null,
  "ext": null
}
```

#### 返回码解释


| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 创建模板失败 |
| 11996 | BAD_REQUEST     | 请求参数错误 |
| 12996 | BAD_REQUEST     | 请求参数错误 |

### 4.删除模板

删除用户创建的机型模板

#### 请求地址

| DELETE方法                          |
| --------------------------------- |
| http://HOST/v1/cluster/:clusterId |

#### 请求参数

| 名称        | 类型     | 是否必须 | 描述   |
| --------- | ------ | ---- | ---- |
| clusterId | string | 是    | 模板ID |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```php
curl -X DELETE "http://HOST/v1/cluster/3"  \
-H "Content-type: application/json"   
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": true,
  "ext": null
}
```

异常返回结果： 

删除模板失败

```json
{
  "code": 10999,
  "msg": "Delete cluster err",
  "content": null,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释


| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 删除模板失败 |
| 12996 | BAD_REQUEST     | 请求参数错误 |

### 5.机器扩容

用户根据某一机型模板扩容N台机器

#### 请求地址

| POST方法                                   |
| ---------------------------------------- |
| http://HOST/v1/cluster/:clusterId/expand/:number |

#### 请求参数

| 名称                 | 类型     | 是否必须 | 描述          |
| ------------------ | ------ | ---- | ----------- |
| clusterId          | int    | 是    | 模板ID        |
| number             | int    | 是    | 机器数量        |
| X-CORRELATION-ID   | string | 是    | 关系ID，用于日志入库 |
| Name               | string | 是    | 模板名称        |
| Provider           | string | 是    | 云厂商         |
| InstanceType       | string | 是    | 实例类型        |
| SystemDiskCategory | string | 是    | 系统盘类型       |
| DataDiskCategory   | string | 是    | 数据盘类型       |
| Zone               | string | 是    | 区域          |
| Network            | string | 是    | 网络          |
| DataDiskNum        | int    | 是    | 数据盘容量       |
| DataDiskSize       | int    | 是    | 数据盘大小       |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X POST "http://HOST/v1/cluster/3"  \
-H "Content-type: application/json" \
-H "X-CORRELATION-ID:1-1"                              \ 
-d '{  
  "Name": "1Core1G经典网", 
  "Provider": "aliyun",  
  "Cpu": 0, 
  "Ram": 0, 
  "InstanceType": "1核1G", 
  "ImageId": "centos7u2_64_40G_cloudinit_20160728.raw",  
  "Network": { 
    "Id": 0, 
    "VpcId": "", 
    "SubnetId": "", 
    "SecurityGroup": "", 
    "InternetChargeType": "PayByBandwidth", 
    "InternetMaxBandwidthOut": 5 
  }, 
  "Zone": { 
    "Id": 0, 
    "RegionName": "cn-beijing", 
    "ZoneName": "cn-beijing-c" 
  }, 
  "Replication": null, 
  "SystemDiskCategory": "cloud_efficiency", 
  "DataDiskSize": 100, 
  "DataDiskNum": 1, 
  "DataDiskCategory": "cloud_efficiency", 
  "FlavorId": ""       
}'                                                        
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": [
    "i-2ze1v31lnkfgubznmmtz"
  ],
  "ext": 3
}
```

异常返回结果： 

扩容失败

```json
{
  "code": 10999,
  "msg": "QuerySeter no row found",
  "content": null,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释


| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 扩容失败   |
| 12996 | BAD_REQUEST     | 请求参数错误 |

### 6.机器数量

获取某一时间段内的在线机器数量信息

#### 请求地址

| GET方法                     |
| ---------------------- |
| http://HOST/v1/cluster/number/:hour

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| hour | int  | 是    | 时间范围 |

#### 返回参数

| 名称      | 类型       | 示例值     | 描述   |
| :------ | :------- | :------ | :--- |
| code    | int      | 0       | 返回码  |
| msg     | string   | null    | 错误信息 |
| content | object[] | [{},{}] | 正常信息 |
| ext     | string   | null    | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/cluster/number/3"  \
-H "Content-type: application/json"   
```

#### 响应示例

正常返回结果：

```json
 "code": 0,
  "msg": null,
  "content": [
    {
      "aliyun": "2",
      "phydev": "5",
      "time": "2017-08-02 17:52:29",
      "total": "7"
    },
    {
      "aliyun": "2",
      "phydev": "5",
      "time": "2017-08-02 18:00:00",
      "total": "7"
    }
  ],
  "ext": null
}
```

异常返回结果： 

获取机器信息失败

```json
{
  "code": 10999,
  "msg": "Get instance detail err",
  "content": null,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释


| 返回码   | 状态              | 描述       |
| ----- | --------------- | -------- |
| 0     | SERVICE_SUCCESS | 执行成功     |
| 10999 | SERVICE_ERROR   | 获取机器信息失败 |
| 12996 | BAD_REQUEST     | 请求参数错误   |

### 7.某时刻的机器数量

获取指定过去某一时间点的在线机器数量信息

#### 请求地址 

| GET方法                     |
| ---------------------- |
| http://HOST/v1/cluster/oldnumber/:time

#### 请求参数

| 名称   | 类型     | 是否必须 | 描述    |
| ---- | ------ | ---- | ----- |
| time | string | 是    | 某一时间点 |

#### 返回参数

| 名称      | 类型       | 示例值     | 描述   |
| :------ | :------- | :------ | :--- |
| code    | int      | 0       | 返回码  |
| msg     | string   | null    | 错误信息 |
| content | object[] | [{},{}] | 正常信息 |
| ext     | string   | null    | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/cluster/number/3"  \
-H "Content-type: application/json"   
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": {
    "aliyun": "2",
    "phydev": "5",
    "time": "2017-08-02 17:52:29",
    "total": "7"
  },
  "ext": null
}
```

异常返回结果： 

获取机器信息失败

```json
{
  "code": 10999,
  "msg": "Get instance detail at the time err",
  "content": null,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释


| 返回码   | 状态              | 描述       |
| ----- | --------------- | -------- |
| 0     | SERVICE_SUCCESS | 执行成功     |
| 10999 | SERVICE_ERROR   | 获取机器信息失败 |
| 12996 | BAD_REQUEST     | 请求参数错误   |

## 机器实例接口

### 1.创建机器  

 创建一个台机器实例

#### 请求地址

| POST方法                  |
| ----------------------- |
| http://HOST/v1/instance |

#### 请求参数

| 名称                 | 类型     | 是否必须 | 描述    |
| ------------------ | ------ | ---- | ----- |
| Provider           | string | 是    | 云厂商   |
| InstanceType       | string | 是    | 实例类型  |
| ImageId            | string | 是    | 镜像ID  |
| Network            | string | 是    | 网络    |
| Zone               | string | 是    | 区域    |
| SystemDiskCategory | string | 是    | 系统盘类型 |
| DataDiskSize       | int    | 是    | 数据盘容量 |
| DataDiskNum        | int    | 是    | 数据盘数量 |
| DataDiskCategory   | string | 是    | 数据盘类型 |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X POST "http://HOST/v1/instance"  \
-H "Content-type: application/json"     \
-d '{               
  "Provider": "aliyun", 
  "InstanceType": "ecs.c2.medium", 
  "ImageId": "centos7u2_64_40G_cloudinit_20160728.raw", 
  "Network": { 
    "Id": 0, 
    "VpcId": "", 
    "SubnetId": "", 
    "SecurityGroup": "", 
    "InternetChargeType": "PayByBandwidth", 
    "InternetMaxBandwidthOut": 5 
  }, 
  "Zone": { 
    "Id": 0, 
    "RegionName": "cn-beijing", 
    "ZoneName": "cn-beijing-c" 
  },                                           
  "SystemDiskCategory": "cloud_efficiency", 
  "DataDiskSize": 100, 
  "DataDiskNum": 1, 
  "DataDiskCategory": "cloud_efficiency", 
}'
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": "10.41.52.201",
  "ext": null
}
```

异常返回结果：  

请求参数错误

```json
{
  "code": 12996,
  "msg": "Input error",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 12996 | BAD_REQUEST     | 请求参数错误 |

### 2.机器信息

获取某个机器实例的信息

#### 请求地址

| GET方法                               |
| ----------------------------------- |
| http://HOST/v1/instance/:instanceId |

#### 请求参数

| 名称         | 类型     | 是否必须 | 描述   |
| ---------- | ------ | ---- | ---- |
| instanceId | string | 是    | 实例ID |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/i-2ze1v31lnkfgubznmmtz"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": {
    "Id": 20,
    "Cluster": {
      "Id": 4,
      "Name": "Physical device",
      "Provider": "phydev",
      "LastestPartNum": 0,
      "Desc": "About physical device",
      "CreateTime": "2017-07-05T10:14:06+08:00",
      "DeleteTime": "0001-01-01T00:00:00Z",
      "Cpu": 0,
      "Ram": 0,
      "InstanceType": "",
      "ImageId": "",
      "PostScript": "",
      "KeyName": "",
      "Network": {
        "Id": 2,
        "VpcId": "",
        "SubnetId": "",
        "SecurityGroup": "",
        "InternetChargeType": "",
        "InternetMaxBandwidthOut": 0
      },
      "Zone": {
        "Id": 2,
        "RegionName": "",
        "ZoneName": ""
      },
      "Replication": null,
      "SystemDiskCategory": "",
      "DataDiskSize": 0,
      "DataDiskNum": 0,
      "DataDiskCategory": "",
      "FlavorId": ""
    },
    "Provider": "phydev",
    "CreateTime": "2017-07-21T10:56:20+08:00",
    "Cpu": 1,
    "Ram": 1,
    "InstanceId": "i-b5omql5sajq2oum6srf0",
    "ImageId": "",
    "InstanceType": "",
    "KeyName": "",
    "VpcId": "",
    "SubnetId": "",
    "SecurityGroupId": "",
    "RegionId": "",
    "ZoneId": "",
    "DataDiskNum": 0,
    "DataDiskSize": 0,
    "DataDiskCategory": "",
    "SystemDiskCategory": "",
    "CostWay": "",
    "PreBuyMonth": 0,
    "PrivateIpAddress": "1.11.1.1",
    "PublicIpAddress": "101.1.2.42",
    "NatIpAddress": "",
    "Status": 3,
    "TenantID": "",
    "UserID": "",
    "Name": ""
  },
  "ext": null
}
```

异常返回结果：  

请求出错

```json
{
  "code": 10999,
  "msg": "get one instance err",
  "content": null,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 12997,
  "msg": "Missing requisite parameter!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 12997 | BAD_REQUEST     | 请求参数错误 |

### 3.启动机器

启动某个机器实例

#### 请求地址

| GET方法                            |
| -------------------------------- |
| http://HOST/v1/start/:instanceId |

#### 请求参数

| 名称         | 类型     | 是否必须 | 描述   |
| ---------- | ------ | ---- | ---- |
| instanceId | string | 是    | 实例ID |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/start/i-2ze1v31lnkfgubznmmtz"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": true,
  "ext": null
}
```

异常返回结果：  

启动机器失败

```json
{
  "code": 10999,
  "msg": "Could not start instances",
  "content": false,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 12997,
  "msg": "Missing requisite parameter!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 启动机器失败 |
| 12997 | BAD_REQUEST     | 请求参数错误 |

### 4.机器状态

查看某个机器的状态

#### 请求地址

| GET方法                             |
| --------------------------------- |
| http://HOST/v1/status/:instanceId |

#### 请求参数

| 名称         | 类型     | 是否必须 | 描述   |
| ---------- | ------ | ---- | ---- |
| instanceId | string | 是    | 实例ID |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/start/i-2ze1v31lnkfgubznmmtz"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": [
    {
      "instance_id": "i-fsfsfasfaflflsj23",
      "status": 7,
      "ip_address": ""
    }
  ],
  "ext": null
}
```

异常返回结果：  

查看机器状态失败

```json
{
  "code": 10999,
  "msg": "get multi instance err:",
  "content": false,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 12997,
  "msg": "Missing requisite parameter!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述       |
| ----- | --------------- | -------- |
| 0     | SERVICE_SUCCESS | 执行成功     |
| 10999 | SERVICE_ERROR   | 查看机器状态失败 |
| 12997 | BAD_REQUEST     | 请求参数错误   |

### 5.更新机器状态

更新某个机器实例的状态

#### 请求地址

| POST方法                |
| --------------------- |
| http://HOST/v1/status |

#### 请求参数

| 名称         | 类型     | 是否必须 | 描述   |
| ---------- | ------ | ---- | ---- |
| InstanceId | string | 是    | 实例ID |
| Status     | int    | 是    | 状态值  |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X POST "http://HOST/v1/status"  \
-H "Content-type: application/json"  \
-d '{ 
  "InstanceId": "i-b5omql5sajq2oum6srf0", 
  "Status": 1 
}'
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": 1,
  "ext": null
}
```

异常返回结果：  

请求出错

```json
{
  "code": 10999,
  "msg": "update instance status err",
  "content": false,
  "ext": null
}
```

请求参数错误

```json
{
  "code": 12996,
  "msg": "Input error",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 更新状态失败 |
| 12996 | BAD_REQUEST     | 请求参数错误 |

### 6. 删除机器 

删除用户创建和录入的实例

#### 请求地址

| DELETE方法                             |
| ------------------------------------ |
| http://HOST/v1/instance/:instanceIds |

#### 请求参数

| 名称               | 类型       | 是否必须 | 描述          |
| ---------------- | -------- | ---- | ----------- |
| instanceids      | string[] | 是    | 一组实例ID      |
| X-CORRELATION-ID | string   | 是    | 关系ID，用于日志入库 |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X DELETE "http://HOST/v1/instance/i-2ze1v31lnkfgubznmmtz"  \
-H "Content-type: application/json"  -H "X-CORRELATION-ID：1-1"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": null,
  "ext": null
}
```

异常返回结果：  

请求参数错误

```json
{
  "code": 12997,
  "msg": "Missing requisite parameter(X-CORRELATION-ID)!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 12997 | BAD_REQUEST     | 请求参数错误 |

### 7.下载密钥

下载机器的私钥

#### 请求地址

| GET方法                              |
| ---------------------------------- |
| http://HOST/v1/instance/sshkey/:ip |

#### 请求参数

| 名称   | 类型     | 是否必须 | 描述   |
| ---- | ------ | ---- | ---- |
| ip   | string | 是    | 机器IP |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/101.1.2.42"  \
-H "Content-type: application/json"   
```

#### 响应示例

正常返回结果：

```json
 开始下载密钥
```

异常返回结果：  

请求参数错误

```json
{
  "code": 12997,
  "msg": "QuerySeter no row found",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 12997 | BAD_REQUEST     | 请求参数错误 |

### 8.上传密钥

上传用户生成的密钥

#### 请求地址

| PUT方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/sshkey/:instanceId |

#### 请求参数

| 名称         | 类型     | 是否必须 | 描述   |
| ---------- | ------ | ---- | ---- |
| instanceId | string | 是    | 实例ID |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X PUT "http://HOST/v1/instance/sshkey/i-2ze1v31lnkfgubznmmtz"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": {密钥内容},
  "ext": null
}
```

异常返回结果：  

请求参数错误

```json
{
  "code": 12997,
  "msg": "QuerySeter no row found",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 12997 | BAD_REQUEST     | 请求参数错误 |

### 9.云厂商列表

返回所有云厂商的名称

#### 请求地址

| GET方法                            |
| -------------------------------- |
| http://HOST/v1/instance/provider |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| 无    | 无    | 否    | 无    |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/provider"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": [
    "aliyun",
    "aws",
    "openstack"
  ],
  "ext": null
}
```

异常返回结果：  

请求参数错误

```json
{
  "code": 12997,
  "msg": "QuerySeter no row found",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 12997 | BAD_REQUEST     | 请求参数错误 |

### 10.区域信息

获取某个云厂商的区域信息

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/regions/:provider |

#### 请求参数

| 名称       | 类型     | 是否必须 | 描述    |
| -------- | ------ | ---- | ----- |
| provider | string | 是    | 云厂商名称 |

#### 返回参数

| 名称      | 类型       | 示例值     | 描述   |
| :------ | :------- | :------ | :--- |
| code    | int      | 0       | 返回码  |
| msg     | string   | null    | 错误信息 |
| content | object[] | [{},{}] | 正常信息 |
| ext     | string   | null    | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/regions/aliyun"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": [ 
    {
      "Endpoint": "",
      "RegionName": "cn-beijing"
    },
    {
      "Endpoint": "",
      "RegionName": "cn-zhangjiakou"
    }
  ],
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 10999,
  "msg": "unknown backend provider driver: ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 获取信息失败 |

### 11.位置信息

获取某个云厂商的具体位置的信息


#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/zones/:provider/:regionId |

#### 请求参数

| 名称       | 类型     | 是否必须 | 描述    |
| -------- | ------ | ---- | ----- |
| provider | string | 是    | 云厂商名称 |
| regionId | string | 是    | 区域ID  |

#### 返回参数

| 名称      | 类型       | 示例值     | 描述   |
| :------ | :------- | :------ | :--- |
| code    | int      | 0       | 返回码  |
| msg     | string   | null    | 错误信息 |
| content | object[] | [{},{}] | 正常信息 |
| ext     | string   | null    | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/zones/aliyun/cn-beijing"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": [
    {
      "Messages": null,
      "RegionName": "华北 2 可用区 D",
      "State": "",
      "ZoneName": "cn-beijing-d"
    },
    {
      "Messages": null,
      "RegionName": "华北 2 可用区 A",
      "State": "",
      "ZoneName": "cn-beijing-a"
    }
  ],
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 10999,
  "msg": "unknown backend provider driver: ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 获取信息失败 |

### 12. VPC信息

获取某个云厂商某个区域的VPC信息

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/vpc/:provider/:regionId |

#### 请求参数

| 名称       | 类型     | 是否必须 | 描述    |
| -------- | ------ | ---- | ----- |
| provider | string | 是    | 云厂商名称 |
| regionId | string | 是    | 区域ID  |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/vpc/aliyun/cn-beijing"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": [
    {
      "CidrBlock": "172.17.0.0/16",
      "DhcpOptionsId": "",
      "InstanceTenancy": "",
      "IsDefault": false,
      "State": "Available",
      "Tags": null,
      "VpcId": "vpc-2ze0sufdt8yiguyg8u2lp"
    }     
  ],
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 10999,
  "msg": "unknown backend provider driver: ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 获取信息失败 |

### 13. 子网信息

获取某个云厂商某个区域的子网信息

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/subnet/:provider/:zoneId/:vpcId |

#### 请求参数

| 名称       | 类型     | 是否必须 | 描述     |
| -------- | ------ | ---- | ------ |
| provider | string | 是    | 云厂商名称  |
| zoneId   | string | 是    | 位置ID   |
| vpcID    | string | 是    | VPC ID |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/subnet/aliyun/cn-beijing-a/vpc-2ze0sufdt8yiguyg8u2lp"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": [
    {
      "AvailabilityZone": "cn-beijing-a",
      "AvailableIpAddressCount": 0,
      "CidrBlock": "172.17.192.0/20",
      "DefaultForAz": false,
      "MapPublicIpOnLaunch": false,
      "State": "Available",
      "SubnetId": "vsw-2zeci6ok1lzo3wjdgvtua",
      "Tags": null,
      "VpcId": "vpc-2ze0sufdt8yiguyg8u2lp"
    }
  ],
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 10999,
  "msg": "unknown backend provider driver: ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 获取信息失败 |

### 14.实例类型

获取云厂商机器实例的类型

#### 请求地址

| GET方法                                  |
| -------------------------------------- |
| http://HOST/v1/instance/type/:provider |

#### 请求参数

| 名称       | 类型     | 是否必须 | 描述    |
| -------- | ------ | ---- | ----- |
| provider | string | 是    | 云厂商名称 |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/type/aliyun"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": [
    "1Core-1GB",
    "1Core-2GB",
    "4Cores-8GB",
    "16Cores-16GB",
    "16Cores-64GB"
  ],
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 10999,
  "msg": "unknown backend provider driver: ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 获取信息失败 |

### 15.付费类型

获取云厂商机器的付费类型

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/charge/:provider |

#### 请求参数

| 名称       | 类型     | 是否必须 | 描述    |
| -------- | ------ | ---- | ----- |
| provider | string | 是    | 云厂商名称 |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/charge/aliyun"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
 {
  "code": 0,
  "msg": null,
  "content": [
    "PayByBandwidth",
    "PayByTraffic"
  ],
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 10999,
  "msg": "unknown backend provider driver: ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 获取信息失败 |

### 16.磁盘类型

查询云厂商机器的磁盘类型

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/disk_category/:provider |

#### 请求参数

| 名称       | 类型     | 是否必须 | 描述    |
| -------- | ------ | ---- | ----- |
| provider | string | 是    | 云厂商名称 |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/disk_category/aliyun"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": [
    "cloud_efficiency",
    "cloud_ssd"
  ],
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 10999,
  "msg": "unknown backend provider driver: ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 获取信息失败 |

### 17.镜像信息

获取公有云机器的镜像信息

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/image/:provider/:regionId |

#### 请求参数

| 名称       | 类型     | 是否必须 | 描述    |
| -------- | ------ | ---- | ----- |
| provider | string | 是    | 云厂商名称 |
| regionId | string | 是    | 区域ID  |

#### 返回参数

| 名称      | 类型       | 示例值     | 描述   |
| :------ | :------- | :------ | :--- |
| code    | int      | 0       | 返回码  |
| msg     | string   | null    | 错误信息 |
| content | object[] | [{},{}] | 正常信息 |
| ext     | string   | null    | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/iamge/aliyun/cn-beijing"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": [
    {
      "Architecture": "x86_64",
      "BlockDeviceMappings": null,
      "CreationDate": "2017-07-31T09:00:02Z",
      "Description": "升级docker到1.12\n添加mesos-slave",
      "EnaSupport": false,
      "Hypervisor": "",
      "ImageId": "m-2zefnwiq237sjwnt7oez",
      "ImageLocation": "",
      "ImageOwnerAlias": "",
      "ImageType": "",
      "KernelId": "",
      "Name": "mapi-image-2017073101",
      "OwnerId": "self",
      "Platform": "",
      "ProductCodes": [
        {
          "ProductCodeId": "",
          "ProductCodeType": ""
        }
      ],
      "Public": false,
      "RamdiskId": "",
      "RootDeviceName": "",
      "RootDeviceType": "",
      "SriovNetSupport": "",
      "State": "Available",
      "StateReason": {
        "Code": "",
        "Message": ""
      },
      "Tags": null,
      "VirtualizationType": ""
    }
  ],
  "ext": null
}
```

异常返回结果：  

请求参数错误

```json
{
  "code": 10999,
  "msg": "unknown backend provider driver: ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 获取信息失败 |

### 18.安全组

获取公有云机器的安全组信息

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/security_group/:provider/:regionId |

#### 请求参数

| 名称       | 类型     | 是否必须 | 描述    |
| -------- | ------ | ---- | ----- |
| provider | string | 是    | 云厂商名称 |
| regionId | string | 是    | 区域ID  |

#### 返回参数

| 名称      | 类型       | 示例值     | 描述   |
| :------ | :------- | :------ | :--- |
| code    | int      | 0       | 返回码  |
| msg     | string   | null    | 错误信息 |
| content | object[] | [{},{}] | 正常信息 |
| ext     | string   | null    | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/security_group/aliyun/cn-beijing"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": [
    {
      "Description": "P2P",
      "GroupId": "sg-2ze727lmif5l3o5x5vf5",
      "GroupName": "p2p_group",
      "IpPermissions": null,
      "IpPermissionsEgress": null,
      "OwnerId": "",
      "Tags": null,
      "VpcId": ""
    },
    {
      "Description": "System created security group.",
      "GroupId": "sg-25vpbkswc",
      "GroupName": "sg-25vpbkswc",
      "IpPermissions": null,
      "IpPermissionsEgress": null,
      "OwnerId": "",
      "Tags": null,
      "VpcId": ""
    }
  ],
  "ext": null
}
```

异常返回结果：  

请求参数错误

```json
{
  "code": 10999,
  "msg": "unknown backend provider driver: ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 10999 | SERVICE_ERROR   | 获取信息失败 |

### 19.机器列表

获取所有正在运行中机器实例

#### 请求地址

| GET方法                        |
| ---------------------------- |
| http://HOST/v1/instance/list |

#### 请求参数

| 名称   | 类型   | 是否必须 | 描述   |
| ---- | ---- | ---- | ---- |
| 无    | 无    | 否    | 无    |

#### 返回参数

| 名称      | 类型       | 示例值     | 描述   |
| :------ | :------- | :------ | :--- |
| code    | int      | 0       | 返回码  |
| msg     | string   | null    | 错误信息 |
| content | object[] | [{},{}] | 正常信息 |
| ext     | string   | null    | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/list"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": [
    {
      "Id": 89,
      "Cluster": {
        "Id": 4,
        "Name": "1Core-1Gib",
        "Provider": "aws",
        "LastestPartNum": 0,
        "Desc": "",
        "CreateTime": "2017-07-27T13:38:08Z",
        "DeleteTime": "0001-01-01T00:00:00Z",
        "Cpu": 1,
        "Ram": 1,
        "InstanceType": "t2.micro",
        "ImageId": "ami-3965b454",
        "PostScript": "",
        "KeyName": "zhaowei9",
        "Network": {
          "Id": 2,
          "VpcId": "vpc-62a57d06",
          "SubnetId": "subnet-f303e897",
          "SecurityGroup": "",
          "InternetChargeType": "PayByBandwidth",
          "InternetMaxBandwidthOut": 5
        },
        "Zone": {
          "Id": 3,
          "RegionName": "cn-north",
          "ZoneName": "cn-north-1"
        },
        "Replication": null,
        "SystemDiskCategory": "standard",
        "DataDiskSize": 100,
        "DataDiskNum": 1,
        "DataDiskCategory": "standard",
        "FlavorId": ""
    },
      "Provider": "aws",
      "CreateTime": "2017-08-01T07:28:22Z",
      "Cpu": 1,
      "Ram": 1,
      "InstanceId": "i-0b2fbd8b652431b62",
      "ImageId": "ami-3965b454",
      "InstanceType": "t2.micro",
      "KeyName": "",
      "VpcId": "vpc-b7b907d2",
      "SubnetId": "subnet-f303e897",
      "SecurityGroupId": "sg-d6d62bb3",
      "RegionId": "aws region ",
      "ZoneId": "aws zone",
      "DataDiskNum": 1,
      "DataDiskSize": 100,
      "DataDiskCategory": "standard",
      "SystemDiskCategory": "standard",
      "CostWay": "aws costway",
      "PreBuyMonth": 0,
      "PrivateIpAddress": "172.31.0.233",
      "PublicIpAddress": "52.80.61.125",
      "NatIpAddress": "aws nia",
      "Status": 7,
      "TenantID": "",
      "UserID": "",
      "Name": ""
    }
  ],
  "ext": null
}
```

异常返回结果：  

获取机器列表失败

```json
{
  "code": 10999,
  "msg": "get all instances error ",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述       |
| ----- | --------------- | -------- |
| 0     | SERVICE_SUCCESS | 执行成功     |
| 10999 | SERVICE_ERROR   | 获取机器列表失败 |

### 20.机型实例

获取某个机型模板的所有机器实例

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/cluster/:clusterId |

#### 请求参数

| 名称        | 类型   | 是否必须 | 描述     |
| --------- | ---- | ---- | ------ |
| clusterId | int  | 是    | 机型模板ID |

#### 返回参数

| 名称      | 类型       | 示例值     | 描述   |
| :------ | :------- | :------ | :--- |
| code    | int      | 0       | 返回码  |
| msg     | string   | null    | 错误信息 |
| content | object[] | [{},{}] | 正常信息 |
| ext     | string   | null    | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/cluster/3"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": [
    {
      "Id": 84,
      "Cluster": {
        "Id": 3,
        "Name": "1Core1G经典网",
        "Provider": "aliyun",
        "LastestPartNum": 0,
        "Desc": "",
        "CreateTime": "2017-07-27T13:38:08Z",
        "DeleteTime": "0001-01-01T00:00:00Z",
        "Cpu": 1,
        "Ram": 1,
        "InstanceType": "ecs.n1.tiny",
        "ImageId": "centos7u2_64_40G_cloudinit_20160728.raw",
        "PostScript": "",
        "KeyName": "key",
        "Network": {
          "Id": 1,
          "VpcId": "",
          "SubnetId": "",
          "SecurityGroup": "",
          "InternetChargeType": "PayByBandwidth",
          "InternetMaxBandwidthOut": 5
        },
        "Zone": {
          "Id": 1,
          "RegionName": "cn-beijing",
          "ZoneName": "cn-beijing-c"
        },
        "Replication": null,
        "SystemDiskCategory": "cloud_efficiency",
        "DataDiskSize": 100,
        "DataDiskNum": 1,
        "DataDiskCategory": "cloud_efficiency",
        "FlavorId": ""
      },
      "Provider": "aliyun",
      "CreateTime": "2017-08-01T07:18:38Z",
      "Cpu": 1,
      "Ram": 1,
      "InstanceId": "i-2zeeptz3hom3c5c9xl5s",
      "ImageId": "centos7u2_64_40G_cloudinit_20160728.raw",
      "InstanceType": "ecs.n1.tiny",
      "KeyName": "",
      "VpcId": "",
      "SubnetId": "",
      "SecurityGroupId": "sg-25vpbkswc",
      "RegionId": "cn-beijing",
      "ZoneId": "cn-beijing-c",
      "DataDiskNum": 1,
      "DataDiskSize": 100,
      "DataDiskCategory": "cloud_efficiency",
      "SystemDiskCategory": "cloud_efficiency",
      "CostWay": "PostPaid",
      "PreBuyMonth": 0,
      "PrivateIpAddress": "",
      "PublicIpAddress": "101.201.227.192",
      "NatIpAddress": "",
      "Status": 1,
      "TenantID": "",
      "UserID": "",
      "Name": ""
    }
  ],
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 10999,
  "msg": "get all instances error ",
  "content": null,
  "ext": null
}
```
请求参数出错

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述       |
| ----- | --------------- | -------- |
| 0     | SERVICE_SUCCESS | 执行成功     |
| 10999 | SERVICE_ERROR   | 获取信息失败   |
| 12996 | BAD_REQUEST     | 获取机器列表失败 |

### 21. 获取日志

根据correlationId 和 instanceId获取日志

#### 请求地址

| GET方法                                    |
| ---------------------------------------- |
| http://HOST/v1/instance/log/:correlationId/:instanceId |

#### 请求参数

| 名称            | 类型     | 是否必须 | 描述              |
| ------------- | ------ | ---- | --------------- |
| correlationId | string | 是    | 关系ID，用于从数据库过滤日志 |
| instanceId    | string | 是    | 机器实例ID          |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/log/20-20/i-2zeeptz3hom3c5c"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": {日志信息},
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 0,
  "msg": "result to json error",
  "content": null,
  "ext": null
}
```
请求参数出错

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 0     | BAD_REQUEST     | 获取信息失败 |
| 12996 | BAD_REQUEST     | 请求参数出错 |

### 22. 获取日志

根据instanceId获取日志

#### 请求地址

| GET方法                                   |
| --------------------------------------- |
| http://HOST/v1/instance/log/:instanceId |

#### 请求参数

| 名称         | 类型     | 是否必须 | 描述     |
| ---------- | ------ | ---- | ------ |
| instanceId | string | 是    | 机器实例ID |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X GET "http://HOST/v1/instance/log/i-2zeeptz3hom3c5c"  \
-H "Content-type: application/json"  
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": {日志信息},
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 0,
  "msg": "result to json error",
  "content": null,
  "ext": null
}
```

请求参数出错

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 0     | BAD_REQUEST     | 获取信息失败 |
| 12996 | BAD_REQUEST     | 请求参数出错 |

### 23.上传机器信息 

上传机器信息到数据库 

#### 请求地址

| PUT方法                          |
| ------------------------------ |
| http://HOST/v1/instance/phydev |

#### 请求参数

| 名称                 | 类型     | 是否必须 | 描述    |
| ------------------ | ------ | ---- | ----- |
| Provider           | string | 是    | 云厂商   |
| InstanceType       | string | 是    | 实例类型  |
| ImageId            | string | 是    | 镜像ID  |
| Network            | string | 是    | 网络    |
| Zone               | string | 是    | 区域    |
| SystemDiskCategory | string | 是    | 系统盘类型 |
| DataDiskSize       | int    | 是    | 数据盘容量 |
| DataDiskNum        | int    | 是    | 数据盘数量 |
| DataDiskCategory   | string | 是    | 数据盘类型 |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X PUT "http://HOST/v1/instance/phydev"  \
-H "Content-type: application/json"  \
-d '{                
  "Provider":"aliyun",                                      
  "InstanceType":"ecs.n2.3xlarge"     
  "ImageId":"centos7u2_64_40G_cloudinit_20160728.raw",           
  "Network": { 
    "Id": 0, 
    "VpcId": "", 
    "SubnetId": "", 
    "SecurityGroup": "", 
    "InternetChargeType": "PayByBandwidth", 
    "InternetMaxBandwidthOut": 5 
  }, 
  "Zone": { 
    "Id": 0, 
    "RegionName": "cn-beijing", 
    "ZoneName": "cn-beijing-c" 
  }, 
  "SystemDiskCategory": "cloud_efficiency",           
  "DataDiskSize": 100, 
  "DataDiskNum": 1,   
  "DataDiskCategory": "cloud_efficiency",                    
}' 
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": {机器信息},
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 0,
  "msg": "result to json error",
  "content": null,
  "ext": null
}
```
请求参数出错

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 0     | BAD_REQUEST     | 获取信息失败 |
| 12996 | BAD_REQUEST     | 请求参数出错 |

### 24.录入机器

录入用户自己的机器

#### 请求地址

| POST方法                         |
| ------------------------------ |
| http://HOST/v1/instance/phydev |

#### 请求参数

| 名称               | 类型     | 是否必须 | 描述          |
| ---------------- | ------ | ---- | ----------- |
| PublicIP         | string | 是    | 公网IP        |
| PrivateIP        | string | 是    | 内网IP        |
| Password         | string | 是    | 密码          |
| X-CORRELATION-ID | string | 是    | 关系ID，用于日志入库 |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X POST "http://HOST/v1/instance/phydev"  \
-H "Content-type: application/json"  -H "X-CORRELATION-ID：1-1" \
-d '{  
  "instancelist": [ 
    { 
      "PublicIp": "101.1.2.146", 
      "PrivateIp": "1.11.1.123", 
      "Password": "adminpadd", 
      "Port": 0 
    } 
  ] 
}'
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": {
    "success": 1,
    "failed": 0,
    "errors": []
  },
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 0,
  "msg": null,
  "content": {
    "success": 0,
    "failed": 1,
    "errors": [
      "Instance: 101.1.2.46 is already in DB"
    ]
  },
  "ext": null
}
```
请求参数出错

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

请求参数出错

```json
{
  "code": 12997,
  "msg": "Missing requisite parameter(X-CORRELATION-ID)!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 12996 | BAD_REQUEST     | 获取信息失败 |
| 12997 | BAD_REQUEST     | 请求参数出错 |

### 25.OpenStack配置

修改OpenStack配置信息

#### 请求地址

| POST方法                            |
| --------------------------------- |
| http://HOST/v1/instance/openstack |

#### 请求参数

| 名称         | 类型     | 是否必须 | 描述   |
| ---------- | ------ | ---- | ---- |
| OpIP       | string | 是    | 机器IP |
| OpPort     | string | 是    | 机器端口 |
| OpUserName | string | 是    | 用户名  |
| OpPassword | string | 是    | 密码   |

#### 返回参数

| 名称      | 类型     | 示例值  | 描述   |
| :------ | :----- | :--- | :--- |
| code    | int    | 0    | 返回码  |
| msg     | string | null | 错误信息 |
| content | object | {}   | 正常信息 |
| ext     | string | null | 额外信息 |

#### 请求示例

```bash
curl -X POST "http://HOST/v1/instance/openstack"  \
-H "Content-type: application/json"  \
-d '{
  "OpIP":"127.0.0.1",
  "OpPort":"26018",
  "OpUserName":"root",
  "OpPassword":"test12345"
}'
```

#### 响应示例

正常返回结果：

```json
{
  "code": 0,
  "msg": null,
  "content": {配置信息}，
  "ext": null
}
```

异常返回结果：  

获取信息失败

```json
{
  "code": 0,
  "msg": null,
  "content": {
    "success": 0,
    "failed": 1,
    "errors": [
      "Instance: 101.1.2.46 is already in DB"
    ]
  },
  "ext": null
}
```
请求参数出错

```json
{
  "code": 12996,
  "msg": "Input error!",
  "content": null,
  "ext": null
}
```

#### 返回码解释

| 返回码   | 状态              | 描述     |
| ----- | --------------- | ------ |
| 0     | SERVICE_SUCCESS | 执行成功   |
| 12996 | BAD_REQUEST     | 请求参数出错 |

