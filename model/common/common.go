package common

import "github.com/casbin/casbin"

var CasbinEnforcer *casbin.SyncedEnforcer

const (
	SUPER_USER = 0
)
const (
	CACHE_USER       = "USER"
	CACHE_MENU       = "MENU"
	CACHE_ROLE       = "ROLE"
	KeyPrefix        = "bluebell_"
	KeyPostTimeZset  = "post_time"
	KeyPostScoreZset = "post_score"
	KeyPostVotedZset = "post_voted"
)
