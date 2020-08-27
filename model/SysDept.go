package model

import "time"

type SysDept struct {
	Id         int       `json:"id"`
	Pid        int       `json:"pid"`
	Title      string    `json:"title"`
	Open       int       `json:"open"`
	Remark     string    `json:"remark"`
	Address    string    `json:"address"`
	Available  int       `json:"available"` // 状态【0不可用1可用】
	Ordernum   int       `json:"ordernum"`  // 排序码【为了调事显示顺序】
	Createtime time.Time `json:"createtime"`
}
