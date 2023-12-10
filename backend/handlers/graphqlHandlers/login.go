package graphqlHandlers

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/graph/model"
	"Knoxiaes/fairesults/helpers"
	"database/sql"
	"log"
	"time"
	_"github.com/joho/godotenv/autoload"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context, input model.Login) (string, error) {
	res := database.DB.QueryRow("select password,verified from users where username=?;", input.Username)
	var password string
	var verified int
	err := res.Scan(&password, &verified)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", helpers.CustomError{Message: "User doesnt exists", Code: 3}
		} else {
			log.Println(err)
			return "", helpers.CustomError{Message: err.Error(), Code: 0}
		}
	}
	if verified == 0 {
		return "", helpers.CustomError{Message: "User not verified", Code: 4}
	}

	if !helpers.CheckPasswordHash(input.Password, password) {
		return "", helpers.CustomError{Message: "Wrong username or password", Code: 5}
	}
	token, err := helpers.GenerateToken(input.Username,time.Hour * 24)
	if err != nil{
		log.Println(err)
		return "", helpers.CustomError{Message:err.Error(), Code:0}
	}
	c.SetCookie("jwt_token", token, 86400, "/","",false,false) // Remainder in production last two have to be true
	return "", nil
}
