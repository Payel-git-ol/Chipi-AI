package producer

import (
	"ApiGateway/pkg/models/request"
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log"
)

func SendMessageNewUser(topic string, msg request.UserRequest) {
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
}
