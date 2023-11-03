package main

import (
	"context"
	printhelloPB "github.com/upupnoah/iam-geektime/grpc-example/printhello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
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
	client := printhelloPB.NewPrintHelloClient(conn)

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()

	r, err := client.PrintHello(timeout, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("could not print hello world: %v", err)
	}
	log.Print(r.GetMessage())
}
