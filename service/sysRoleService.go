package service

import (
	"encoding/json"
	"web_app/dao/mysql"
	"web_app/dao/redis"
	"web_app/model"
	"web_app/service/Cache_service"
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
