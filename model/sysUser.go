package model

import "time"

type SysUser struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Loginname string    `json:"loginname"`
	Address   string    `json:"address"`
	Sex       int       `json:"sex"`
	Remark    string    `json:"remark"`
	Pwd       string    `json:"pwd"`
	Deptid    int       `json:"deptid"`
	Hiredate  time.Time `json:"hiredate"`
	Mgr       int       `json:"mgr"`
	Available int       `json:"available"`
	Ordernum  int       `json:"ordernum"`
	Type      int       `json:"type"`
	Imgpath   string    `json:"imgpate"`
	Salt      string    `json:"salt"`
}
