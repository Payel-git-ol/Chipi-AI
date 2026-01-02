package producer

import (
	"ChipiAiAuth/pkg/models"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

func SendMessageUser(topic string, msg models.User) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   topic,
	})

	defer w.Close()

	jsonData, err := json.Marshal(msg)
	if err != nil {
		log.Println(err)
	}

	err = w.WriteMessages(context.Background(), kafka.Message{
		Value: jsonData,
	})

	if err != nil {
		log.Println(err)
	}

	log.Println("Message sent: ", jsonData)
}
