package main

import (
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/psanodiya94/grpc-go-learning/greet/proto"
)

var greetWithDeadlineTime time.Duration = 1 * time.Second

var addr string = "0.0.0.0:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Printf("Server listening at %v\n", addr)

	opts := []grpc.ServerOption{}
	tls := true // Enable TLS

	if tls {
		// Create a new TLS server
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	serv := grpc.NewServer(opts...)

	pb.RegisterGreetServiceServer(serv, &Server{})

	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
