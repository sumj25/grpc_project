package main

import (
	"context"
	"log"

	pb "github.com/sumj25/grpc_go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {

	log.Print("inside sayHello function")

	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
