package main

import (
	"context"
	printhelloPB "github.com/upupnoah/iam-geektime/grpc-example/printhello"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
)

const (
	port = ":50051"
)

type service struct {
	printhelloPB.UnimplementedPrintHelloServer
}

func (s *service) PrintHello(context.Context, *emptypb.Empty) (*printhelloPB.HelloResponse, error) {
	return &printhelloPB.HelloResponse{
		Message: "Hello World",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 注册
	s := grpc.NewServer()
	printhelloPB.RegisterPrintHelloServer(s, &service{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
