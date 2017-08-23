locale_messages = {
    cloud:{
        "id": false,
        "Id": "序号",
        "Name": "名称",
        "Provider": "云厂商",
        "LastestPartNum": false,
        "PartOfInstances": false,
        "Desc": "描述",
        "CreateTime": "创建时间",
        "DeleteTime": false,
        "Cpu": "CPU",
        "Ram": "内存",
        "InstanceType": "实例类型",
        "ImageId": "镜像ID",
        "PostScript": false,
        "Network": "网络选项",
        "VpcId": "专有网络",
        "SubnetId": "子网",
        "SecurityGroup": "安全组",
        "Zone":"区域选项",
        "RegionName": "可用地域",
        "ZoneName": "可用区",
        "Replication": false,
        "SystemDiskCategory": "系统盘种类",
        "DataDiskNum": "数据盘数量",
        "DataDiskSize": "数据盘大小",
        "DataDiskCategory": "数据盘种类",
        "SecurityGroupId": "安全组",
        "InstanceId": "ECS序号",
        "KeyName": false,
        "RegionId": "可用地域",
        "ZoneId": "可用区域",
        "CostWay": "付费方式",
        "PrivateIpAddress" : "私有IP",
        "PublicIpAddress" : "公有IP",
        "NatIpAddress" : "转发IP",
        "Status" : false,
        "Cluster" : false,
        "PreBuyMonth" : "预付费购买月份",
        "RequestId" : false,
        "HostId" : false,
        "Code" : false,
        "Message" : false,
        "LoadBalancerId" : "SLB_Id",
        "RegionIdAlias" : "地域别名",
        "LoadBalancerName" : "SLB名称",
        "LoadBalancerStatus" : "状态",
        "Address" : "IP",
        "AddressType" : "地址类型",
        "NetworkType" : "网络类型",
        "VswitchId" : "子网",
        "InternetChargeType" : "付费类型",
        "Bandwidth" : "带宽",
        "ListenerPorts" : false,
        "ListenerPortsAndProtocal" : false,
        "ListenerPortsAndProtocol" : false,
        "BackendServers" : false,
        "ListenerPort" : "服务端口",
        "BackendServerPort" : "后端端口",
        //"XForwardedFor" : "XForwardedFor",
        "Scheduler" : "监听端口",
        "StickySession" : "回话保持",
        "StickySessionType" : "回话保持类型",
        "CookieTimeout" : "Cookie超时时间",
        //"Cookie" : "Cookie",
        "HealthCheck" : "监控检查",
        "HealthCheckDomain" : "检查域名",
        "HealthCheckURI" : "检查URI",
        "HealthyThreshold" : "健康阈值",
        "UnhealthyThreshold" : "不健康阈值",
        "HealthCheckTimeout" : "健康检查超时时间",
        "HealthCheckInterval" : "健康检查间隔",
        "HealthCheckHttpCode" : "HTTP状态码",
        "HealthCheckConnectPort" : "健康检查端口",
        "InternetMaxBandwidthOut" : "公网出带宽峰值",
    },
    package: {
        creator: '创建者',
        name: '项目名称',
        createTime: '创建时间',
        lastModifyTime: '修改时间',
        lastModifyOperator: '操作用户',
        Cluster: '隶属集群',
        DefineDockerFileType: '定义方式'
    },
    repos: {},
    layout: {
        role:{
            id: '序号',
            name: '名称',
            desc: '描述',
            tasks: 'task',
            vars: 'var',
            templates: 'template',
            user: '操作用户',
            create_time: '创建时间',
            update_time: '修改时间',
            role_file_path: '存取路径',
            handles:'handle',
            state: '任务状态',
            files: '文件',
            meta: 'meta'
        },
        resource:{
            id: '序号',
            name: '名称',
            desc: '描述',
            hidden: 'hidden',
            resource_content: '资源内容',
            resource_type: '资源类型',
            user: '最后修改者',
            create_time: '创建时间',
            update_time: '修改时间',
            template_file_owner:'文件owner',
            template_file_path:'文件路径',
            template_file_perm:'文件权限',
            type:'type',
            state: '任务状态',
            associate_role:'相关role',

        },
        id: '序号',
        name: '名称',
        desc: '描述',
        biz: false,
        service_type: '服务类型',
        docker_image: 'Docker镜像',
        cluster_id: '隶属集群序号',
        vm_type: '机型模板',
        sd_id: '服务发现序号',
        service_id: '隶属服务序号',
        tasks: '扩缩容任务',
        node_count: '节点数量',
        template_id: '任务模板序号',
        template_name: '任务模板名称',
        task_name: '任务名称',
        pool_name: '服务池名称',
        state: '任务状态',
        options: '任务参数',
        step_len: false,
        opr_user: '操作用户',
        created: '创建时间',
        updated: '修改时间',
        Stat: false,
        steps: '步骤',
        arg: '参数',
        params: '参数',
        param_values: '参数值',
        retry: '重试',
        actions: '命令'
    },
    hubble: {
        balance: {
            id: '序号',
            name: '名称',
            type: '服务发现类型',
            create_time: '创建时间',
            update_time: '修改时间',
            opr_user: '操作用户',
        },
        nginx: {
            id: '序号',
            name: '名称',
            create_time: '创建时间',
            update_time: '修改时间',
            opr_user: '操作用户',
            group_id: '隶属分组序号',
            deprecated: '已废弃(0:否)',
            is_consul: '已Consul化(0:否)',
            unit_id: '隶属单元序号',
            version: '当前版本',
            is_release: '已发布(0:否)',
            type: '发布类型',
            release_id: '发布序号'
        },
        shell: {
            id: '序号',
            name: '脚本名称',
            desc: '脚本描述',
            create_time: '创建时间',
            update_time: '修改时间',
            opr_user: '操作用户',
        },
        opr_log: {
            id: '序号',
            module: '模块',
            operation: '行为',
            opr_time: '操作时间',
            appkey: 'AppKey',
            user: '操作用户',
        }
    }
}