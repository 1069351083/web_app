package service

import (
	"encoding/json"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/model"
	"web_app/model/common"
	"web_app/service/Cache_service"

	"go.uber.org/zap"
)

func IndexLeftMenu(page *model.PageInfo, clamis *model.MyClaims) (menuList []interface{}, count int, err error) {
	var user model.SysUser
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	db := mysql.DB
	userid := clamis.Id
	db.Where("id=?", userid).Find(&user)
	if user.Type == common.SUPER_USER {
		db.Where("type=? AND available=?", "menu", true).Limit(limit).Offset(offset).Find(&menuList).Count(&count)
		return menuList, count, err
	}
	db.Where("type=? AND available=?", "menu", true).Limit(limit).Offset(offset).Find(&menuList).Count(&count)
	return menuList, count, err

}

func LoadMenuManagerLeftTreeJson() (menuTreeList []interface{}, err error) {
	var spread bool
	var menuList []model.SysPermission
	db := mysql.DB
	err = db.Where("type=?", "menu").Find(&menuList).Error
	for _, value := range menuList {
		if value.Open == 1 {
			spread = true
		} else {
			spread = false
		}
		menuTree := model.NewTreeNode(value.Id, value.Pid, spread, value.Title)

		menuTreeList = append(menuTreeList, menuTree)

	}
	return menuTreeList, err
}

func LoadAllMenu(page *model.PageInfo, menu *model.SysPermission) (menuList []*model.SysPermission, count int, err error) {
	limit := page.PageSize
	offset := page.PageSize * (page.Page - 1)
	db := mysql.DB
	cache := &Cache_service.MenuCache{
		Id:       menu.Id,
		Tittle:   menu.Title,
		Page:     page.Page,
		PageSize: page.PageSize,
	}
	key := cache.GetMenusKey()
	if redis.Exists(key) {
		data, err := redis.Get(key)
		if err != nil {
			zap.L().Error("获取menu缓存失败", zap.Error(err))
			return nil, 0, err
		}
		json.Unmarshal(data, &menuList)
		count := len(menuList)
		return menuList, count, nil
	}
	if menu.Title != "" {
		err = db.Where("title like ?", "%menu.Title%").Count(&count).Find(&menuList).Error
		err = db.Where("title like ?", "%menu.Title%").Limit(limit).Offset(offset).Find(&menuList).Error
		return menuList, count, err

	}
	err = db.Find(&menuList).Count(&count).Error
	err = db.Limit(limit).Offset(offset).Find(&menuList).Error
	if err != nil {
		zap.L().Error("获取menu数据失败", zap.Error(err))
		return nil, 0, err
	}
	redis.Set(key, menuList, 3600)
	return menuList, count, err

}

func AddMenu(menu *model.SysPermission) error {
	db := mysql.DB

	err := db.Create(&menu).Error
	if err != nil {
		return err

	}

	return nil
}

func UpdateMenu(menu *model.SysPermission) error {

	db := mysql.DB
	err := db.Model(&menu).Updates(menu).Error
	if err != nil {
		return err

	}
	return nil
}

func DeleteMenu(menu *model.SysPermission) error {
	db := mysql.DB
	err := db.Unscoped().Delete(&menu).Error
	if err != nil {
		return err

	}

	return nil
}

func CheckMenuHasChildrenNode(m interface{}) (bool, error) {
	var menus []*model.SysPermission
	db := mysql.DB
	menu, ok := m.(*model.SysPermission)
	if ok {

		err := db.Where("pid = ?", menu.Pid).Find(&menus).Error
		if len(menus) != 0 {
			return true, err
		}
		return false, err
	}
	return false, nil
}

func LoadMenuMaxOrderNum() (err error, max int) {
	var list []*model.SysPermission
	db := mysql.DB
	err = db.Order("id desc").Find(&list).Error
	max = list[0].Id
	return err, max
}
