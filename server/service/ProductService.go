package service

import (
	"context"
	"fmt"
	"github.com/ewenliu/grpc-demo/pb"
	"time"
)

type ProdService struct {

}

func (server *ProdService) GetProductStock(ctx context.Context, request *pb.ProductRequest) (*pb.ProductResponse, error) {
	fmt.Println(fmt.Sprintf("%s: this is golang server", time.Now()))
	return &pb.ProductResponse{ProdStock: 100}, nil
}
