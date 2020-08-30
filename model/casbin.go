package model

type CasbinModel struct {
	ID     uint   `json:"id" gorm:"column:_id"`
	Ptype  string `json:"ptype" gorm:"column:ptype"`
	Name   string `json:"name" gorm:"column:v0"`
	Path   string `json:"path" gorm:"column:v1"`
	Method string `json:"method" gorm:"column:v2"`
}

type CasbinInReceive struct {
	Name        string       `json:"name"`
	CasbinInfos []CasbinInfo `json:"casbinInfos"`
}

type CasbinInfo struct {
	Path   string `json:"path"`
	Method string `json:"method"`
}
