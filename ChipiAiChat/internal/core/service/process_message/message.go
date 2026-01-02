package process_message

import (
	"ChipiAiChat/pkg/models/request"
	"encoding/json"
	"fmt"
)

func ProcessMessage(data []byte) (request.UserRequest, error) {
	fmt.Println("Consumer started")

	var req request.UserRequest
	if err := json.Unmarshal(data, &req); err != nil {
		return request.UserRequest{}, err
	}

	return req, nil
}
