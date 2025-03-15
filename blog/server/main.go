package main

import (
	"context"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"

	pb "github.com/psanodiya94/grpc-go-learning/blog/proto"
)

var addr string = "0.0.0.0:50051"

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to mongo!")

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	log.Printf("Server listening at %v\n", addr)

	serv := grpc.NewServer()

	pb.RegisterBlogServiceServer(serv, &Server{})

	if err := serv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
