package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"personFrame/ini/response"
	commonModel "personFrame/model/common"
	"personFrame/pkg/common"
	commonService "personFrame/service/common"
	"personFrame/utils"
	"strconv"
	"time"
)

var jwtService = commonService.JwtService{}

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		//JWT鉴权取头部信息x-token 登录时返回token，前端需要存储token
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithDetail(gin.H{"reload": true}, "未登录或非法访问", c)
			c.Abort()
			return
		}
		if jwtService.IsBlackList(token) {
			response.FailWithDetail(gin.H{"reload": true}, "您的账户异地登录或令牌失效", c)
			c.Abort()
			return
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.FailWithDetail(gin.H{"reload": true}, "授权已过期", c)
				c.Abort()
				return
			}
			response.FailWithDetail(gin.H{"reload": true}, err.Error(), c)
			c.Abort()
			return
		}
		if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt = time.Now().Unix() + common.Conf.JWT.ExpiresTime
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
			if common.Conf.System.UseMultiPoint { //允许多点登录
				err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					common.GlobalLog.Error("get redis jwt failed", zap.Error(err))
				} else {
					//当旧token取成功方可进行拉黑操作
					_ = jwtService.JsonInBlackList(commonModel.JwtBlackList{Jwt: RedisJwtToken})
				}
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
