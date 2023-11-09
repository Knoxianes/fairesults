package handlers

import (
	"Knoxianes/fairecords/db/users"
	"Knoxianes/fairecords/db/users/schema"
	"Knoxianes/fairecords/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPOST(c *gin.Context){
    var user schema.User
    err := c.BindJSON(&user)
    if err != nil{
        c.JSON(http.StatusBadRequest,gin.H{"message":"error"})
        return
    }
    
    result, err := users.GetUserByUsername(user.Username)
    if err != nil{
        c.JSON(http.StatusBadRequest,gin.H{"message":"bad login"})
        return
    }

    if !helpers.CheckPasswordHash(user.Password,result.Password){
        c.JSON(http.StatusBadRequest,gin.H{"message":"bad login"})
        return
    }

    if result.Verified.Int64 == 0{
        c.JSON(http.StatusBadRequest,gin.H{"message":"not verified"})
        return
    }
    
    c.SetCookie("token",result.Token.String,3.156e+8,"/","localhost",false,true)  
    c.JSON(http.StatusOK,gin.H{"message":"done"})
}

func LoginGET(c *gin.Context){
    c.String(http.StatusOK, "Log In")    
}
