package main

import pb "github.com/psanodiya94/grpc-go-learning/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
