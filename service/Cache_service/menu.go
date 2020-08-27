package Cache_service

import (
	"strconv"
	"strings"
	"web_app/model/common"
)

type MenuCache struct {
	Id     int
	Tittle string

	Page     int
	PageSize int
}

func (a *MenuCache) GetMenuCacheKey() string {

	return common.CACHE_USER + "_" + strconv.Itoa(a.Id)
}

func (a *MenuCache) GetMenusKey() string {
	keys := []string{
		common.CACHE_MENU,
		"LIST",
	}

	if a.Id > 0 {
		keys = append(keys, strconv.Itoa(a.Id))
	}
	if a.Tittle != "" {
		keys = append(keys, a.Tittle)
	}

	if a.Page > 0 {
		keys = append(keys, strconv.Itoa(a.Page))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}
