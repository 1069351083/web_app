package routers

import (
	"web_app/logger"
	"web_app/response"
	"web_app/routers/BaseRouter"
	"web_app/routers/casbin"
	"web_app/routers/sysPermission"
	"web_app/routers/sysRole"
	"web_app/routers/sysUser"

	"github.com/gin-gonic/gin"
)

func SetUp() *gin.Engine {
	r := gin.New()

	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	ApiGroup := r.Group("")
	BaseRouter.InitBaseRouter(ApiGroup)
	sysUser.InitUserRouter(ApiGroup)
	sysPermission.InitMenuRouter(ApiGroup)
	sysRole.InitRoleRouter(ApiGroup)
	casbin.InitCasBinRouter(ApiGroup)
	r.GET("/", func(context *gin.Context) {
		response.OkWithMessage("操作成功", context)
	})
	return r
}
