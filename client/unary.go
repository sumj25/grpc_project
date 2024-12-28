package main

import (
	"context"
	"log"
	"time"

	pb "github.com/sumj25/grpc_go/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Print("inside CallsayHello function")
	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("error during say hello %v", err)
	}

	log.Printf("%s", res.Message)
}
