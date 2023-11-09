package handlers

import (
	"Knoxianes/fairecords/db/users"
	"Knoxianes/fairecords/helpers"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyGET(c *gin.Context) {
	tokenStr := c.Param("token")

	token, err := helpers.VerifyJWT(tokenStr)
    if err != nil{
        c.String(http.StatusBadRequest,fmt.Sprintf("Error while verifying a token, token probably expired. %s",err))
        return
    }

    claims  := token.Claims.(jwt.MapClaims)
    username := claims["user"]

    err = users.UpdateUserVerified(fmt.Sprintf("%v",username),true)
    if err != nil {
        c.String(http.StatusBadRequest,"Error while updating database try again later")
        return
    }

    c.String(http.StatusOK,"Verification successful.") 
}
