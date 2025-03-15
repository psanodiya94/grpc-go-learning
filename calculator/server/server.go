package main

import pb "github.com/psanodiya94/grpc-go-learning/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
