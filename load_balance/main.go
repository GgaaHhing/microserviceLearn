package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"testProject/microservice/account_srv/proto/pb"
	"testProject/microservice/internal"
	"testProject/microservice/log"
	// 要使用consul的负载算法，还需要导入一个匿名包：
	_ "github.com/mbobakov/grpc-consul-resolver"
)

func main() {
	addr := fmt.Sprintf("%s:%d", internal.AppConf.ConsulConfig.Host, internal.AppConf.ConsulConfig.Port)
	// consul://{address}/{srvName}?wait=14
	dialAddr := fmt.Sprintf("consul://%s/%s?wait=14", addr, internal.AppConf.AccountSrvConfig.SrvName)
	conn, err := grpc.Dial(dialAddr, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(`{"load_balancing_policy": "round_robin"}`))
	if err != nil {
		log.Logger.Info(err.Error())
	}
	defer conn.Close()
	client := pb.NewAccountServiceClient(conn)
	res, err := client.GetAccountList(context.Background(), &pb.PagingRequest{
		PageNo:   1,
		PageSize: 3,
	})
	if err != nil {
		log.Logger.Info(err.Error())
	}
	for i, v := range res.AccountList {
		fmt.Println(i, v)
	}
}
