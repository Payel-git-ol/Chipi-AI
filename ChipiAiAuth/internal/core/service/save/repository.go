package save

import (
	"ChipiAiAuth/pkg/models"
	"ChipiAiAuth/pkg/models/request"
)

type UserRepository interface {
	UserService(req request.UserRequest) (*models.User, error)
}
