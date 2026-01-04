package grpcserver

import (
	"ContextEnhancementService/internal/fetcher/gemini"
	"context"
	"log"

	pb "ContextEnhancementService/internal/fetcher/grpc/messagepb"
)

type ContextServer struct {
	pb.UnimplementedMessageServiceServer
}

func (s *ContextServer) Message(ctx context.Context, req *pb.NewMessageContent) (*pb.Empty, error) {
	log.Printf("Get message %s: %s", req.Username, req.Content, req.RoomId)

	gemini.SendContentInGemini(req.Username, req.Content, req.RoomId)

	return &pb.Empty{}, nil
}
