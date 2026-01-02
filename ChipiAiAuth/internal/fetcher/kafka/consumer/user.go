package consumer

import (
	"ChipiAiAuth/internal/core/service/process_message"
	"ChipiAiAuth/internal/core/service/save"
	"context"
	"github.com/segmentio/kafka-go"
	"log"
	"sync"
)

func GetMessageNewUser(wg *sync.WaitGroup) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:29092"},
		Topic:   "register",
		GroupID: "user",
	})

	var saveUser save.User

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

		_, err = saveUser.UserService(message)
		if err != nil {
			log.Println(err)
		}
	}
}
