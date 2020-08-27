package middleware

import (
	"errors"
	"web_app/model"
	"web_app/response"
	"web_app/settings"

	"github.com/gin-gonic/gin"

	"github.com/dgrijalva/jwt-go"
)

func JWTAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带Token有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// 这里的具体实现方式要依据你的实际业务情况决定
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.FailWithMessage("请求头为空", c)
			c.Abort()
			return
		}

		Jwt := &SingleKeyConf{SingleKey: []byte(settings.Conf.SingleKey)}
		// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		mc, err := Jwt.ParseToken(token)
		if err != nil {
			response.FailWithMessage("无效的Token", c)
			c.Abort()
			return
		}
		// 将当前请求的username信息保存到请求的上下文c上
		c.Set("claims", mc)
		c.Next() // 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}

type SingleKeyConf struct {
	SingleKey []byte
}

func (JWT SingleKeyConf) CreateJWT(claims model.MyClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWT.SingleKey)
}

// ParseToken 解析JWT
func (JWT SingleKeyConf) ParseToken(tokenString string) (*model.MyClaims, error) {
	// 解析token
	token, err := jwt.ParseWithClaims(tokenString, &model.MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return settings.Conf.SingleKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*model.MyClaims); ok && token.Valid { // 校验token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
