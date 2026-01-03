package main

import (
	pb "ResponseAiService/internal/fetcher/grpc/responsepb"
	"ResponseAiService/internal/fetcher/grpc/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterMessageServiceResponseServer(s, &server.ContextServer{})

	log.Println("ContextEnhancementService :50052")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
