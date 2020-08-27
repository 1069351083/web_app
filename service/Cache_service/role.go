package Cache_service

import (
	"strconv"
	"strings"
	"web_app/model/common"
)

type RoleCache struct {
	Id   int
	Name string

	Page     int
	PageSize int
}

func (a *RoleCache) GetRoleCacheKey() string {

	return common.CACHE_ROLE + "_" + strconv.Itoa(a.Id)
}

func (a *RoleCache) GetRolesKey() string {
	keys := []string{
		common.CACHE_ROLE,
		"LIST",
	}

	if a.Id > 0 {
		keys = append(keys, strconv.Itoa(a.Id))
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
