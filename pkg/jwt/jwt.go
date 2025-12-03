package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// token expiration duration
const TokenExpireDuration = time.Hour * 2

// secret key
var mySecret = []byte("cugxc.top")

type MyClaims struct {
	UserId               uint64 `json:"user_id"`
	Username             string `json:"username"`
	jwt.RegisteredClaims        // Embedding RegisteredClaims
}

func GenToken(userid uint64, username string) (string, error) {
	// Create our own claims
	claims := MyClaims{
		userid,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpireDuration)),
			Issuer:    "lelForum", // Issuer
		},
	}
	// Create token object using specified signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Sign with secret and get the complete encoded token string
	return token.SignedString(mySecret)
}

// ParseToken parses JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	// Parse token
	var mc = &MyClaims{}
	// Use ParseWithClaims method for custom Claim structure
	token, err := jwt.ParseWithClaims(tokenString, mc,
		// KeyFunc to return the secret key
		func(token *jwt.Token) (i interface{}, err error) {
			return mySecret, nil
		})
	if err != nil {
		return nil, err
	}
	// Type assert the Claim in token object
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid { // Validate token
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
