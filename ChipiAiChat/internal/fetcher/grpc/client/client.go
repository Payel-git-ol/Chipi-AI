package client

import (
	pb "ChipiAiChat/internal/fetcher/grpc/messagepb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func SendContent(username string, content string) {
	ctx := context.Background()
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		log.Printf("could not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewMessageServiceClient(conn)

	_, err = client.Message(ctx, &pb.NewMessageContent{
		Username: username,
		Content:  content,
	})

	if err != nil {
		log.Printf("could not send: %v", err)
	}

	log.Printf("Message from %s success send in ContextEnhancementService", username)
}
