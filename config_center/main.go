package main

//
//import (
//	"fmt"
//	"github.com/nacos-group/nacos-sdk-go/clients"
//	"github.com/nacos-group/nacos-sdk-go/common/constant"
//	"github.com/nacos-group/nacos-sdk-go/vo"
//	"testProject/microservice/internal"
//)
//
//func main() {
//	nacosConf := internal.ViperConf.NacosConfig
//	serverConfigs := []constant.ServerConfig{
//		{
//			IpAddr: nacosConf.Host,
//			Port:   nacosConf.Port,
//		},
//	}
//	clientConfig := constant.ClientConfig{
//		//dev: b40636e1-e8eb-45d7-bbfc-b2a74dc591c2
//		//pro: a07dcd7a-d516-4d7e-864d-2c662239baba
//		NamespaceId:         nacosConf.Namespace,
//		TimeoutMs:           5000,
//		NotLoadCacheAtStart: true,
//		LogDir:              "nacos/log",
//		CacheDir:            "nacos/cache",
//		LogLevel:            "debug",
//	}
//
//	configClient, err := clients.NewConfigClient(
//		vo.NacosClientParam{
//			ClientConfig:  &clientConfig,
//			ServerConfigs: serverConfigs,
//		},
//	)
//	if err != nil {
//		panic(err)
//	}
//	content, err := configClient.GetConfig(vo.ConfigParam{
//		DataId: nacosConf.DataId,
//		Group:  nacosConf.Group,
//	})
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println(content)
//}
