package model

import (
	"web_app/model/common"
)

type SysRole struct {
	Id         int          `json:"id"`
	Name       string       `json:"name"` // 角色名称
	Remark     string       `json:"remark"`
	Available  int          `json:"available"`
	Createtime common.XTime `json:"createtime"`
}

type SysRolePage struct {
	Id         int          `json:"id"`
	Name       string       `json:"name"` // 角色名称
	Remark     string       `json:"remark"`
	Available  int          `json:"available"`
	Createtime common.XTime `json:"createtime"`
	Page       int          `json:"page" `
	PageSize   int          `json:"pageSize" form:"pageSize"`
	StartTime  common.XTime `json:"starttime"`
	EndTime    common.XTime `json:"endtime"`
}

type MenuRole struct {
	Rid  int   `json:"rid"`
	Mids []int `json:"mids"`
}
