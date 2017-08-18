reveal = {
  for_cloud: [
    [
      '<h1>多云对接</h1>',
    ],
    [
      '<h2>首次使用注意事项</h2>' +
      '<ol>' +
      '<li>云厂商: 只支持阿里云</li>' +
      '<li>主要功能:</li>' +
      '<ul>' +
      '<li>机型模板的创建和删除</li>' +
      '<li>配额追加</li>' +
      '</ul>' +
      '<li>依赖关系:</li>' +
      '<ul>' +
      '<li>不依赖其它模块</li>' +
      '<li>被服务编排模块依赖</li>' +
      '</ul>' +
      '</ol>',
    ],
    [
      '<h2>机型模板</h2>' +
      '<ol>' +
      '<li>支持经典网络 和 专有网络</li>' +
      '<li>两种网络的机器仅根据机房内网络底层网络类型不同而做的划分</li>' +
      '<li>经典网络: IP由云厂商统一分配, 使用简便</li>' +
      '<li>专有网络: 逻辑隔离的私有网络</li>' +
      '<li>更多使用帮助: 请参见<a href="https://help.aliyun.com/product/25365.html" target="_blank">阿里云官方文档</a></li>' +
      '</ol>',

      '<h2>配额管理</h2>' +
      '<ol>' +
      '<li>配额以小时为单位</li>' +
      '<li>只可以追加配额, 不可以减少配额</li>' +
      '<li>配额仅在删除机器时结算, 支持透支</li>' +
      '<li>配额不足时, 影响范围:</li>' +
      '<ul>' +
      '<li>不能创建此机型的机器</li>' +
      '<li>服务编排模块内使用此机型的服务池扩容将失败</li>' +
      '</ul>' +
      '</ol>',
    ],
    [
      '<h2>机器管理</h2>' +
      '<ol>' +
      '<li>创建机器</li>' +
      '<ul>' +
      '<li>根据机型模板创建机器</li>' +
      '<li>配额不足时将创建失败</li>' +
      '<li>机器创建后将自动进行初始化</li>' +
      '<li>追加额外的初始化配置需要在后台增加响应的ROLE</li>' +
      '</ul>' +
      '<li>删除机器</li>' +
      '<ul>' +
      '<li>创建中的机器不允许删除</li>' +
      '<li>已在服务编排模块下使用的机器不允许删除</li>' +
      '<li>删除机器后将进行配额结算</li>' +
      '</ul>' +
      '</ol>',
    ],
    [
      '<h1>THANK YOU</h1>' +
      '<h1 style="color: orange;">YOU ARE THE ONE</h1>',
    ],
  ],
  for_repos: [
    [
      '<h1>镜像市场</h1>',
    ],
    [
      '<h2>首次使用注意事项</h2>' +
      '<ol>' +
      '<li class="text-danger">服务编排模块部署服务时依赖此模块</li>' +
      '<li>首先创建一个项目(服务)</li>' +
      '<li>然后构建项目镜像</li>' +
      '</ol>',
    ],
    [
      '<h2>镜像仓库子模块</h2>' +
      '<ol>' +
      '<li>打包系统</li>' +
      '<ul>' +
      '<li>目标是简化打包和上传镜像的操作</li>' +
      '<li>提供镜像构建的入口</li>' +
      '<li>配置时默认为直接编辑Dockerfile模式</li>' +
      '<ul>' +
      '<li>可选: 从Git下载、使用工具定义</li>' +
      '</ul>' +
      '</ul>' +
      '<li>镜像仓库</li>' +
      '<ul>' +
      '<li>采用开源项目Harbor</li>' +
      '<li>Harbor项目开源地址: <a href="https://github.com/vmware/harbor" target="_blank">跳转</a></li>' +
      '</ul>' +
      '</ol>',
    ],
    [
      '<h1>THANK YOU</h1>' +
      '<h1 style="color: orange;">YOU ARE THE ONE</h1>',
    ],
  ],
  for_layout: [
    [
      '<h1>服务编排</h1>',
    ],
    [
      '<h2>首次使用注意事项</h2>' +
      '<ol>' +
      '<li>所有数据的名称字段不可修改</li>' +
      '<li>关系: 集群 / 服务 / 服务池 / 节点 </li>' +
      '<li>任务模板依赖命令组, 命令组依赖命令</li>' +
      '</ol>',
    ],
    [
      '<h2>服务编排子模块</h2>' +
      '<ol>' +
      '<li>集群管理</li>' +
      '<li>服务管理</li>' +
      '<li>任务管理</li>' +
      '<li>远程命令和命令组</li>' +
      '</ol>',
    ],
    [
      '<h2>集群管理</h2>' +
      '<ul>' +
      '<li>为服务隔离而生</li>' +
      '<li>集群创建后, 名称不可修改</li>' +
      '</ul>',
    ],
    [
      '<h2>服务管理 (一) 服务</h2>' +
      '<ul>' +
      '<li>服务名称: 创建后不可修改</li>' +
      '<li>服务类型: 应用服务软件环境</li>' +
      '<li>镜像名称: </li>' +
      '<ul>' +
      '<li>服务池扩容依赖此配置</li>' +
      '<li>服务池上线会自动更新此配置</li>' +
      '<li class="text-danger">下属多个服务池使用不同镜像时, 扩容时需要检查此配置</li>' +
      '</ul>' +
      '</ul>',

      '<h2>服务管理 (二) 服务池</h2>' +
      '<ul>' +
      '<li>服务池名称: 创建后不可修改</li>' +
      '<li>机型模板: 依赖多云对接模块, 需要预先创建</li>' +
      '<li>服务发现类型: 依赖服务发现模块, 用于自动引流</li>' +
      '<li>扩容任务模板: 依赖任务管理, 需要预先创建</li>' +
      '<li>缩容任务模板: 依赖任务管理, 需要预先创建</li>' +
      '<li>上线任务模板: 依赖任务管理, 需要预先创建</li>' +
      '</ul>',

      '<h2>服务管理 (三) 节点</h2>' +
      '<ul>' +
      '<li>建议使用缩容进行节点移除操作</li>' +
      '<li>可导入节点:</li>' +
      '<ul>' +
      '<li class="text-danger">导入的节点可删除</li>' +
      '<li class="text-danger">但不支持缩容</li>' +
      '</ul>' +
      '</ul>',
    ],
    [
      '<h2>任务管理 (一) 任务模板一</h2>' +
      '<ul>' +
      '<li>扩容模板</li>' +
      '<ul>' +
      '<li class="text-danger">第一个步骤必须是create_vm(创建VM)</li>' +
      '<li class="text-danger">应包含服务启动的步骤, 步骤名称必须是start_service</li>' +
      '<li class="text-danger">若需使用服务发现, 需包含register(服务注册)步骤</li>' +
      '</ul>' +
      '<li>缩容模板</li>' +
      '<ul>' +
      '<li class="text-danger">若使用服务发现,包含unregister(服务注销)步骤</li>' +
      '<li class="text-danger">需包含return_vm(归还VM)</li>' +
      '</ul>' +
      '<li>上线模板</li>' +
      '<ul>' +
      '<li class="text-danger">应包含服务启动的步骤, 步骤名称必须是start_service</li>' +
      '</ul>' +
      '<li>注意事项</li>' +
      '<ul>' +
      '<li>模板内各步骤可通过忽略错误字段设置依赖关系</li>' +
      '<li>忽略错误: 否, 强依赖, 必须前一个步骤成功才继续</li>' +
      '<li>忽略错误: 是, 无依赖, 前一个步骤失败也会继续</li>' +
      '</ul>' +
      '</ul>',

      '<h2>任务管理 (一) 任务模板二</h2>' +
      '<ul>' +
      '<li class="text-danger">任务步骤内有一些内置的命令组, 使用时关注以下事项:</li>' +
      '<ol>' +
      '<li>步骤create_vm: 参数vm_type_id对应多云对接/机型模板(详情->序号)</li>' +
      '<li>步骤return_vm: 同上</li>' +
      '<li>步骤register: 参数service_discovery_id对应服务发现/服务注册(详情->序号)</li>' +
      '<li>步骤unregister: 同上</li>' +
      '<li>步骤install_nginx: 参数octans_host指的是下发通道(Ansible)的IP</li>' +
      '</ol>' +
      '</ul>',

      '<h2>任务管理 (二) 列表</h2>' +
      '<ul>' +
      '<li>创建任务依赖任务模板</li>' +
      '<li>每个任务只允许操作一个服务池内的机器</li>' +
      '<li>任务以IP维度下发</li>' +
      '<li>任务下发的IP, 可手动输入, 也可按服务池选择</li>' +
      '</ul>',

      '<h2>任务管理 (三) 任务详情</h2>' +
      '<ul>' +
      '<li>执行中任务, 前端每5秒自动刷新详情数据</li>' +
      '<li>节点执行结果:</li>' +
      '<ul>' +
      '<li>此步骤暂无日志: 未执行此步骤</li>' +
      '<li>此步骤日志为空: 此步骤已执行, 未产生日志</li>' +
      '</ul>' +
      '</ul>',
    ],
    [
      '<h2>命令和命令组</h2>' +
      '<ol>' +
      '<li>此模块所有名称(含参数名称), 不支持中文</li>' +
      '<li>命令实现只支持Ansible</li>' +
      '<li>命令实现参数支持专家模式</li>' +
      '<li>删除命令时, 若任务模板或命令组正在使用, 将失败</li>' +
      '</ol>',
    ],
    [
      '<h1>THANK YOU</h1>' +
      '<h1 style="color: orange;">YOU ARE THE ONE</h1>',
    ],
  ],
  for_hubble: [
    [
      '<h1>服务发现</h1>',
    ],
    [
      '<h2>首次使用注意事项</h2>' +
      '<ol>' +
      '<li>确认七层类型: 只支持 Nginx/Tengine 和 阿里云SLB</li>' +
      '<li>若使用Nginx/Tengine:</li>' +
      '<ul>' +
      '<li>依次创建分组, 单元, 然后导入节点</li>' +
      '<li>导入主配置和Upstream配置, 需要拆分Upstream</li>' +
      '<li>创建服务注册, 类型选择Nginx</li>' +
      '</ul>' +
      '<li>若使用阿里云SLB:</li>' +
      '<ul>' +
      '<li>选择SLB</li>' +
      '<li>创建服务注册, 类型选择阿里云SLB</li>' +
      '</ul>' +
      '</ol>',
    ],
    [
      '<h2>服务发现子模块</h2>' +
      '<ol>' +
      '<li>服务注册</li>' +
      '<li>七层Nginx变更管理</li>' +
      '<li>阿里云SLB控制台</li>' +
      '<li>脚本管理</li>' +
      '<li>授权管理</li>' +
      '<li>操作日志&变更历史</li>' +
      '</ol>',
    ],
    [
      '<h2>服务注册</h2>' +
      '<ul>' +
      '<li class="text-danger">仅适用于自动扩缩容流程</li>' +
      '<li>每个服务池绑定一个服务注册ID</li>' +
      '<li>服务池扩容时, 自动触发七层变更, 实现自动引入流量</li>' +
      '<li>系统内置的服务注册类型: Nginx 和 阿里云SLB</li>' +
      '</ul>',
    ],
    [
      '<h2>七层Nginx (一)</h2>' +
      '<ul>' +
      '<li>此模块提供WEB化的Nginx变更管理功能</li>' +
      '<li class="text-danger">如果服务注册类型选的是Nginx, 需要将Nginx配置导入此模块进行管理</li>' +
      '<li>系统设计默认完美兼容Tengine</li>' +
      '</ul>',

      '<h2>七层Nginx (二) 系统结构</h2>' +
      '<img style="width:100%;margin: 0px;" data-src="../images/Hubble_Nginx.png">',

      '<h2>七层Nginx (三) 相关概念</h2>' +
      '<ul>' +
      '<li>分组: 相当于Nginx集群, 每个分组可包含多个单元</li>' +
      '<li>单元: Nginx服务池, 通常以机房为维度, 每个机房一个单元</li>' +
      '<li>节点: Nginx服务器IP</li>' +
      '<li>主配置: 除Upstream外的所有配置文件</li>' +
      '</ul>',

      '<h2>七层Nginx (四) Upstream配置</h2>' +
      '<ul>' +
      '<li>以分组(集群)为单位, 全局唯一</li>' +
      '<li>每个单元Include其中某一些Upstream文件</li>' +
      '<li>Upstream的文件名是给人看的</li>' +
      '<li>Upstream文件内的名称是给Nginx看的</li>' +
      '<li class="text-danger">同一个服务各机房的Upstream文件名不相同</li>' +
      '<li class="text-danger">同一个服务各机房的Upstream文件内的名称相同</li>' +
      '<li>无版本化管理, 发布时以文件为单位</li>' +
      '</ul>',

      '<h2>七层Nginx (五) 主配置</h2>' +
      '<ul>' +
      '<li>除Upstream配置外的所有配置文件集合</li>' +
      '<li>以单元(服务池)为单位</li>' +
      '<li>每个单元使用各自独立的配置文件</li>' +
      '<li class="text-danger">主配置文件采用版本化管理</li>' +
      '<li class="text-danger">创建版本时, 可选择任意文件和任一版本</li>' +
      '<li class="text-danger">生成版本时, 自动选择每个文件的最新版本</li>' +
      '</ul>',
    ],
    [
      '<h2>阿里云SLB控制台</h2>' +
      '<ul>' +
      '<li>阿里云SLB的管理,包括Listener和Backend</li>' +
      '<li>Listener入口: SLB列表的Listener列的图标<i class="fa fa-bars"></i></li>' +
      '<li>Backend入口: SLB列表的Backend列的图标<i class="fa fa-bars"></i></li>' +
      '<li>更多使用帮助: 请参见<a href="https://help.aliyun.com/product/27537.html" target="_blank">阿里云官方文档</a></li>' +
      '</ul>',
    ],
    [
      '<h2>脚本管理</h2>' +
      '<ul>' +
      '<li>用于管理服务发现模块变更用到脚本</li>' +
      '<li>如Nginx模块使用的:</li>' +
      '<ul>' +
      '<li>updateMainConf.sh - 更新主配置的脚本</li>' +
      '<li>updateUpstreamConf.sh - 更新upstream配置的脚本</li>' +
      '</ul>' +
      '</ul>',
    ],
    [
      '<h2>授权管理</h2>' +
      '<ul>' +
      '<li>服务发现模块的授权体系</li>' +
      '<li>外部系统接入时需要生成独立的APPKEY并配置授权接口</li>' +
      '<li>不接入外部系统时,请忽略</li>' +
      '</ul>',
    ],
    [
      '<h2>操作日志 & 变更历史</h2>' +
      '<ul>' +
      '<li>操作日志: 用户在服务发现模块内的增删改操作记录</li>' +
      '<li>变更历史: 对线上服务或配置进行的变更记录</li>' +
      '</ul>',
    ],
    [
      '<h1>THANK YOU</h1>' +
      '<h1 style="color: orange;">YOU ARE THE ONE</h1>',
    ],
  ]
}

var showHelp=function(){
  var uri=window.location.pathname,title='帮助',text='',flag=false,height='';
  var arrUri=uri.split(/[/.]/);
  var module=(arrUri.length>1)?arrUri[1]:'';
  //var page=(arrUri.length>2)?arrUri[2]:module;
  if(module){
    if(typeof reveal[module] == 'object'){
      if(!$.isEmptyObject(reveal[module])){
        text+='<div class="reveal"><div class="slides">';
        $.each(reveal[module],function(k1,v1){
          text+='<section>';
          if($.isArray(v1)){
            if(v1.length>1){
              $.each(v1,function(k2,v2){
                text+='<section>'+v2+'</section>';
              });
            }else{
              text+=v1;
            }
          }else{
            text+=v2;
          }
          text+='</section>';
        });
        text+='</div></div>';
      }
    }
  }
  if(!text){
    text='<div class="alert alert-warning">未设置帮助内容</div>';
    $('#myRevealModalBody').css('height','');
  }else{
    flag=true;
  }
  $('#myRevealModalLabel').html(title);
  $('#myRevealModalBody').html(text);
  if(flag){
    Reveal.initialize({
      controls: true,
      progress: true,
      history: false,
      center: true,
      slideNumber: true,
      overview: true,
      autoSlideStoppable: true,

      transition: 'slide', // none/fade/slide/convex/concave/zoom

      dependencies: [
        {
          src: '../gentelella/vendors/reveal330/lib/js/classList.js', condition: function () {
          return !document.body.classList;
        }
        },
        {
          src: '../gentelella/vendors/reveal330/plugin/markdown/marked.js', condition: function () {
          return !!document.querySelector('[data-markdown]');
        }
        },
        {
          src: '../gentelella/vendors/reveal330/plugin/markdown/markdown.js', condition: function () {
          return !!document.querySelector('[data-markdown]');
        }
        },
        {
          src: '../gentelella/vendors/reveal330/plugin/highlight/highlight.js', async: true, callback: function () {
          hljs.initHighlightingOnLoad();
        }
        },
        {src: '../gentelella/vendors/reveal330/plugin/zoom-js/zoom.js', async: true},
        {src: '../gentelella/vendors/reveal330/plugin/notes/notes.js', async: true}
      ]
    });
    Reveal.configure({autoSlide: 1500});
  }
}