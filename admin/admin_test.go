package admin

import (
	"bytes"
	"encoding/json"
	"github.com/sentiers/api-server/config"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// =============================================================
func TestCreateUser(t *testing.T) {
	var userResult User

	err := config.InitDatabase()
	if err != nil {
		t.Error(err)
	}

	err = config.DB.AutoMigrate(&User{})
	assert.NoError(t, err)

	user := User{
		Email:    "test@email.com",
		Password: os.Getenv("passwordHash"),
	}

	err = user.CreateUser()
	assert.NoError(t, err)

	config.DB.Where("email = ?", user.Email).Find(&userResult)

	config.DB.Unscoped().Delete(&user)

	assert.Equal(t, "test@email.com", userResult.Email)

}
func TestHashPassword(t *testing.T) {
	user := User{
		Password: "secret",
	}

	err := user.HashPassword(user.Password)
	assert.NoError(t, err)

	os.Setenv("passwordHash", user.Password)
}

func TestCheckPassword(t *testing.T) {
	hash := os.Getenv("passwordHash")

	user := User{
		Password: hash,
	}

	err := user.CheckPassword("secret")
	assert.NoError(t, err)
}

// ===================================================================
func TestGenerateToken(t *testing.T) {
	jwtWrapper := JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "Alchera",
		ExpirationHours: 24,
	}

	generatedToken, err := jwtWrapper.GenerateToken("jwt@email.com")
	assert.NoError(t, err)

	os.Setenv("testToken", generatedToken)
}

func TestValidateToken(t *testing.T) {
	encodedToken := os.Getenv("testToken")

	jwtWrapper := JwtWrapper{
		SecretKey: "verysecretkey",
		Issuer:    "Alchera",
	}

	claims, err := jwtWrapper.ValidateToken(encodedToken)
	assert.NoError(t, err)

	assert.Equal(t, "jwt@email.com", claims.Email)
	assert.Equal(t, "Alchera", claims.Issuer)
}

// ================================================================

func TestSignUp(t *testing.T) {
	var actualResult User

	user := User{
		Email:    "jwt@email.com",
		Password: "secret",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = config.InitDatabase()
	assert.NoError(t, err)
	config.DB.AutoMigrate(&User{})

	Signup(c)

	assert.Equal(t, 200, w.Code)

	err = json.Unmarshal(w.Body.Bytes(), &actualResult)
	assert.NoError(t, err)

	assert.Equal(t, user.Email, actualResult.Email)
}

func TestSignUpInvalidJSON(t *testing.T) {
	user := "test"

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/signup", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	Signup(c)

	assert.Equal(t, 400, w.Code)
}

func TestLogin(t *testing.T) {
	user := LoginPayload{
		Email:    "jwt@email.com",
		Password: "secret",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = config.InitDatabase()
	assert.NoError(t, err)

	config.DB.AutoMigrate(&User{})

	Login(c)

	assert.Equal(t, 200, w.Code)

}

func TestLoginInvalidJSON(t *testing.T) {
	user := "test"

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	Login(c)

	assert.Equal(t, 400, w.Code)
}

func TestLoginInvalidCredentials(t *testing.T) {
	user := LoginPayload{
		Email:    "jwt@email.com",
		Password: "invalid",
	}

	payload, err := json.Marshal(&user)
	assert.NoError(t, err)

	request, err := http.NewRequest("POST", "/login", bytes.NewBuffer(payload))
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(w)
	c.Request = request

	err = config.InitDatabase()
	assert.NoError(t, err)

	config.DB.AutoMigrate(&User{})

	Login(c)

	assert.Equal(t, 401, w.Code)

	config.DB.Unscoped().Where("email = ?", user.Email).Delete(&User{})
}
