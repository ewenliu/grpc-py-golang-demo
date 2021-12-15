package service

import (
	"context"
	"github.com/ewenliu/grpc-demo/pb"
)

type ProdService struct {

}

func (server *ProdService) GetProductStock(ctx context.Context, request *pb.ProductRequest) (*pb.ProductResponse, error) {
	return &pb.ProductResponse{ProdStock: 100}, nil
}
