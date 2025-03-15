package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/psanodiya94/grpc-go-learning/calculator/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Printf("Server listening at %v\n", addr)

	serv := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(serv, &Server{})

	reflection.Register(serv)

	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
