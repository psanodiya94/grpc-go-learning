package main

import (
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "github.com/psanodiya94/grpc-go-learning/greet/proto"
)

var addr string = "localhost:50051"

func main() {
	tls := true // Enable TLS
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			log.Fatalf("Failed loading certificates: %v\n", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()

	log.Printf("Connected to %v\n", addr)

	client := pb.NewGreetServiceClient(conn)

	doGreet(client)
	// doGreetManyTimes(client)
	// doLongGreet(client)
	// doGreetEveryone(client)
	// doGreetWithDeadline(client, 5)
}
