package save

import (
	"ChipiAiAuth/pkg/database"
	"ChipiAiAuth/pkg/models"
	"ChipiAiAuth/pkg/models/request"
)

type User struct{}

func (u *User) UserService(req request.UserRequest) (*models.User, error) {
	newUser := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
	}

	user := database.Db.Where("username = ? AND email = ?", req.Username, req.Email).First(&newUser)
	if user.Error != nil {
		database.Db.Create(&newUser)
	}

	return &newUser, nil
}
