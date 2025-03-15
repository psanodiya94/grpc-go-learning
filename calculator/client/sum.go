package main

import (
	"context"
	"log"

	pb "github.com/psanodiya94/grpc-go-learning/calculator/proto"
)

func doSum(client pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := client.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  1,
		SecondNumber: 2,
	})
	if err != nil {
		log.Fatalf("Could not sum: %v\n", err)
	}
	log.Printf("Sum: %v\n", res.Result)
}
