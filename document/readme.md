**OpenDCP**

**概览**

OpenDCP是一个基于Docker容器的云服务运维平台，集镜像仓库、多云支持、服务编排、服务发现等功能与一身，支持业务服务的扩缩容，其体系技术源于微博用于支持节假日及热点事件期间高峰访问请求的DCP系统。OpenDCP允许您在不够买一台服务器的情况下，利用公有云服务器搭建起适应互联网应用的IT基础设置，并且将运维的工作量降到最低。

**特性**

-   基于Docker容器技术进行封装，不要求使用者掌握Docker。

-   支持阿里云、AWS等多个公有云平台，可根据需要随时申请和释放服务器，显著降低运营成本

-   分部门记录使用时长，资源用量一目了然

-   涵盖虚拟机创建、镜像打包、镜像部署、服务发现各个环节，简单易用，易于组织内推广

**快速开始**

以下介绍了系统快速部署方法，系统的详细使用请参照使用手册。

1.  按照 **Docker-1.10.0+** 和 **Docker-Compose-1.6.0+**, 详细请查看 [Docker
    Compose Install](https://docs.docker.com/compose/install/).

2.  下载**opendcp-0.1.0.tgz**, 解压到要部署的位置.

3.  修改配置.

    1.  切换到解压位置.

    2.  编辑 conf/hubble.conf, 修改配置项 'HUBBLE\_HOST' 和
        'HUBBLE\_ANSIBLE\_HTTP' 为本机IP.

    3.  Edit docker-compose.yml, find service 'imagebuild', and change
        environment 'SERVER\_IP' to the ip of the machine as above.

    4.  Edit conf/jupiter.json, and add your Aliyun key id and secrets there.

    5.  Edit conf/web\_conf.json, and change REPOS\_DOMAIN to the ip of Harbor.

4.  Run ./run.sh to start opendcp, and
    visit [http://localhost:8888](http://localhost:8888/) to see the web UI.

5.  Run scripts/installHarbor.sh using sudo on a Linux machine to deploy harbor.

6.  Run ./build.sh to build the project, which will build all sub projects into
    docker images.

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
