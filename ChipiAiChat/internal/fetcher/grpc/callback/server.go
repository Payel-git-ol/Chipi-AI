package callback

import (
	"ChipiAiChat/internal/core/service/chat"
	"ChipiAiChat/internal/fetcher/grpc/callbackpb"
	"ChipiAiChat/pkg/database"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CallbackServer struct {
	callbackpb.UnimplementedAiCallbackServer
}

func (s *CallbackServer) SendAiMessage(ctx context.Context, req *callbackpb.AiMessage) (*emptypb.Empty, error) {
	ws := chat.Connections[req.Username]
	if ws != nil {
		ws.WriteMessage(1, []byte(req.Content))
		err := database.UpdateMessageContent(req.RoomId, req.Username, req.Content)
		if err != nil {
			return nil, err
		}
	}
	return &emptypb.Empty{}, nil
}
