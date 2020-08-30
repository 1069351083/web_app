package middleware

import (
	"web_app/model"
	"web_app/response"
	"web_app/service"
	"web_app/settings"

	"github.com/gin-gonic/gin"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, _ := c.Get("claims")
		j := claims.(model.MyClaims)
		//获取请求的URI
		obj := c.Request.URL.RequestURI()
		//获取请求方法
		act := c.Request.Method
		//获取用户的角色
		sub := j.Name
		//数据库持久化
		e := service.Casbin()
		//判断策略中是否存在
		if settings.Conf.Mode == "dev" || e.Enforce(sub, obj, act) {
			c.Next()
		} else {
			response.Result(response.ERROR, gin.H{}, "权限不足", c)
			c.Abort()
			return
		}

	}
}
