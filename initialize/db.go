package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"personFrame/pkg/common"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := common.Conf.DB.Host
	port := common.Conf.DB.Port
	username := common.Conf.DB.Username
	password := common.Conf.DB.Password
	database := common.Conf.DB.Database
	charset := common.Conf.DB.Charset

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	mysqlConfig := mysql.Config{
		DSN:                       args,  // DSN data source name
		DefaultStringSize:         191,   // string 类型字段的默认长度
		SkipInitializeWithVersion: false, // 根据版本自动配置
	}
	db, err := gorm.Open(mysql.New(mysqlConfig))
	if err != nil {
		panic("连接数据库失败" + err.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
