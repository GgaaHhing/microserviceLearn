package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"testProject/microservice/account_srv/biz"
	"testProject/microservice/account_srv/internal"
	"testProject/microservice/account_srv/proto/pb"
)

func init() {
	internal.InitDB()
}

func main() {
	ip := flag.String("ip", "127.0.0.1", "ip 输入")
	port := flag.Int("port", 9095, "port 输入")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterAccountServiceServer(grpcServer, &biz.AccountServer{})
	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}
