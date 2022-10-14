package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"greeter/proto/greeter"
	"net"
)
// 1.定義遠程調用的結構體和方法，需要實現Greeterserver的街口

type Hello struct{}

// 方法名稱要跟proto文件的方法名稱一樣
func (this Hello) SayHello(c context.Context,req *greeter.HelloReq)(*greeter.HelloRes,error){
	fmt.Println(req)
	return &greeter.HelloRes{
		Message: "你好"+ req.Name,
	},nil
}

func main(){
	// 1.初始化一個grpc對象
	grpcServer :=grpc.NewServer()
	// 2.註冊服務
	greeter.RegisterGreeterServer(grpcServer,&Hello{})
	// 3.監聽端口
	Listener,err := net.Listen("tcp","127.0.0.1:8080")
	if err !=nil{
		fmt.Println(err)
	}
	//監聽退出關閉監聽
	defer Listener.Close()
	// 4.啟動服務
	grpcServer.Serve(Listener)
}


