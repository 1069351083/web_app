package settings

import (
	"fmt"
	"time"

	"go.uber.org/zap"

	"github.com/fsnotify/fsnotify"

	"github.com/spf13/viper"
)

var Conf = new(Config)

type Config struct {
	Name       string `mapstructure:"name"`
	Mode       string `mapstructure:"mode"`
	Port       int    `mapstructure:"port"`
	SingleKey  string `mapstructure:"single-key"`
	*LogConf   `mapstructure:"log"`
	*MysqlConf `mapstructure:"mysql"`
	*RedisConf `mapstructure:"redis"`
	*Casbin    `mapstructure:"casbin"`
}

type LogConf struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type MysqlConf struct {
	MysqlHost string `mapstructure:"host"`
	MysqlPort int    `mapstructure:"port"`
	UserName  string `mapstructure:"username"`
	PassWord  string `mapstructure:"password"`
	DbName    string `mapstructure:"dbName"`
}

type RedisConf struct {
	RedisHost   string        `mapstructure:"host"`
	RedisPort   int           `mapstructure:"port"`
	DB          int           `mapstructure:"db"`
	MaxIdle     int           `mapstructure:"MaxIdle"`
	MaxActive   int           `mapstructure:"MaxActive"`
	IdleTimeout time.Duration `mapstructure:"IdleTimeout "`
	PassWord    string        `mapstructure:"PassWord"`
}

func Init() (err error) {
	viper.SetConfigFile("config.yaml") // 指定配置文件
	viper.AddConfigPath(".")           // 指定查找配置文件的路径（这里使用相对路径）
	err = viper.ReadInConfig()         // 读取配置信息
	if err != nil {                    // 读取配置信息失败
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
		return
	}
	if err := viper.Unmarshal(Conf); err != nil {
		zap.L().Error("读取配置失败", zap.Error(err))
		fmt.Sprint(err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Sprint("配置文件被修改")
		viper.Unmarshal(Conf)
	})
	return
}

type Casbin struct {
	ModelPath string `mapstructure:"model_path" yaml:"model_path"`
}
