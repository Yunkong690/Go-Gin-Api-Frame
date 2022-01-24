package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	common2 "personFrame/ini/request"
	"personFrame/pkg/common"
	"time"
)

type JWT struct {
	SigningKey []byte
}

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		[]byte(common.Conf.JWT.SigningKey),
	}
}

func (j *JWT) CreatClaims(baseClaims common2.BaseClaims) common2.CustomClaims {
	claims := common2.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: common.Conf.JWT.BufferTime, //缓存时间1天 缓冲时间内会获取新token刷新令牌，此时一个用户虽有两个有效令牌，但是前端只留一个
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                        //签名生效时间
			ExpiresAt: time.Now().Unix() + common.Conf.JWT.ExpiresTime, //过期时间7天 来自配置文件
			Issuer:    common.Conf.JWT.Issuer,                          //q签名发行者
		},
	}
	return claims
}

// CreateToken 生成token
func (j *JWT) CreateToken(claims common2.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims common2.CustomClaims) (string, error) {
	v, err, _ := common.ConcurrencyControl.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.CreateToken(claims)
	})
	return v.(string), err
}

// ParseToken 解析token
func (j *JWT) ParseToken(tokenString string) (*common2.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &common2.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*common2.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}
