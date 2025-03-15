package main

import pb "github.com/psanodiya94/grpc-go-learning/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
