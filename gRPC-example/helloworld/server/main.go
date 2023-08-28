package main

import (
	"context"
	helloWorldPb "github.com/upupnoah/IAM-GeekTime/gRPC-example/helloworld"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct {
	helloWorldPb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *helloWorldPb.HelloRequest) (*helloWorldPb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &helloWorldPb.HelloReply{
		Message: "Hello " + in.GetName(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	helloWorldPb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
