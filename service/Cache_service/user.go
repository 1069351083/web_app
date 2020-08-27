package Cache_service

import (
	"strconv"
	"strings"
	"web_app/model/common"
)

type UserCache struct {
	Id        int
	LoginName string
	Name      string

	Page     int
	PageSize int
}

func (a *UserCache) GetUserCacheKey() string {

	return common.CACHE_USER + "_" + strconv.Itoa(a.Id)
}

func (a *UserCache) GetUsersKey() string {
	keys := []string{
		common.CACHE_USER,
		"LIST",
	}

	if a.Id > 0 {
		keys = append(keys, strconv.Itoa(a.Id))
	}
	if a.LoginName != "" {
		keys = append(keys, a.LoginName)
	}
	if a.Name != "" {
		keys = append(keys, a.Name)
	}

	if a.Page > 0 {
		keys = append(keys, strconv.Itoa(a.Page))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
