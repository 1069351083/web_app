package mysql

import (
	"web_app/settings"

	"go.uber.org/zap"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func Init(cfg *settings.MysqlConf) (err error) {
	//dns := fmt.Sprintf("%s:%s@(%s:3306)/leige?charset=utf8mb4&parseTime=True&loc=Local",
	//	cfg.UserName, cfg.PassWord, cfg.MysqlHost, cfg.MysqlPort, cfg.DbName)
	DB, err = gorm.Open("mysql", "root:root@(127.0.0.1:3306)/leige?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		zap.L().Error("mysql connect failed", zap.Error(err))
		panic(err)
	}
	DB.SingularTable(true)
	return
}
