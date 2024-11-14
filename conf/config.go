package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

type AppConfig struct {
	JWTConfig JWTConfig `mapstructure:"jwt_op"`
}

var AppConf AppConfig

func init() {
	v := viper.New()
	configName := "microservice/dev-config.yaml"
	v.SetConfigFile(configName)
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = v.Unmarshal(&AppConf)
	if err != nil {
		panic(err)
	}
	fmt.Println("AppConfig 初始化成功....")
}
