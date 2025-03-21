package main

import (
	"context"
	"io"
	"testing"

	pb "github.com/psanodiya94/grpc-go-learning/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func TestGreetEveryone(t *testing.T) {
	creds := grpc.WithTransportCredentials(insecure.NewCredentials())
	conn, err := grpc.NewClient("bufnet", grpc.WithContextDialer(bufDialer), creds)

	if err != nil {
		t.Fatalf("Failed to dial bufnet: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		t.Errorf("Error while creating stream: %v", err)
	}

	requests := []*pb.GreetRequest{
		{
			FirstName: "Pradumn",
		},
		{
			FirstName: "Shivika",
		},
		{
			FirstName: "Test",
		},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range requests {
			stream.Send(req)
		}

		stream.CloseSend()
	}()

	go func() {
		idx := 0

		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				t.Errorf("Error while receiving: %v", err)
				break
			}

			expected := "Hello " + requests[idx].FirstName + "!"
			if res.Result != expected {
				t.Errorf("Expected \"%s\", got: \"%s\"", expected, res.Result)
			}

			idx++
		}

		close(waitc)
	}()

	<-waitc
}
