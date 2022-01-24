package main

import (
	"personFrame/initialize"
	"personFrame/middleware"
	"personFrame/model"
	"personFrame/pkg/common"
	common2 "personFrame/service/common"
)

func main() {
	common.InitConfig("conf/config.yaml") //导入配置文件
	common.DB = initialize.InitDB()       //初始化数据库
	common.Redis = initialize.InitRedis() //初始化Redis
	common.GlobalLog = middleware.Zap()   //使用Log
	if common.DB != nil {
		model.RegisterTables(initialize.DB) //注册数据表
		common2.LoadAll()                   //从数据库加载 Jwt信息
		db, _ := common.DB.DB()
		defer db.Close()
	}
	redis := common.Redis
	defer redis.Close()
	Router := initialize.Routers() //注册路由
	s := initialize.InitServer(common.Conf.System.ListenAddr, Router)
	common.GlobalLog.Error(s.ListenAndServe().Error())
}
