package main

import (
	"github.com/ewenliu/grpc-demo/pb"
	"github.com/ewenliu/grpc-demo/server/service"
	"google.golang.org/grpc"
	"log"
	"net"
)

// protoc --go_out=plugins=grpc_py:. Product.proto


func main() {
	s := grpc.NewServer()
	pb.RegisterProdServiceServer(s, new(service.ProdService))

	listen, err := net.Listen("tcp", "127.0.0.1:8082")
	if err != nil{
		log.Fatal("tcp 监听失败", err)
	}

	_ = s.Serve(listen)
}
