# Hubble

## 简介

Hubble 是weibo平台开源的一个服务发现组件, 兼容Nginx,slb,elb等。


## 项目结构

**注意:模块的命名和编写请遵循 PSR-1 和 PSR-4**

**注意:模块的命名和编写请遵循 PSR-1 和 PSR-4**

**注意:模块的命名和编写请遵循 PSR-1 和 PSR-4**

重要的事情说三遍。

好,下面我们介绍目录

所有的公共函数放入到 Common/Common/function.php中。

如:日志,权限校验等。

命名规则为snake。

所有直接操作数据库的操作放入 Common/Dao 文件夹中,遵循 PSR-1和PSR-4。

Controller中不允许出现任何直接操作数据库的函数和过程。

数据库操作返回规则:
    
    add, update, delete 三种操作 只可以返回 code,msg 的 Array
    
    select 操作 需要返回 code,msg,content 的Array
    
## 代码规则

函数命名:

    驼峰式,动名词结构
    
变量命名:

    驼峰式
    
    不允许出现 无意义变量名。如: $a,$b

注释规则:

    使用多行注释 /* xxx \*/
    
    函数注释需要写明参数及返回值,如:
    
>         /*
         * 通过服务名称获取服务ID
         * @param string name 表示服务发现的类型
         * @return mixed  
         * 成功  int 服务ID
         * 失败  bool false
         */
         
函数之间最少间隔一行,最多间隔2行,最好统一一行。
