package main

import (
	"log"
	"net"

	pb "grpc-demo/service/shopping"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Listen error is %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterShoppingServiceServer(s, pb.NewShoppingService())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Start server is error :%v", err)
	}
}
