package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"testProject/microservice/account_srv/biz"
	"testProject/microservice/account_srv/proto/pb"
	"testProject/microservice/internal"
)

func init() {
	internal.InitDB()
}

func main() {
	conf := internal.ViperConf
	addr := fmt.Sprintf("%s:%d", conf.AccountSrvConfig.Host, conf.AccountSrvConfig.Port)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		zap.S().Error("Account_srv 启动异常" + err.Error())
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAccountServiceServer(grpcServer, &biz.AccountServer{})

	//让grpc在consul注册，集成服务
	grpc_health_v1.RegisterHealthServer(grpcServer, health.NewServer())
	defaultConfig := api.DefaultConfig()
	defaultConfig.Address = addr
	client, err := api.NewClient(defaultConfig)
	if err != nil {
		zap.S().Error("Account_srv 构造client异常" + err.Error())
		panic(err)
	}
	check := &api.AgentServiceCheck{
		Interval:                       "1s",
		Timeout:                        "3s",
		GRPC:                           addr,
		DeregisterCriticalServiceAfter: "5s",
	}
	req := &api.AgentServiceRegistration{
		Name:    conf.AccountSrvConfig.SrvName,
		ID:      conf.AccountSrvConfig.SrvName,
		Check:   check,
		Port:    conf.AccountSrvConfig.Port,
		Tags:    conf.AccountSrvConfig.Tags,
		Address: addr,
	}
	err = client.Agent().ServiceRegister(req)
	if err != nil {
		zap.S().Error("Account_srv consul 构造异常" + err.Error())
		panic(err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
