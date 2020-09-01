package service

import (
	"errors"
	"web_app/dao/mysql"
	"web_app/model"
	"web_app/settings"

	"github.com/casbin/casbin"
	gormadapter "github.com/casbin/gorm-adapter"
)

func Casbin() *casbin.Enforcer {
	a := gormadapter.NewAdapterByDB(mysql.DB)
	e := casbin.NewEnforcer(settings.Conf.ModelPath, a)
	_ = e.LoadPolicy()
	return e
}

func UpdateCasbin(name string, casbinInfos []model.CasbinInfo) error {
	ClearCasbin(0, name)
	for _, v := range casbinInfos {
		cm := model.CasbinModel{
			ID:     0,
			Ptype:  "p",
			Name:   name,
			Path:   v.Path,
			Method: v.Method,
		}
		addflag := AddCasbin(cm)
		if addflag == false {
			return errors.New("存在相同api,添加失败,请联系管理员")
		}
	}
	return nil
}

func GetPolicyPathByAuthorityId(name string) []model.CasbinInfo {
	enforcer := Casbin()
	var infos []model.CasbinInfo
	policys := enforcer.GetFilteredPolicy(0, name)
	for _, p := range policys {
		info := model.CasbinInfo{
			Path:   p[1],
			Method: p[2],
		}
		infos = append(infos, info)
	}
	return infos

}

func ClearCasbin(v int, p ...string) bool {
	e := Casbin()
	return e.RemoveFilteredPolicy(v, p...)

}

func AddCasbin(cm model.CasbinModel) bool {
	e := Casbin()
	return e.AddPolicy(cm.Name, cm.Path, cm.Method)
}
