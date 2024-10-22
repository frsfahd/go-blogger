package server

import (
	"errors"
	"os"
	"time"

	_ "github.com/dotenv-org/godotenvvault/autoload"
	"github.com/frsfahd/go-blogger/internal/sqlc"
	"github.com/golang-jwt/jwt/v5"
)

var (
	SECRET = []byte(os.Getenv("SECRET"))
)

type LoginClaims struct {
	LoginData
	jwt.RegisteredClaims
}

type LoginData struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

func signToken(user sqlc.User) string {

	// Create claims with multiple fields populated
	claims := LoginClaims{
		LoginData{
			Email: user.Email,
			Role:  user.Role,
		},
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, _ := token.SignedString(SECRET)

	return ss

}

func parseToken(tokenString string) (LoginData, error) {
	token, err := jwt.ParseWithClaims(tokenString, &LoginClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SECRET), nil
	})
	if err != nil {
		return LoginData{}, err
	} else if claims, ok := token.Claims.(*LoginClaims); ok {
		return LoginData{Email: claims.Email, Role: claims.Role}, nil
	} else {
		return LoginData{}, errors.New("unknown claims type, cannot proceed")
	}
}
