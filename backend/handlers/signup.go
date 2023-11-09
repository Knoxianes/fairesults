package handlers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"Knoxianes/fairecords/db/users"
	"Knoxianes/fairecords/db/users/schema"
	"Knoxianes/fairecords/helpers"

	"github.com/gin-gonic/gin"
)

func SignupPOST(c *gin.Context) {
	var user schema.User

	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	_, err = users.GetUserByUsername(user.Username)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad username"})
		log.Println(err)
		return
	}
	_, err = users.GetUserByEmail(user.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad email"})
		return
	}

	user.Password, err = helpers.HashPassword(user.Password)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message":"error"})
		return
	}

	token, err := helpers.GenerateJWT(time.Time{}, user.Username)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message":"error"})
		return
	}
	user.Token = sql.NullString{String: token}

	token, err = helpers.GenerateJWT(time.Now().Add(time.Hour), user.Username)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message":"error"})
		return
	}
	user.Verification_token = sql.NullString{String: token}

	err = users.InsertUser(user)
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message":"error"})
		return
	}

	mail := fmt.Sprintf(`<h1>Email verification</h1><p>Please <a href="http://localhost:4000/verify/%s">click here</a> to verify email.`, user.Verification_token.String)
	if err := helpers.SendMail(mail, "Email verification", user.Email); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"message":"error"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "done"})
}
func SignupGET(c *gin.Context) {
	c.String(http.StatusOK, "Oke")
}
