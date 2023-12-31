package helpers

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)

func GenerateToken(username string, duration time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(duration).Unix()
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT")))
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	return tokenString, nil
}

func ParseToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT")), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}

func ParseTokenFromContext(c *gin.Context) (string, error) {
	tokenString, err := c.Cookie("jwt_token")
	if err != nil || tokenString == "" {
		return "", CustomError{Message: err.Error(), Code: 0}
	}
	username, err := ParseToken(tokenString)
	if err != nil {
		log.Println(err)
		return "", CustomError{Message: err.Error(), Code: 6}
	}
	return username, nil
}
