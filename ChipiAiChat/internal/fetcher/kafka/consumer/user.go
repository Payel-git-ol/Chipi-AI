package consumer

import (
	"ChipiAiChat/internal/core/service/process_message"
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetMessageUser(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "user-in-chat",
		GroupID: "user-in-chat",
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err)
		}

		message, err := process_message.ProcessMessage(m.Value)
		if err != nil {
			log.Println(err)
		}

		log.Println(message)
	}
}
