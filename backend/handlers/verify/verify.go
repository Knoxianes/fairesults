package verifyHandler

import (
	"Knoxiaes/fairesults/database"
	"Knoxiaes/fairesults/helpers"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GET(c *gin.Context) {
	tokenString := c.Param("token")
	username, err := helpers.ParseToken(tokenString)
	if err != nil {
		log.Println(err)
		c.String(http.StatusUnauthorized, "Token expired or invalid, try registering again")
		return
	}
	row := database.DB.QueryRow("select verified from users where username = ?;", username)

	var verified int
	err = row.Scan(&verified)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
			c.String(http.StatusBadRequest, "User not registered, try registering again")
			return
		} else {
			log.Println(err)
			c.String(http.StatusInternalServerError, fmt.Sprintf("Error %s", err))
			return
		}
	}

	if verified == 1 {
		c.String(http.StatusAlreadyReported, "Email already verified")
		return
	}

	_, err = database.DB.Query("update users set verified = 1 where username = ?;", username)
	if err != nil {
		log.Println(err)
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error %s", err))
		return
	}

	c.String(http.StatusOK, "Verification successful")

}
