package main

import (
	"context"
	"fmt"
	"github.com/ewenliu/grpc-demo/pb"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err:= grpc.Dial("127.0.0.1:8083", grpc.WithInsecure())
	if err!=nil{
		log.Fatal(err)
	}
	defer func() {
		_ = conn.Close()
	}()
	client := pb.NewProdServiceClient(conn)
	resp, err := client.GetProductStock(context.TODO(), &pb.ProductRequest{ProdId: 123})
	fmt.Println(err)
	fmt.Println(resp)
}
