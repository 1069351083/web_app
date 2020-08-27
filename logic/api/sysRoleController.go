package api

import (
	"web_app/model"
	"web_app/response"
	"web_app/service"

	"github.com/gin-gonic/gin"
)

func LoadAllRole(ctx *gin.Context) {

	page := new(model.SysRolePage)
	ctx.ShouldBindJSON(&page)
	list, count, err := service.LoadAllRole(page)
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.OkWithData(response.ResponsePage{
		List:     list,
		Total:    count,
		PageSize: page.PageSize,
		Page:     page.Page,
	}, ctx)

}

func AddRole(ctx *gin.Context) {
	role := new(model.SysRole)
	ctx.ShouldBindJSON(&role)
	err := service.AddRole(role)
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.Ok(ctx)
}
