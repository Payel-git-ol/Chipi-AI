package client

import (
	pb "ContextEnhancementService/internal/fetcher/grpc/responsepb"
	"context"
	"google.golang.org/grpc"
	"log"
)

func SendNewMessage(username string, content string) {
	ctx := context.Background()
	conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())

	if err != nil {
		log.Printf("could not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewMessageServiceResponseClient(conn)

	_, err = client.Message(ctx, &pb.NewMessageContentResponse{
		Username: username,
		Content:  content,
	})

	if err != nil {
		log.Printf("could not send: %v", err)
	}

	log.Printf("Message from %s success send in ContextEnhancementService", username)
}
