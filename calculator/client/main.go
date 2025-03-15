package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/psanodiya94/grpc-go-learning/calculator/proto"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	log.Printf("Connected to %v\n", addr)

	client := pb.NewCalculatorServiceClient(conn)

	doSum(client)
	doPrimes(client)
	doAvg(client)
	doMax(client)
	doSqrt(client, 10)
	doSqrt(client, -2)
}
