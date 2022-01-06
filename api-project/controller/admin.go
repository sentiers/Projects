package controller

import (
	"log"
	"manager/auth"
	"manager/config"
	"manager/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// login body
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// token response
type LoginResponse struct {
	Token string `json:"token"`
}

// login for admin user
func Login(c *gin.Context) {
	var payload LoginPayload
	var user models.User

	err := c.ShouldBindJSON(&payload)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	result := config.DB.Where("email = ?", payload.Email).First(&user)

	// if couldn't find the matching email
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	// if the password is not matching
	err = user.CheckPassword(payload.Password)
	if err != nil {
		log.Println(err)
		c.JSON(401, gin.H{
			"msg": "invalid user credentials",
		})
		c.Abort()
		return
	}

	// success

	// start generating token ...

	jwtWrapper := auth.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "Alchera",
		ExpirationHours: 24,
	}

	signedToken, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error signing token",
		})
		c.Abort()
		return
	}

	tokenResponse := LoginResponse{
		Token: signedToken,
	}

	c.JSON(200, tokenResponse)

	return
}

// register admin user
func Signup(c *gin.Context) {
	var user models.User

	// BindJSON -> ShouldBindJSON because of the [WARNING]
	err := c.ShouldBindJSON(&user)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"msg": "invalid json",
		})
		c.Abort()
		return
	}

	err = user.HashPassword(user.Password)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"msg": "error hashing password",
		})
		c.Abort()
		return
	}

	err = user.CreateUser()
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"msg": "error creating user",
		})
		c.Abort()
		return
	}
	c.JSON(200, user) // success
}
