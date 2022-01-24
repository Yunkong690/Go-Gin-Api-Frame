package common

import (
	"context"
	"go.uber.org/zap"
	commonModel "personFrame/model/common"
	"personFrame/pkg/common"
	"time"
)

type JwtService struct {
}

// GetRedisJWT 从Redis取出JWT信息
func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = common.Redis.Get(context.Background(), userName).Result()
	return err, redisJWT
}

// SetRedisJWT 向Redis缓存JWT信息
func (jwtService *JwtService) SetRedisJWT(jwt, userName string) (err error) {
	timer := time.Duration(common.Conf.JWT.ExpiresTime) * time.Second
	err = common.Redis.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

// JsonInBlackList Jwt加入黑名单
func (jwtService *JwtService) JsonInBlackList(jwtList commonModel.JwtBlackList) (err error) {
	err = common.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	common.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

// IsBlackList 判断是否在黑名单
func (jwtService *JwtService) IsBlackList(jwt string) bool {
	_, ok := common.BlackCache.Get(jwt)
	return ok
}

func LoadAll() {
	var data []string
	err := common.DB.Model(&commonModel.JwtBlackList{}).Select("jwt").Find(&data).Error
	if err != nil {
		common.GlobalLog.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		common.BlackCache.SetDefault(data[i], struct{}{})
	}
}
