package sysRole

import (
	"web_app/logic/api"

	"github.com/gin-gonic/gin"
)

func InitRoleRouter(group *gin.RouterGroup) {

	routerGroup := group.Group("role")
	{
		routerGroup.POST("loadAllRole", api.LoadAllRole)
		routerGroup.POST("addRole", api.AddRole)
		routerGroup.POST("updateRole", api.UpdateRole)
		routerGroup.POST("deleteRole", api.DeleteRole)
		routerGroup.POST("initPermissionByRoleId", api.InitPermissionByRoleId)
		routerGroup.POST("saveRolePermission", api.SaveRolePermission)

	}
}
