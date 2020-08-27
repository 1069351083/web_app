package sysUser

import (
	"web_app/logic/api"
	"web_app/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(group *gin.RouterGroup) {
	UserRouter := group.Group("user").Use(middleware.JWTAuth())
	{
		UserRouter.POST("list", api.QueryListUser)
		UserRouter.POST("loadUserMaxOrderNum", api.LoadUserMaxOrderNum)
	}
}
