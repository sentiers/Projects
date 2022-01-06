package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

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
