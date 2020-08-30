package api

import (
	"web_app/model"
	"web_app/response"
	"web_app/service"

	"github.com/gin-gonic/gin"
)

func UpdateCasbin(ctx *gin.Context) {
	m := new(model.CasbinInReceive)
	ctx.ShouldBindJSON(&m)
	err := service.UpdateCasbin(m.Name, m.CasbinInfos)
	if err != nil {
		response.FailWithMessage("更新权限失败", ctx)
		return
	}
	response.Ok(ctx)

}

func GetPolicyPathByAuthorityId(c *gin.Context) {
	var cmr model.CasbinInReceive
	_ = c.ShouldBindJSON(&cmr)

	paths := service.GetPolicyPathByAuthorityId(cmr.Name)
	response.OkWithData(gin.H{
		"Paths": paths,
	}, c)
}
