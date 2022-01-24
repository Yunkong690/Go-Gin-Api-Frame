package common

import (
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	model "personFrame/model/common"
)

type Config struct {
	System model.System
	DB     model.DBConfig
	Redis  model.RedisConfig
	JWT    model.JWT
	Zap    model.Zap
}

var Conf = &Config{}
var (
	GlobalLog          *zap.Logger             //日志
	ConcurrencyControl = &singleflight.Group{} //防缓存击穿
	BlackCache         local_cache.Cache
	DB                 *gorm.DB
	Redis              *redis.Client
)

func InitConfig(configPath string) *Config {
	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	viper.OnConfigChange(func(in fsnotify.Event) {
		err = viper.Unmarshal(&Conf)
		if err != nil {
			panic(err)
		}
	})
	return Conf
}
