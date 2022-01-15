package utils

import (
	"errors"
	"github.com/GanweiYee/ducker/server/global"
	systemmodel "github.com/GanweiYee/ducker/server/model"
	"github.com/dgrijalva/jwt-go"
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
		[]byte(global.Config.Jwt.SigningKey),
	}
}

func (j *JWT) CreateClaims(baseClaims systemmodel.BaseClaims) systemmodel.CustomClaims {
	claims := systemmodel.CustomClaims{
		BaseClaims: baseClaims,
		BufferTime: global.Config.Jwt.BufferTime, // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                          // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.Config.Jwt.ExpiresTime, // 过期时间 7天  配置文件
			Issuer:    global.Config.Jwt.Issuer,                          // 签名的发行者
		},
	}
	return claims
}

// 创建一个token
func (j *JWT) CreateToken(claims systemmodel.CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
//func (j *JWT) CreateTokenByOldToken(oldToken string, claims systemmodel.CustomClaims) (string, error) {
//	v, err, _ := global.GVA_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
//		return j.CreateToken(claims)
//	})
//	return v.(string), err
//}

// 解析 token
func (j *JWT) ParseToken(tokenString string) (*systemmodel.CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &systemmodel.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*systemmodel.CustomClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}
}
