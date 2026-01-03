package gemini

import (
	grpclient "ContextEnhancementService/internal/fetcher/grpc/client"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/genai"
	"log"
	"os"
)

func SendContentInGemini(username string, content string) {
	godotenv.Load()

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  os.Getenv("GEMINI_API_KEY"),
		Backend: genai.BackendGeminiAPI,
	})

	if err != nil {
		log.Println(err)
	}

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-2.5-flash",
		genai.Text(
			"Rewrite the user’s question into a clearer, more grammatically correct, and well‑formulated question. "+
				"Do not answer the question. Do not expand it. Do not add emotional or artistic interpretation. "+
				"Do not make it longer than necessary. Just rewrite it in a clean, neutral, professional form. "+
				"User "+username+" asks: \""+content+"\"",
		),
		nil,
	)

	if err != nil {
		log.Println(err)
	}

	fmt.Println(result.Text())

	grpclient.SendNewMessage(username, result.Text())

	log.Printf("Message from %s success send in ResponseService", username)
}
