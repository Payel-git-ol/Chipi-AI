package client

import (
	call "ResponseAiService/internal/fetcher/grpc/callbackpb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func SendNewMessageInChat(username string, content string) {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())

	if err != nil {
		log.Printf("could not connect: %v", err)
	}

	defer conn.Close()

	callbackClient := call.NewAiCallbackClient(conn)

	_, err = callbackClient.SendAiMessage(ctx, &call.AiMessage{
		Username: username,
		Content:  content,
	})

	if err != nil {
		log.Printf("could not send: %v", err)
	}

	log.Printf("Message from %s success send in ContextEnhancementService", username)
}
