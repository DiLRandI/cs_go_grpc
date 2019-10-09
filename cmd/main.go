package main

import (
	"context"
	"log"
	"net"

	calcpb "github.com/dilrandi/calc-service/protobuff/calcpb/protobuff"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
}

// SayHello implements helloworld.GreeterServer
func (s *server) Add(ctx context.Context, r *calcpb.AddRequest) (*calcpb.AddResponse, error) {
	log.Printf("Received: %v And %v", r.No1, r.No2)
	res := r.No1 + r.No2
	return &calcpb.AddResponse{
		Result: res,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	calcpb.RegisterCalculatorServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
