package model

import (
	"github.com/google/uuid"
	database "github.com/sam-app/hackernews/packages/db/postgress"
)

type User struct {
	ID       string `json:"id" gorm:"primary_key;default:uuid_generate_v4()"`
	Name     string `json:"name"`
	Posts    []Post `json:"posts" gorm:"foreignkey:UserID"`
	Username string `json:"username" gorm:"unique not null"`
	Password string `json:"password" gorm:"not null"`
}

func (user *User) Save() (*User, error) {
	uuid := uuid.New().String()
	result := database.Db.Create(&User{ID: uuid, Username: user.Username, Password: user.Password, Name: user.Name})
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func GetUserById(id string) (*User, error) {
	var user User
	result := database.Db.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
