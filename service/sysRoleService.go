package service

import (
	"encoding/json"
	"errors"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/model"
	"web_app/service/Cache_service"

	"go.uber.org/zap"
)

func LoadAllRole(info *model.SysRolePage) (list []*model.SysRole, count int, err error) {
	db := mysql.DB
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	cache := Cache_service.RoleCache{
		Id:       info.Id,
		Name:     info.Name,
		Page:     info.Page,
		PageSize: info.PageSize,
	}
	key := cache.GetRolesKey()
	if redis.Exists(key) {
		get, e := redis.Get(key)
		json.Unmarshal(get, &list)
		count = len(list)
		err = e
		return
	}

	err = db.Where("name like ? AND remark like ?  AND  ? <  createtime < ?", "%"+info.Name+"%", "%"+info.Remark+"%", info.StartTime, info.EndTime).Find(&list).Count(&count).Error
	err = db.Where("name like ? AND remark like ?", "%"+info.Name+"%", "%"+info.Remark+"%").Limit(limit).Offset(offset).Find(&list).Error
	redis.Set(key, list, 3600)
	return
}

func AddRole(role *model.SysRole) error {
	db := mysql.DB
	err := db.Create(&role).Error
	return err

}

func UpdateRole(role *model.SysRole) error {
	db := mysql.DB
	err := db.Save(&role).Error
	return err

}

func DeleteRole(role *model.SysRole) error {
	db := mysql.DB
	err := db.Where("rid = ?", role.Id).Find(&model.SysRoleUser{}).Error
	if err == nil {
		zap.L().Error("该角色已被分配给用户")
		return errors.New("此角色有用户正在使用禁止删除")
	}
	e := db.Where("id = ?", role.Id).Delete(&role).Error
	return e

}

func InitPermissionByRoleId(role *model.SysRole) ([]model.TreeNode, error) {
	db := mysql.DB
	var menus []*model.SysPermission
	var proles []*model.SysRolePermission
	var menuByIds []*model.SysPermission
	var nodes []model.TreeNode
	var menuIds []int
	var checkArr string
	var spread bool
	db.Find(&menus)
	db.Where("rid = ?", role.Id).Find(&proles)
	if len(proles) != 0 {
		for menuId, _ := range proles {
			menuIds = append(menuIds, menuId)
		}
	}
	err := db.Where("id IN (?)", menuIds).Find(&menuByIds).Error
	if err != nil {
		zap.L().Error("查询失败")
		return nil, err
	}
	for _, menu := range menuByIds {
		for _, m := range menus {
			if menu.Id == m.Id {
				checkArr = "1"
			}
		}
		if menu.Open == 1 {
			spread = true
		} else {
			spread = false
		}
		Nmenu := model.TreeNode{
			Id:       menu.Id,
			Pid:      menu.Pid,
			Spread:   spread,
			CheckArr: checkArr,
		}
		nodes = append(nodes, Nmenu)
	}
	return nodes, nil

}

func SaveRolePermission(mr *model.MenuRole) error {
	db := mysql.DB
	err := db.Where("rid = ?", "%jinzhu%").Delete(&model.SysRolePermission{}).Error
	if err != nil {
		zap.L().Error("删除失败")
		return err
	}
	for _, mid := range mr.Mids {
		e := db.Create(&model.SysRolePermission{
			Rid: mr.Rid,
			Pid: mid,
		}).Error
		if e != nil {
			zap.L().Error("添加失败")
			return e
		}

	}
	return nil

}
