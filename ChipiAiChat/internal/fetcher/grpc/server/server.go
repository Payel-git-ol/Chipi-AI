package server

import (
	pb "ChipiAiChat/internal/fetcher/grpc/chatpb"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

type ContextServer struct {
	pb.UnimplementedMessageServiceChatServer
}

func (s *ContextServer) Message(ctx context.Context, req *pb.NewMessageContentInChat) (*emptypb.Empty, error) {
	log.Printf("Get message %s: %s", req.Username, req.Content)

	return &emptypb.Empty{}, nil
}
