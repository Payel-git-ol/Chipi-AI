package server

import (
	g "ResponseAiService/internal/fetcher/gemini"
	pb "ResponseAiService/internal/fetcher/grpc/responsepb"
	"context"
	"log"

	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type ContextServer struct {
	pb.UnimplementedMessageServiceResponseServer
}

func (s *ContextServer) Message(ctx context.Context, req *pb.NewMessageContentResponse) (*emptypb.Empty, error) {
	log.Printf("Get message %s: %s", req.Username, req.Content, req.RoomId)

	g.SendMessageInGemini(req.Username, req.Content, req.RoomId)

	return &emptypb.Empty{}, nil
}
