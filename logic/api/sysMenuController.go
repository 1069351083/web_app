package api

import (
	"web_app/model"
	"web_app/response"
	"web_app/service"

	"github.com/gin-gonic/gin"
)

func IndexLeftMenu(ctx *gin.Context) {

	claims, ok := ctx.Get("claims")
	if !ok {
		response.FailWithMessage("claims获取失败", ctx)
		return
	}
	cla := claims.(*model.MyClaims)

	page := new(model.PageInfo)
	ctx.ShouldBindJSON(&page)
	menuList, count, err := service.IndexLeftMenu(page, cla)

	if err != nil {
		response.Fail(ctx)
		return
	}
	response.OkWithData(response.ResponsePage{
		List:     menuList,
		Total:    count,
		PageSize: page.PageSize,
		Page:     page.Page,
	}, ctx)

}

func LoadMenuManagerLeftTreeJson(ctx *gin.Context) {

	menuTreeList, err := service.LoadMenuManagerLeftTreeJson()
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.OkWithData(menuTreeList, ctx)

}

func LoadAllMenu(ctx *gin.Context) {
	menu := new(model.SysPermission)
	page := new(model.PageInfo)
	ctx.ShouldBindJSON(&page)

	menList, count, err := service.LoadAllMenu(page, menu)
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.OkWithData(response.ResponsePage{
		List:     menList,
		Total:    count,
		PageSize: page.PageSize,
		Page:     page.Page,
	}, ctx)
}

func AddMenu(ctx *gin.Context) {
	menu := new(model.SysPermission)
	ctx.ShouldBindJSON(&menu)

	err := service.AddMenu(menu)
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.Ok(ctx)

}

func UpdateMenu(ctx *gin.Context) {
	menu := new(model.SysPermission)
	ctx.ShouldBindJSON(&menu)
	err := service.UpdateMenu(menu)
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.Ok(ctx)
}

func DeleteMenu(ctx *gin.Context) {
	menu := new(model.SysPermission)
	ctx.ShouldBind(&menu)
	err := service.DeleteMenu(menu)
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.Ok(ctx)

}

func CheckMenuHasChildrenNode(ctx *gin.Context) {
	menu := new(model.SysPermission)
	ctx.ShouldBindJSON(&menu)
	ok, err := service.CheckMenuHasChildrenNode(menu)
	if err != nil {
		response.Fail(ctx)
		return
	}
	m := make(map[string]bool)
	m["value"] = ok
	response.OkWithData(
		m, ctx)
}

func LoadMenuMaxOrderNum(ctx *gin.Context) {
	menu := new(model.SysPermission)
	ctx.ShouldBindJSON(&menu)
	err, max := service.LoadMenuMaxOrderNum()
	if err != nil {
		response.Fail(ctx)
		return
	}
	response.OkWithData(max+1, ctx)
}
