**OpenDCP**  

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://github.com/weibocom/opendcp/LICENSE)

**概览**

OpenDCP是一个基于Docker容器的云服务运维平台，集镜像仓库、多云支持、服务编排、服务发现等功能与一身，支持服务池的扩缩容，其体系技术源于微博用于支持节假日及热点事件期间高峰访问请求的DCP系统。OpenDCP允许您在不够买一台服务器的情况下，利用公有云服务器搭建起适应互联网应用的IT基础设置，并且将运维的工作量降到最低。

**特性**

-   基于Docker容器技术进行封装，不要求使用者掌握Docker。

-   支持阿里云、AWS等多个公有云平台，可根据需要随时申请和释放服务器，显著降低运营成本

-   分部门记录使用时长，资源用量一目了然

-   涵盖虚拟机创建、镜像打包、镜像部署、服务发现各个环节，简单易用，易于组织内推广

**快速开始**

以下介绍了系统快速部署方法，系统的详细使用请参照使用手册。

1.  安装 **Docker-1.10.0**和**Docker-Compose-1.6.0**以上版本, 详细请查看 [Docker
    Compose Install](https://docs.docker.com/compose/install/).  
    并确保docker-daemon已正常启动.

2.  下载源码，`git clone http://github.com/weibo/opendcp.git`
	>由于项目会下载所需的必需基础镜像,建议将下载源码放到空间大于50G以上的目录中。    
    
3. Harbor作为系统的一个模块.用于存储用户自定义的镜像,安装Harbor
    1.  `cd opendcp/deploy`.

    2.  使用`sudo`运行`scripts/installHarbor.sh <镜像仓库机器IP> 12380 /data1 aliyun` 以部署Harbor(镜像市场)，Harbor可以在单独的一台Linux机器上运行.脚本中会通过wget下载harbor安装文件,目前下载文件为harbor-online-installer-0.4.5.tgz,如果网络不佳,可将下载好的安装文件放到/data1/harbor目录下.然后执行`scripts/installHarbor.sh <镜像仓库机器IP> 12380 /data1 aliyun`即可。最后一个参数表示云服务,ali表示阿里云机器,从对应的仓库中拉取镜像,速度更快。如果需要停止harbor服务,需要进入到安装目录/data1/harbor/harbor下,执行`docker-compose down`。

4.  修改配置.
    * `cd opendcp/deploy`
    * 执行`./scripts/change_conf.sh`，按照提示修改配置文件。   
	* 或者按以下几个步骤手工修改：

        1.  编辑 conf/hubble_conf.php, 修改配置项 'HUBBLE_HOST','HUBBLE_ANSIBLE_HTTP' 和 'HUBBLE_SLB_HTTP' 为本机IP。
        2.  编辑 docker-compose.yml, 找到配置项 'imagebuild', 修改  'SERVER_IP' 为本机IP, 'HARBOR_ADDRESS'改为Harbor服务器的地址.
        3.  在conf/octans_roles/init/files/docker里，增加--insecure-registry 'harbor_ip:12380'。
        4.  编辑 conf/jupiter.json, 修改`KeyId`和`KeySecret`为您的阿里云账号和密码。
        5.  编辑 conf/web_conf.php, 修改  REPOS_DOMAIN 为 Harbor部署的IP。
 		6.  关于登录系统和认证
            * 当选择本地认证方式时，用户需要自行维护所有用户的登录账号。缺省只有管理员账号“root”能够登录，密码是“admin”。登录成功后，到导航栏 系统管理 -> 用户管理 可以增加管理用户。
			* 如果使用LDAP，需配置LDAP授权：修改deploy/conf/web_conf.php中“LDAP配置信息”一节。其中，“LDAP_HOST”是LDAP服务器的IP地址，“LDAP_PORT”是端口号，“LDAP_USER”是访问认证用的账号，“LDAP_PASS”是账号对应的密码，“LDAP_BIND”是LDAP绑定目录，“LDAP_SEARCH”是LDAP查找目录。首次使用可以忽略此步骤，登陆时采用本地登录不使用LDAP登录。

5.  构建镜像.
    1. `cd opendcp`
    2. 执行`./build.sh source [tag]`，构建各个模块的Docker镜像. 参数说明如下:
        - source 基础镜像来源,目前支持的来源有: aliyun, dockerio，国内用户推荐使用aliyun
        - tag   本次build命令执行输出的tag版本，默认不填为latest

       



6.  在`opendcp/deploy`目录下，运行 ./run.sh & 启动 OpenDCP, 访问网址
     [http://localhost:8888](http://localhost:8888/) 可以看到管理控制台界面。
    
7.  在`opendcp/deploy`目录下，运行 ./stop.sh 停止 OpenDCP。




**用户手册**

-   [Wiki](document/usermanual.md)

**作者**

-   Fu Wen

-   Fu Yuhui

-   Ke Yinan

-   Liu Peng

-   Ma Sihua

-   Sun Mingchao

-   Wang Guansheng

-   Wang Xiao

-   Yao Junxian

-   Zhang Yingeng

-   Zhuang Wenhui

**授权**

OpenDCP使用[Apache License
2.0](http://www.apache.org/licenses/LICENSE-2.0)授权协议进行授权
