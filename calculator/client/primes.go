package main

import (
	"context"
	"io"
	"log"

	pb "github.com/psanodiya94/grpc-go-learning/calculator/proto"
)

func doPrimes(client pb.CalculatorServiceClient) {
	log.Println("doPrimes was invoked")

	stream, err := client.Primes(context.Background(), &pb.PrimeRequest{
		Number: 12390392840,
	})
	if err != nil {
		log.Fatalf("error while calling Primes: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive primes: %v\n", err)
		}

		log.Printf("Primes: %v\n", res.Result)
	}
}
