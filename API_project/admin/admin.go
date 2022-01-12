package admin

import (
	"log"
	"manager/config"

	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	jwt "github.com/dgrijalva/jwt-go"
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

// Signup and Login--------------------------------------

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
	var user User

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

	jwtWrapper := JwtWrapper{
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
	var user User

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

// Jwt authorization --------------------------------------

// wrap the signing key and the issuer
type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

// add email as a claim to the token
type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

// JwtWrapper's method
// generate a jwt token
func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {

	claims := &JwtClaim{ // use JwtClaim structure
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}

	return
}

// JwtWrapper's method
// validate the jwt token
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) { // anonymous function
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}

	// if claims.ExpiresAt < time.Now().Local().Unix() {
	// 	err = errors.New("JWT is expired")
	// 	return
	// }

	return

}
