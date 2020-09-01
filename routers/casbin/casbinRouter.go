package casbin

import (
	"web_app/logic/api"

	"github.com/gin-gonic/gin"
)

func InitCasBinRouter(group *gin.RouterGroup) {
	routerGroup := group.Group("casbin")
	{
		routerGroup.POST("updateCasbin", api.UpdateCasbin)
		routerGroup.POST("queryAllCasbin", api.GetPolicyPathByAuthorityId)
	}

}
