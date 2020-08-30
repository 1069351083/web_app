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

func UpdateRole(ctx *gin.Context) {
	role := new(model.SysRole)
	ctx.ShouldBindJSON(&role)
	err := service.UpdateRole
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.OkWithMessage("更新成功", ctx)

}

func DeleteRole(ctx *gin.Context) {
	role := new(model.SysRole)
	ctx.ShouldBindJSON(&role)
	err := service.DeleteRole(role)
	if err != nil {
		response.FailWithMessage("删除失败", ctx)
		return
	}
	response.OkWithMessage("删除成功", ctx)
}

func InitPermissionByRoleId(ctx *gin.Context) {
	role := new(model.SysRole)
	ctx.ShouldBindJSON(&role)
	treeNodes, err := service.InitPermissionByRoleId(role)
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.OkWithData(response.ResponsePage{
		List:     treeNodes,
		Total:    0,
		PageSize: 0,
		Page:     0,
	}, ctx)
}

func SaveRolePermission(ctx *gin.Context) {

	mr := new(model.MenuRole)
	ctx.ShouldBindJSON(&mr)
	err := service.SaveRolePermission(mr)
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.Ok(ctx)

}
