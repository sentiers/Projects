package middleware

import (
	"fmt"
	"github.com/serntiers/api-server/admin"
	"github.com/serntiers/api-server/config"
	"github.com/serntiers/api-server/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthzNoHeader(t *testing.T) {
	router := gin.Default()
	router.Use(Authz())

	router.GET("/company", controllers.GetAllCompanies)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/company", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 403, w.Code)
}

func TestAuthzInvalidTokenFormat(t *testing.T) {
	router := gin.Default()
	router.Use(Authz())

	router.GET("/company", controllers.GetAllCompanies)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/company", nil)
	req.Header.Add("Authorization", "test")

	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestAuthzInvalidToken(t *testing.T) {
	invalidToken := "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
	router := gin.Default()
	router.Use(Authz())

	router.GET("/company", controllers.GetAllCompanies)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/company", nil)
	req.Header.Add("Authorization", invalidToken)

	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
}

func TestValidToken(t *testing.T) {

	err := config.InitDatabase()
	assert.NoError(t, err)

	err = config.DB.AutoMigrate(&admin.User{})
	assert.NoError(t, err)

	user := admin.User{
		Email:    "test@email.com",
		Password: "secret",
	}

	jwtWrapper := admin.JwtWrapper{
		SecretKey:       "verysecretkey",
		Issuer:          "Alchera",
		ExpirationHours: 24,
	}

	token, err := jwtWrapper.GenerateToken(user.Email)
	assert.NoError(t, err)

	err = user.HashPassword(user.Password)
	assert.NoError(t, err)

	result := config.DB.Create(&user)
	assert.NoError(t, result.Error)

	router := gin.Default()
	router.Use(Authz())

	router.GET("/company", controllers.GetAllCompanies)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/company", nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))

	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)

	config.DB.Unscoped().Where("email = ?", user.Email).Delete(&admin.User{})
}
