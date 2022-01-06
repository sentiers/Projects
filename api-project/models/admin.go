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

// create a admin user in the db
func (user *User) CreateUser() error {
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// encrypt admin user password
func (user *User) HashPassword(password string) error {
	// use bycrpt
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// check admin user password is correct
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
