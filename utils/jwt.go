package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(userid uint) string {
	claims := jwt.MapClaims{
		"exp" : time.Now().Add(time.Hour * 3).Unix(),
		"iat" : time.Now().Unix(),
		"userID" : userid,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	return t
}

func ValidateToken(token string) (*jwt.Token, error) {
	//2nd arg function return secret key after checking if the signing method is HMAC and returned key is used by 'Parse' to decode the token)
	return jwt.Parse(token, func(token *jwt.Token) (any, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			//nil secret key
			return nil, fmt.Errorf("System Error Signing Method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
}
