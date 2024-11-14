package internal

import (
	"fmt"
	"github.com/spf13/viper"
)

var ViperConf ViperConfig
var fileName = "microservice/dev-config.yaml"

func init() {
	v := viper.New()
	v.SetConfigFile(fileName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	//将配置文件中的内容解构成结构体，并赋值
	err = v.Unmarshal(&ViperConf)
	if err != nil {
		panic(err)
	}
	fmt.Println(ViperConf)
	fmt.Println("Redis Viper 初始化成功.....")

	InitRedis()
}

type ViperConfig struct {
	DBConfig         DBConfig         `mapstructure:"db"`
	RedisConfig      RedisConfig      `mapstructure:"redis"`
	ConsulConfig     ConsulConfig     `mapstructure:"consul"`
	AccountWebConfig AccountWebConfig `mapstructure:"account_web"`
	AccountSrvConfig AccountSrvConfig `mapstructure:"account_srv"`
}
