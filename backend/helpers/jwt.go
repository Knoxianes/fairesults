package helpers

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv/autoload"
)



func GenerateJWT(expireDate time.Time, username string)(string,error){
    token := jwt.New(jwt.SigningMethodHS256)
    claims := token.Claims.(jwt.MapClaims)
    tmp := time.Time{}
    if expireDate != tmp {
        claims["exp"] = expireDate.Unix()
    }
    claims["user"] = username
    return token.SignedString([]byte(os.Getenv("JWT")))
}

func VerifyJWT(tokenStr string)(*jwt.Token, error){
    return jwt.Parse(tokenStr,func(token *jwt.Token)(interface{},error){
        return []byte(os.Getenv("JWT")),nil
    })
}
