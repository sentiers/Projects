package middleware

import (
	"strings"

	"github.com/sentiers/api-server/admin"

	"github.com/gin-gonic/gin"
)

// ensure if user has validate token
func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {

		// try reading HTTP Header
		clientToken := c.Request.Header.Get("Authorization")
		// when there is no token in authorization header
		if clientToken == "" {
			c.JSON(403, "No Authorization header provided")
			c.Abort()
			return
		}

		// extract token from string "Bearer token"
		extractedToken := strings.Split(clientToken, "Bearer ")

		if len(extractedToken) == 2 {
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			// when token format is incorrect
			c.JSON(400, "Incorrect Format of Authorization Token")
			c.Abort()
			return
		}

		jwtWrapper := admin.JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "Alchera",
		}

		claims, err := jwtWrapper.ValidateToken(clientToken)
		// when token is invalid
		if err != nil {
			c.JSON(401, err.Error())
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}
