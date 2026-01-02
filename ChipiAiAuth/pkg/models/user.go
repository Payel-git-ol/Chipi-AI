package models

type User struct {
	Id       uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
