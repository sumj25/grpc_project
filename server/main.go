package main

import (
	"log"
	"net"

	pb "github.com/sumj25/grpc_go/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

const (
	port = ":8080"
)

func main() {
	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})

	log.Printf("server started at %v", lis.Addr())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("error during start %v", err)
	}

}
