# OpenDCP 项目说明文档

## 一、项目简介

**OpenDCP** 是一个基于 Docker 的云资源管理与调度平台，集镜像仓库、多云支持、服务编排、服务发现等功能于一身，支持服务池的弹性扩缩容。其技术体系源于 微博 内部用于支持节假日及热点峰值流量的弹性调度 DCP 系统。

OpenDCP 允许利用公有云服务器快速搭建适应互联网应用的 IT 基础设施，并将运维工作量降至最低。

---

## 二、核心特性

- **基于 Docker 容器技术**：对应用进行容器化封装，使用者无需深入掌握 Docker 命令即可使用。
- **多云支持**：支持阿里云公有云平台，可根据业务负载随时申请和释放服务器，显著降低运营成本。
- **全链路覆盖**：涵盖虚拟机创建、镜像打包、镜像部署、服务发现等各个环节，简单易用。
- **弹性扩缩容**：支持服务池的快速扩容与缩容，轻松应对突发流量和峰值场景。

---

## 三、系统架构

OpenDCP 由以下核心模块组成：

| 模块 | 名称 | 功能说明 |
|------|------|----------|
| **orion** | 服务编排模块 | 负责部署和管理平台中的所有服务，支持服务池的扩容、缩容与上线。 |
| **octans** | 多云对接模块 | 负责向公有云（如阿里云）申请和创建服务器资源，并提供下发通道进行机器初始化。 |
| **imagebuild** | 镜像打包模块 | 帮助用户将业务代码打包成 Docker 镜像，并推送到镜像仓库。 |
| **hubble** | 服务注册及发现模块 | 维护可用资源列表，提供服务节点的注册与注销，支持 Nginx 和阿里云 SLB 等负载均衡。 |
| **harbor** | 私有镜像仓库 | 用于存储用户自定义的 Docker 镜像，作为系统镜像市场的后端仓库。 |
| **jupiter** | 多云对接网关 | 对外提供云资源管理的 API 接口，封装不同云厂商的驱动。 |
| **ui** | 前端管理界面 | 提供可视化的 Web 控制台，统一管理各模块功能。 |

### 架构示意图

```
┌─────────────────────────────────────────────────────────────┐
│                         用户访问层                           │
│              (Web UI / API / 负载均衡入口)                    │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                      OpenDCP 核心平台                        │
│  ┌─────────┐  ┌─────────┐  ┌─────────┐  ┌───────────────┐  │
│  │  Orion  │  │ Octans  │  │ Hubble  │  │  ImageBuild   │  │
│  │ 服务编排 │  │ 多云对接 │  │ 服务发现 │  │   镜像打包     │  │
│  └─────────┘  └─────────┘  └─────────┘  └───────────────┘  │
│  ┌─────────┐  ┌─────────┐  ┌─────────────────────────────┐  │
│  │ Jupiter │  │   UI    │  │         Harbor 仓库         │  │
│  │ 云网关  │  │ 管理界面 │  │      (Docker Registry)      │  │
│  └─────────┘  └─────────┘  └─────────────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                              │
┌─────────────────────────────────────────────────────────────┐
│                        基础设施层                            │
│         阿里云 ECS / 私有云 / 物理服务器 / 负载均衡            │
└─────────────────────────────────────────────────────────────┘
```

---

## 四、目录结构

```
opendcp/
├── orion/              # 服务编排模块
├── octans/             # 多云对接模块
├── jupiter/            # 云资源管理网关
├── hubble/             # 服务发现模块
├── imagebuild/         # 镜像打包模块
├── ui/                 # Web 前端界面
├── db_init/            # 数据库初始化脚本
├── deploy/             # 部署脚本与配置文件
│   ├── scripts/        # 安装/部署脚本
│   └── conf/           # 各模块配置文件
├── document/           # 项目文档
│   ├── install.md      # 安装手册
│   ├── quick_start.md  # 快速开始
│   ├── usermanual.md   # 用户手册
│   ├── api.md          # 开发接口文档
│   └── openstack.md    # 虚拟化安装手册
├── build.sh            # 构建脚本
├── push.sh             # 镜像推送脚本
└── README.md           # 项目简介
```

---

## 五、快速开始

### 5.1 环境要求

- Docker >= 1.10.0
- Docker Compose >= 1.6.0
- Git
- Linux 操作系统（推荐 CentOS 7+ / Ubuntu 16.04+）
- 磁盘空间 >= 50GB（用于存储镜像和日志）

### 5.2 一键部署

```bash
# 1. 克隆项目
git clone https://github.com/weibocom/opendcp.git
cd opendcp

# 2. 安装 Harbor（私有镜像仓库）
cd deploy/scripts
./installHarbor.sh <镜像仓库IP> 12380 /data1 aliyun

# 3. 修改配置
cd ../deploy
./scripts/change_conf.sh

# 4. 构建镜像
cd ../
./build.sh source aliyun

# 5. 启动服务
cd deploy
./run.sh -d

# 6. 访问管理界面
# http://<服务器IP>:8888
# 默认账号: root / 密码: admin
```

> 详细安装步骤请参考 [document/install.md](document/install.md)。

---

## 六、核心功能详解

### 6.1 多云对接

多云对接模块用于公有云资源的申请和创建。平台用到的所有服务器、负载均衡等资源均来自公有云。

- **机型模板**：预定义云服务器规格（CPU、内存、磁盘、镜像等），简化创建流程。
- **机器管理**：统一管理平台上所有服务器，支持创建、分配、回收。
- **配额管理**：为模板分配使用额度，按小时粒度计费。

### 6.2 镜像市场

服务部署的第一步是将业务代码打包成镜像。

- **打包系统**：支持直接编辑 Dockerfile、从 Git/SVN 拉取、或使用图形化工具定义镜像。
- **镜像仓库**：管理已构建的镜像，支持多集群（base、default_cluster、library）。
- **Base 镜像维护**：支持自定义基础镜像的上传与管理。

### 6.3 服务编排

系统的核心模块，负责部署和管理平台中的所有服务。

- **集群管理**：按部门或业务线划分资源集群。
- **服务管理**：定义服务名称、镜像来源、服务发现方式等。
- **服务池管理**：为服务配置运行环境（机型模板、扩容/缩容/上线模板）。
- **任务管理**：维护部署服务所用的任务模板，支持自定义执行流程。
- **远程命令**：使用 Ansible 脚本等方式实现命令的远程执行。

### 6.4 服务发现

维护可用资源列表，实现服务节点的自动注册与注销。

- **Nginx 服务发现**：支持七层负载均衡，提供主配置和 Upstream 的自动下发。
- **阿里云 SLB**：封装阿里云负载均衡服务，实现流量分发和故障转移。
- **自动发布**：通过 API 实现服务节点的自动注册与注销，与服务编排模块打通。

---

## 七、二次开发

### 7.1 多云对接 Driver 开发

OpenDCP 支持通过实现 `ProviderDriver` 接口扩展新的云厂商：

```go
type ProviderDriver interface {
    Create(cluster *models.Cluster, number int) (instanceIds []string, errs []error)
    Delete(instanceId string) (time.Time, error)
    Start(instanceId string) (bool, error)
    Stop(instanceId string) (bool, error)
    // ... 其他接口方法
}
```

实现后通过 `provider.RegisterProviderDriver` 注册即可。

### 7.2 API 接口文档

详细的二次开发接口说明请参考 [document/api.md](document/api.md)。

---

## 八、常见问题

### 8.1 Harbor 安装失败

如果遇到 `docker-compose` 版本或依赖问题，请升级：

```bash
sudo pip install backports.ssl_match_hostname --upgrade
```

### 8.2 打包失败

- 检查 Git/SVN 地址、用户名、密码是否正确。
- 检查 Dockerfile 中的 shell 命令是否合法。
- 确保磁盘空间充足。
- 查看日志：`opendcp/deploy/log/imagebuild/imagebuild.log`

### 8.3 服务发现配置

- 默认配置适用于 Nginx 服务发现类型，只需修改 Nginx IP 即可使用。
- 如需自定义，请创建新的分组、单元、Upstream 和主配置。

---

## 九、开源协议

OpenDCP 使用 [GPL-2.0](https://www.gnu.org/licenses/gpl-2.0.txt) 授权协议进行授权。

---

## 十、社区与交流

- **GitHub**: [https://github.com/weibocom/opendcp](https://github.com/weibocom/opendcp)
- **Gitter**: [https://gitter.im/weibocom/opendcp](https://gitter.im/weibocom/opendcp)
- **QQ 交流群**: 369485460

---

## 十一、贡献者

感谢以下贡献者对 OpenDCP 项目的支持：

- Fu Wen ([@itfuwen](https://github.com/itfuwen))
- Zhang Yingeng ([@goodidle](https://github.com/goodidle))
- Ma Sihua ([@masihua](https://github.com/masihua))
- Zhuang Wenhui ([@hacpai](https://github.com/hacpai))
- Wang Xiao ([@sean-w2011](https://github.com/sean-w2011))
- Yao Junxian ([@yaojunxian170](https://github.com/yaojunxian170))
- Peng Tao ([@PengTaoWW](https://github.com/PengTaoWW))
- He Jigui ([@hexuan1922](https://github.com/hexuan1922))
- Li Meng ([@MyEngIsPoor](https://github.com/MyEngIsPoor))
- Jiang Tao ([@jiangtao12](https://github.com/jiangtao12))
- 以及更多社区贡献者...

---

*文档版本: v1.0 | 最后更新: 2026-08-07*
