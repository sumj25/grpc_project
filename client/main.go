package main

import (
	"log"

	pb "github.com/sumj25/grpc_go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials((insecure.NewCredentials())))

	if err != nil {
		log.Fatalf("did not connect on client side: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	names := &pb.NamesList{
		Names: []string{"Suman", "Rahul"},
	}
	log.Print("inside main function")

	// callSayHello(client)
	// callSayHelloServerStream(client, names)
	// callSayHelloClientStream(client, names)
	callSayHelloBidirectionalStream(client, names)

}
