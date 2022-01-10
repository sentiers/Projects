package models

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)

	os.Setenv("passwordHash", user.Password)
}

// func TestCreateUser(t *testing.T) {
// 	var userResult User

// 	err := config.DB.AutoMigrate(&User{})
// 	assert.NoError(t, err)

// 	user := User{
// 		Email:    "test@email.com",
// 		Password: os.Getenv("passwordHash"),
// 	}

// 	err = user.CreateUser()
// 	assert.NoError(t, err)

// 	config.DB.Where("email = ?", user.Email).Find(&userResult)

// 	config.DB.Unscoped().Delete(&user)

// 	assert.Equal(t, "test@email.com", userResult.Email)

// }

func TestCheckPassword(t *testing.T) {
	hash := os.Getenv("passwordHash")

	user := User{
		Password: hash,
	}

	err := user.CheckPassword("secret")
	assert.NoError(t, err)
}
