package service

import (
	"web_app/dao/mysql"
	"web_app/model"
	"web_app/request"
)

func QueryListUser(info model.PageInfo) (err error, userList []model.SysUser, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := mysql.DB
	err = db.Find(&userList).Count(&total).Error
	err = db.Find(&userList).Limit(limit).Offset(offset).Error
	return err, userList, total
}

func Login(user *request.RegisterAndLoginStruct) (error, *model.SysUser) {
	var u model.SysUser
	db := mysql.DB
	err := db.Where("loginname=? AND pwd=?", user.LoginName, user.Pwd).First(&u).Error
	return err, &u
}

func LoadUserMaxOrderNum(user *model.SysUser) map[string]interface{} {
	map1 := make(map[string]interface{})
	var list []*model.SysUser
	db := mysql.DB
	err := db.Order("ordernum  desc").Find(&list).Error
	if err != nil {
		panic(err)
	}

	if len(list) != 0 {
		max := list[0].Ordernum + 1
		map1["value"] = max
		return map1
	}
	map1["value"] = 1
	return map1

}
