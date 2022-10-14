package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"client/proto/greeter"
)

func main(){

	// 連接服務器
	grpcClient,err := grpc.Dial("127.0.0.1:8080",grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err !=nil{
		fmt.Println(err)
	}
	// 註冊客戶端
	client :=greeter.NewGreeterClient(grpcClient)
	res,err :=client.SayHello(context.Background(),&greeter.HelloReq{
		Name: "張三",
	})

	fmt.Println(res.Message)
}
