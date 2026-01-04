package gemini

import (
	grpclient "ResponseAiService/internal/fetcher/grpc/client"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
	"log"
	"os"
)

func SendMessageInGemini(username string, content string, roomId string) {
	godotenv.Load()

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY_RESPONSE"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		log.Println(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(
			"Name: "+username+"Content: "+content+"An important answer to a user's question",
		),
		nil,
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result.Text())

	grpclient.SendNewMessageInChat(username, result.Text(), roomId)
}
