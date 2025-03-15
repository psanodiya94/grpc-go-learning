package main

import (
	"fmt"
	"log"

	pb "github.com/psanodiya94/grpc-go-learning/greet/proto"
)

func (*Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Println("GreetManyTimes was invoked")

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
