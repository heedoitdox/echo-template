package utils

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JwtCustomClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

var JWTConfig = middleware.JWTConfig{
	ContextKey:  "jwt",
	TokenLookup: "header:" + echo.HeaderAuthorization,
	AuthScheme:  "Bearer",
	Claims:      &JwtCustomClaims{},
	SigningKey:  []byte(os.Getenv("TOKEN_SECRET")),
}

func CreateJWT(userID int, cTime time.Duration) (string, error) {
	s := "[utils/CreateJWT]"

	tk := &JwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * cTime).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tk)
	log.Printf("%v\n", os.Getenv("TOKEN_SECRET"))
	jwt, err := token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		log.Printf("%s %v\n", s, err)
		return "", err
	}

	return jwt, nil
}
