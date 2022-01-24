# Gin-Web-Api-Framework

## 基于Gin封装的个人golang后端开发框架

### 项目背景

规范化封装个人使用的Golang Web后端API框架,目前集成了Gorm+JWT+Zap，摸鱼没事干瞎写的

### 项目结构

```
|-- go.mod
|-- main.go
|-- Readme.md
|-- api                         //Api文件夹，负责接收及及返回数据
|   |-- init.go                 //注册api组
|   |-- v1
|       |-- baseApi.go         //示例
|       |-- init.go            //注册v1版本的路由组
|       |-- userApi.go         //示例Api，里面写了一个登录逻辑
|-- conf                       //项目配置文件夹
|   |-- config.yaml            //项目配置文件
|-- ini                        //定义接收和返回数据
|   |-- request                //接收参数
|   |   |-- jwt.go
|   |   |-- user.go
|   |-- response              //返回参数
|       |-- response.go       
|       |-- user.go
|-- initialize                //基础功能初始化
|   |-- db.go                   
|   |-- redis.go
|   |-- router.go
|   |-- server.go
|-- log                       //项目日志
|   |-- server_error.log
|   |-- server_info.log
|   |-- server_warn.log
|-- middleware                //中间件
|   |-- jwt.go
|   |-- zap.go
|-- model                    //model层，定义数据库结构体
|   |-- init.go              //数据库表注册用
|   |-- user.go
|   |-- common               //定义配置文件结构体与通用结构体
|       |-- DBConfig.go
|       |-- GlobalModel.go    //全局通用结构体
|       |-- JwtConfig.go
|       |-- RedisConfig.go
|       |-- SystemConfig.go
|       |-- ZapConfig.go
|-- pkg                       
|   |-- error.go              //错误信息
|   |-- common                //初始化配置文件，引入全局变量
|       |-- config.go         //初始化全局配置文件
|-- router                    //路由
|   |-- init.go               //注册总路由
|   |-- v1
|       |-- base.go           //示例路由
|       |-- init.go           //注册v1版本路由组
|-- service                   //service层，具体逻辑实现
|   |-- init.go               //注册总逻辑，方便调用
|   |-- common                //全局变量相关逻辑
|   |   |-- jwt_service.go
|   |-- v1
|       |-- init.go           //注册v1版本相关逻辑
|       |-- user_service.go
|-- utils                     //工具包
    |-- dirUtils.go
    |-- jwtUtils.go
    |-- Md5Utils.go
    |-- rotatelogsUtils.go
    |-- validator.go         //参数校验
    |-- verify.go            //定义参数校验规则
```

### 安装

Clone 本项目到本地,打开终端，进入到项目根目录，执行以下命令

```shell
go mod tidy
```

然后打开conf文件里的conf.yaml文件，填写你的配置信息，执行以下命令启动项目

```shell
go run main.go
```

###开发顺序

model层

1.model层编写相关数据库结构体例如User（注意结构体名称要与数据库表名一致，参考gorm的文档）
2.在model文件夹里的init.go里注册User结构体，即完成了表的注册

service层
1.在service层的相对应的版本文件夹里编写相关逻辑
2.在init.go文件夹里注册刚刚编写的service

api层
1.在版本文件夹（如v1）里编写相关api代码
2.在v1版本文件里注册api结构体
3.在api文件夹的init.go注册v1版本的api组

router层
编写逻辑同上，但是写完记得去initialize文件夹里的router.go里注册路由


###个人说明
第一个版本，瑕疵可能不少，参考了gin-vue-admin的server的设计思路，后续可能会慢慢完善，当前版本已经能满足基本开发需求。