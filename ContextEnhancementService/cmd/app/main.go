package main

import (
	pb "ContextEnhancementService/internal/fetcher/grpc/messagepb"
	server "ContextEnhancementService/internal/fetcher/grpc/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageServiceServer(s, &server.ContextServer{})

	log.Println("ContextEnhancementService :50051")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
