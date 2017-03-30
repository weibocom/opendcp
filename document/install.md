# 快速部署

## 概览
OPENDCP主要由以下几个模块构成:  
 - orion  服务编排模块  
 - octans  多云对接模块  
 - imagebuild 镜像打包模块  
 - hubble 服务注册及发现模块  
 - harbor 作为系统的私有仓库.用于存放客户自定义的docker镜像  

系统部署后将运行在docker容器内,所以系统将依赖docker服务.并会在本地打包生成镜像

## 安装步骤
以下介绍了系统快速部署方法，系统的详细使用请参照使用手册。**注意**安装Harbor，Docker Compose等组件的时候，由于国内网络环境的原因，访问github可能失败，遇此情况需要重试安装。

1.  安装 **Docker-1.10.0**和**Docker-Compose-1.6.0**以上版本, 详细请查看[Docker Install](https://www.docker.com/docker-centos-distribution) 和 [Docker
    Compose Install](https://docs.docker.com/compose/install/).  
    并确保docker-daemon已正常启动.

2.  确保**Git**已经安装，然后下载源码，`git clone http://github.com/weibocom/opendcp.git`
    - 由于项目会下载所需的必需基础镜像,建议将下载源码放到空间大于50G以上的目录中。 
	
3. 以下步骤请使用`root`用户或有sudo权限的用户`sudo su -`切换到`root`用户后执行。

3. Harbor作为系统的一个模块.用于存储用户自定义的镜像,安装Harbor
    - `cd opendcp/deploy/scripts` 进入到脚本目录
    - `./installHarbor.sh <镜像仓库机器IP> 12380 /data1 aliyun` 部署Harbor(镜像市场)，最后一个参数表示镜像获取方式,aliyun表示从阿里云获取,其他的从docker.io中获取,Harbor可以在单独的一台Linux机器上运行。
    - 如需停止harbor服务，`cd /data1/harbor/harbor` 进入到安装目录, 执行`docker-compose down`。

4.  修改配置
    - `cd opendcp/deploy`
    - 执行`./scripts/change_conf.sh`，按照提示修改配置文件。   
    - 或者按以下几个步骤手工修改：
        - <p>编辑 conf/hubble_conf.php, 修改配置项 'HUBBLE_HOST','HUBBLE_ANSIBLE_HTTP' 和 'HUBBLE_SLB_HTTP' 为本机IP。</p>
        - <p>编辑 docker-compose.yml, 找到配置项 'imagebuild', 修改  'SERVER_IP' 为本机IP, 'HARBOR_ADDRESS'改为Harbor服务器的地址.</p>
        - <p>在conf/octans_roles/init/files/docker里，增加--insecure-registry 'harbor_ip:12380'。</p>
        - <p>编辑 conf/jupiter.json, 修改`KeyId`和`KeySecret`为您的阿里云账号和密码。</p>
        - <p>编辑 conf/web_conf.php, 修改  REPOS_DOMAIN 为 Harbor部署的IP。</p>
        - <span id="auth">关于登录系统和认证</span>
            - 当选择本地认证方式时，用户需要自行维护所有用户的登录账号。缺省只有管理员账号“root”能够登录，密码是“admin”。登录成功后，到导航栏 系统管理 -> 用户管理 可以增加管理用户。
            - 如果使用LDAP，需配置LDAP授权：修改deploy/conf/web_conf.php中“LDAP配置信息”一节。其中，“LDAP_HOST”是LDAP服务器的IP地址，“LDAP_PORT”是端口号，“LDAP_USER”是访问认证用的账号，“LDAP_PASS”是账号对应的密码，“LDAP_BIND”是LDAP绑定目录，“LDAP_SEARCH”是LDAP查找目录。首次使用可以忽略此步骤，登陆时采用本地登录不使用LDAP登录。
5.  构建镜像.
    - `cd opendcp`
    - 执行`./build.sh source [tag]`，构建各个模块的Docker镜像. 参数说明如下:
        - source 基础镜像来源,目前支持的来源有: aliyun, dockerio，国内用户推荐使用aliyun</p>
        - <p>tag   本次build命令执行输出的tag版本，默认不填为latest</p>

6.  在`opendcp/deploy`目录下，运行 `./run.sh &` 启动 OpenDCP。 
    - 如果想要让OpenDCP在后台运行`./run.sh -d`。
    - 运行日志可在`logs`目录下查看:`tail -f logs/orion.log`。
    - OpenDCP启动完成后，访问网址
     [http://localhost:8888](http://localhost:8888/) 可以看到管理控制台界面。
7.  如要停止OpenDCP，在`opendcp/deploy`目录下，运行 ./stop.sh 。
