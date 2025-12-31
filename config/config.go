package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	Mysql *Mysql `mapstructure:"mysql" json:"mysql"`
	Redis *Redis `mapstructure:"redis" json:"redis"`
	JWT   *JWT   `mapstructure:"jwt" json:"jwt"`
}

var Conf = new(Config)

func InitConfig() {
	dir, err := os.Getwd()
	if err != nil {
		panic(fmt.Errorf("读取目录失败: %s", err))
	}
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AddConfigPath(dir + "/config/")
	// 读取配置信息
	err = viper.ReadInConfig()

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(Conf); err != nil {
			panic(fmt.Errorf("初始化配置文件失败:%s", err))
		}
	})

	if err = viper.Unmarshal(Conf); err != nil {
		panic(fmt.Errorf("初始化配置文件失败:%s", err))
	}
	fmt.Println(Conf.Mysql, Conf.Redis, Conf.JWT)
}
