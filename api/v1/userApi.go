package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"personFrame/ini/request"
	"personFrame/ini/response"
	"personFrame/model"
	commonModel "personFrame/model/common"
	"personFrame/pkg/common"
	"personFrame/utils"
)

// Login 登录示例
func (b *BaseApi) Login(c *gin.Context) {
	var loginInfo request.Login
	_ = c.ShouldBindJSON(&loginInfo)
	if err := utils.Verify(loginInfo, utils.SimpleLoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &model.User{Username: loginInfo.Username, Password: loginInfo.Password}
	if err, user := userService.Login(u); err != nil {
		common.GlobalLog.Error("登陆失败! 用户名不存在或者密码错误!", zap.Error(err))
		response.FailWithMessage("用户名不存在或者密码错误", c)
		return
	} else {
		b.tokenNext(c, *user)
	}
}

//签发Token
func (b *BaseApi) tokenNext(c *gin.Context, user model.User) {
	j := &utils.JWT{SigningKey: []byte(common.Conf.JWT.SigningKey)}
	claims := j.CreatClaims(request.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		Username:    user.Username,
		NickName:    user.NickName,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.CreateToken(claims)
	if err != nil {
		common.GlobalLog.Error("获取token失败！", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	if !common.Conf.System.UseMultiPoint {
		response.SuccessWithDetail(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功", c)
		return
	}

	if err, jwtStr := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			common.GlobalLog.Error("设置登录状态失败！", zap.Error(err))
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.SuccessWithDetail(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功！", c)
	} else if err != nil {
		common.GlobalLog.Error("设置登录状态失败！", zap.Error(err))
		response.FailWithMessage("设置登录状态失败！", c)
	} else {
		var blackJWT commonModel.JwtBlackList
		blackJWT.Jwt = jwtStr
		if err := jwtService.JsonInBlackList(blackJWT); err != nil {
			response.FailWithMessage("jwt作废失败", c)
			return
		}
		if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
			response.FailWithMessage("设置登录状态失败", c)
			return
		}
		response.SuccessWithDetail(response.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
		}, "登录成功！", c)
	}
}
