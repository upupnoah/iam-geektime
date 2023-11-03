package main

import (
	"context"
	helloWorldPb "github.com/upupnoah/iam-geektime/grpc-example/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// 创建一个 gRPC 连接，用来跟服务端进行通信
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Fatalf("did not close: %v", err)
		}
	}(conn)

	// 创建客户端 stub，用来执行 RPC 请求
	c := helloWorldPb.NewGreeterClient(conn)

	// Contact the server and print out its response
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// 创建一个带有超时的 context
	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	// 通过 c.SayHello调用远端的 SayHello 接口
	r, err := c.SayHello(timeout, &helloWorldPb.HelloRequest{
		Name: name,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	//fmt.Printf("Greeting: %s", r.GetMessage())

}
