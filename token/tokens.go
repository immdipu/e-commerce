package tokens

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	email string
	jwt.RegisteredClaims
}

func GenerateToken(email string) (string, error) {
	var SECRET_KEY = os.Getenv("SECRET_KEY")
	log.Println("email: ", email)
	claims := &Claims{
		email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(SECRET_KEY))

	return tokenString, err
}
