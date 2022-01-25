package oauth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
)

// Google Oauth ------------------------------------------

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

var oauthGoogleConf *oauth2.Config

func OauthGoogleInit(clientid string, clientsecret string, redirecturl string) {
	oauthGoogleConf = &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientsecret,
		RedirectURL:  redirecturl,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleLogin(c *gin.Context) {
	b := make([]byte, 32)
	rand.Read(b)
	token := base64.RawStdEncoding.EncodeToString(b)
	url := oauthGoogleConf.AuthCodeURL(token)
	c.Redirect(http.StatusMovedPermanently, url)
}

func GoogleRedirect(c *gin.Context) (result []byte) {
	code := c.Query("code")
	token, err := oauthGoogleConf.Exchange(context.Background(), code)
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
	jsonMap := make(map[string]interface{})
	json.Unmarshal(contents, &jsonMap)
	c.JSON(200, jsonMap)
	return contents
}

// Github =====================================================

const oauthGithubUrlAPI = "https://api.github.com/user"

var oauthGithubConf *oauth2.Config

func OauthGithubInit(clientid string, clientsecret string, redirecturl string) {
	oauthGithubConf = &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientsecret,
		RedirectURL:  redirecturl,
		Scopes:       []string{"user"},
		Endpoint:     github.Endpoint,
	}
}

func GithubLogin(c *gin.Context) {
	b := make([]byte, 32)
	rand.Read(b)
	token := base64.RawStdEncoding.EncodeToString(b)
	url := oauthGithubConf.AuthCodeURL(token)
	c.Redirect(http.StatusMovedPermanently, url)
}

func GithubRedirect(c *gin.Context) (result []byte) {
	code := c.Query("code")
	token, err := oauthGithubConf.Exchange(context.Background(), code)
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
	jsonMap := make(map[string]interface{})
	json.Unmarshal(contents, &jsonMap)
	c.JSON(200, jsonMap)
	return contents
}

// Facebook =====================================================

const oauthFacebookUrlAPI = "https://graph.facebook.com/me?fields=name,email&access_token="

var oauthFacebookConf *oauth2.Config

func OauthFacebookInit(clientid string, clientsecret string, redirecturl string) {
	oauthFacebookConf = &oauth2.Config{
		ClientID:     clientid,
		ClientSecret: clientsecret,
		RedirectURL:  redirecturl,
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	}
}

func FacebookLogin(c *gin.Context) {
	b := make([]byte, 32)
	rand.Read(b)
	token := base64.RawStdEncoding.EncodeToString(b)
	url := oauthFacebookConf.AuthCodeURL(token)
	c.Redirect(http.StatusMovedPermanently, url)
}

func FacebookRedirect(c *gin.Context) (result []byte) {
	code := c.Query("code")
	token, err := oauthFacebookConf.Exchange(context.Background(), code)
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
	jsonMap := make(map[string]interface{})
	json.Unmarshal(contents, &jsonMap)
	c.JSON(200, jsonMap)
	return contents
}
