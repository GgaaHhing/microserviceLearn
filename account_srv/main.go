package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"google.golang.org/grpc"
	"net"
	"testProject/microservice/account_srv/biz"
	"testProject/microservice/account_srv/proto/pb"
	"testProject/microservice/internal"
	"testProject/microservice/log"
)

func init() {
	internal.InitDB()
}

func main() {
	conf := internal.ViperConf
	srvAddress := fmt.Sprintf("%s:%d", conf.AccountSrvConfig.Host, conf.AccountSrvConfig.Port)

	// 创建一个新的 Consul 客户端
	defaultConfig := api.DefaultConfig()
	defaultConfig.Address = fmt.Sprintf("%s:%d", conf.ConsulConfig.Host, conf.ConsulConfig.Port)

	consulClient, err := api.NewClient(defaultConfig)
	if err != nil {
		log.Logger.Error("创建 consulClient失败： " + err.Error())
		panic(err)
	}

	req := &api.AgentServiceRegistration{
		Address: conf.AccountSrvConfig.Host,
		Port:    conf.AccountSrvConfig.Port,
		Name:    conf.AccountSrvConfig.SrvName,
		ID:      conf.AccountSrvConfig.SrvName,
		Tags:    conf.AccountSrvConfig.Tags,
	}

	err = consulClient.Agent().ServiceRegister(req)
	if err != nil {
		log.Logger.Error("GRPC 部署 consul失败：" + err.Error())
		panic(err)
	}

	// 监听并服务 gRPC 请求
	lis, err := net.Listen("tcp", ":9095")
	if err != nil {
		log.Logger.Error(srvAddress + "监听失败：" + err.Error())
		panic(err)
	}
	fmt.Println("gRPC 正在监听 :  " + srvAddress)

	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &biz.AccountServer{}) // 替换为你的服务接口和服务器实例

	if err := s.Serve(lis); err != nil {
		log.Logger.Error("GRPC 部署失败: " + err.Error())
		panic(err)
	}
}
