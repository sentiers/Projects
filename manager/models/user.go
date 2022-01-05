package models

import (
	"manager/config"

	"golang.org/x/crypto/bcrypt"
)

// User Model=====================================================

type User struct {
	Email    string `gorm:"primarykey; type:varchar(255); not null" json:"email"`
	Password string `json:"password"`
}

func (c *User) TableName() string {
	return "user"
}

// func CreateUser(user *User) (err error) {
// 	if err = config.DB.Create(user).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// CreateUserRecord creates a user record in the database
func (user *User) CreateUserRecord() error {
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// HashPassword encrypts user password
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

// CheckPassword checks user password
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}

	return nil
}
