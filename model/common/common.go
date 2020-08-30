package common

import "github.com/casbin/casbin"

var CasbinEnforcer *casbin.SyncedEnforcer

const (
	SUPER_USER = 0
)
const (
	CACHE_USER = "USER"
	CACHE_MENU = "MENU"
	CACHE_ROLE = "ROLE"
)
