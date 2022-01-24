package admin

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"manager/config"
	"net/http"

	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
	"gorm.io/gorm"

	jwt "github.com/dgrijalva/jwt-go"
)

// Local User Model=====================================================

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
	return
}

// Oauth basic functions =====================================================

var oauthConf *oauth2.Config

func getToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.RawStdEncoding.EncodeToString(b)
}

func getLoginURL(state string) string {
	return oauthConf.AuthCodeURL(state)
}

// Google User =====================================================

type User_Google struct {
	Email string `gorm:"primarykey; type:varchar(255); not null" json:"email"`
	Name  string `json:"name"`
}

func (c *User_Google) TableName() string {
	return "user_google"
}

// create a admin user in the db if the user is not exist
func (user *User_Google) CreateUser_Google() error {

	exist := config.DB.First(&user, "email = ?", user.Email)
	if exist.Error == nil { // skip if user is already exist
		return nil
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Google Oauth ------------------------------------------

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

func oauthGoogleInit() {
	oauthConf = &oauth2.Config{
		ClientID:     "clientid",
		ClientSecret: "clientsecret",
		RedirectURL:  "http://localhost:8080/google/redirect",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleLogin(c *gin.Context) {
	oauthGoogleInit()
	token := getToken()
	url := getLoginURL(token)
	c.Redirect(http.StatusMovedPermanently, url)
}

func GoogleRedirect(c *gin.Context) {

	code := c.Query("code")

	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}

	// save user info in database
	var user User_Google
	json.Unmarshal(contents, &user)

	if err := user.CreateUser_Google(); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// start generating token

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

// Github User =====================================================

type User_Github struct {
	Email string `gorm:"primarykey; type:varchar(255); not null" json:"email"`
	Name  string `json:"name"`
}

func (c *User_Github) TableName() string {
	return "user_github"
}

// create a admin user in the db if the user is not exist
func (user *User_Github) CreateUser_Github() error {

	exist := config.DB.First(&user, "email = ?", user.Email)
	if exist.Error == nil { // skip if user is already exist
		return nil
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Github Oauth ------------------------------------------

const oauthGithubUrlAPI = "https://api.github.com/user"

func oauthGithubInit() {
	oauthConf = &oauth2.Config{
		ClientID:     "clientid",
		ClientSecret: "clientsecret",
		RedirectURL:  "http://localhost:8080/github/redirect",
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}
}

func GithubLogin(c *gin.Context) {
	oauthGithubInit()
	token := getToken()
	url := getLoginURL(token)
	c.Redirect(http.StatusMovedPermanently, url)
}

func GithubRedirect(c *gin.Context) {

	code := c.Query("code")

	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}

	// Get request to a set URL
	request, err := http.NewRequest("GET", oauthGithubUrlAPI, nil)
	if err != nil {
		log.Panic("API Request creation failed")
	}
	// set authorization header: token XXXXXXXXXXXXXXXXXXXXXXXXXXX
	authHeader := fmt.Sprintf("token %s", token.AccessToken)
	request.Header.Set("Authorization", authHeader)

	// Make the request
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}

	// save user info in database
	var user User_Github
	json.Unmarshal(contents, &user)

	if err := user.CreateUser_Github(); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// start generating token

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

// Facebook User =====================================================

type User_Facebook struct {
	Email string `gorm:"primarykey; type:varchar(255); not null" json:"email"`
	Name  string `json:"name"`
}

func (c *User_Facebook) TableName() string {
	return "user_facebook"
}

// create a admin user in the db if the user is not exist
func (user *User_Facebook) CreateUser_Facebook() error {

	exist := config.DB.First(&user, "email = ?", user.Email)
	if exist.Error == nil { // skip if user is already exist
		return nil
	}
	result := config.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Facebook Oauth ------------------------------------------

const oauthFacebookUrlAPI = "https://graph.facebook.com/me?fields=name,email&access_token="

func oauthFacebookInit() {
	oauthConf = &oauth2.Config{
		ClientID:     "clientid",
		ClientSecret: "clientsecret",
		RedirectURL:  "{ngrok_url}/facebook/redirect",
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}
}

func FacebookLogin(c *gin.Context) {
	oauthFacebookInit()
	token := getToken()
	url := getLoginURL(token)
	c.Redirect(http.StatusMovedPermanently, url)
}

func FacebookRedirect(c *gin.Context) {

	code := c.Query("code")

	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}

	response, err := http.Get(oauthFacebookUrlAPI + token.AccessToken)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c.JSON(403, gin.H{"Message": err.Error()})
		return
	}

	// save user info in database
	var user User_Facebook
	json.Unmarshal(contents, &user)

	if err := user.CreateUser_Facebook(); err != nil {
		log.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	// start generating token

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
