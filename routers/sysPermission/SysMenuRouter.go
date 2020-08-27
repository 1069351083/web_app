package sysPermission

import (
	"web_app/logic/api"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter(group *gin.RouterGroup) {
	MenuRouter := group.Group("menu")
	{
		MenuRouter.POST("indexLeftMenu", api.IndexLeftMenu)
		MenuRouter.POST("loadMenuManagerLeftTreeJson", api.LoadMenuManagerLeftTreeJson)
		MenuRouter.POST("loadAllMenu", api.LoadAllMenu)
		MenuRouter.POST("addMenu", api.AddMenu)
		MenuRouter.POST("updateMenu", api.UpdateMenu)
		MenuRouter.DELETE("deleteMenu", api.DeleteMenu)
		MenuRouter.POST("checkMenuHasChildrenNode", api.CheckMenuHasChildrenNode)
		MenuRouter.POST("loadMenuMaxOrderNum", api.LoadMenuMaxOrderNum)
	}

}
