package main

import (
	"context"
	"log"

	pb "github.com/psanodiya94/grpc-go-learning/greet/proto"
)

func doGreet(client pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := client.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Pradumn",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %v\n", res.Result)
}
