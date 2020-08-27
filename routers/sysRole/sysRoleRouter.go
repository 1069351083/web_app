package sysRole

import (
	"web_app/logic/api"

	"github.com/gin-gonic/gin"
)

func InitRoleRouter(group *gin.RouterGroup) {

	routerGroup := group.Group("role")
	{
		routerGroup.POST("loadAllRole", api.LoadAllRole)
		routerGroup.POST("addRole")
	}
}
