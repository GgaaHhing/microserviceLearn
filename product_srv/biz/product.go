package biz

import (
	"context"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"testProject/microservice_part2/proto/google/pb"
)

type ProductServer struct {
	pb.ProductServiceServer
}

func (p ProductServer) ProductList(ctx context.Context, req *pb.ProductionConditionReq) (*pb.ProductRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) BatchGetProduct(ctx context.Context, req *pb.BatchProductIdReq) (*pb.ProductRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) CreateProduct(ctx context.Context, item *pb.CreateProductItem) (*pb.ProductItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) DeleteProduct(ctx context.Context, item *pb.DeleteProductItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) UpdateProduct(ctx context.Context, item *pb.CreateProductItem) (*emptypb.Empty, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) GetProductDetail(ctx context.Context, req *pb.ProductItemReq) (*pb.ProductItemRes, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductServer) mustEmbedUnimplementedProductServiceServer() {
	//TODO implement me
	panic("implement me")
}
