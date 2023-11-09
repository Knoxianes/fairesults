package main

import (
	"Knoxianes/fairecords/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router.LoadHTMLGlob("./pages/*.html")
	// router.Static("/assets", "./assets")
	router.GET("/signup", handlers.SignupGET)
	router.POST("/signup", handlers.SignupPOST)
	router.GET("/login", handlers.LoginGET)
	router.POST("/login", handlers.LoginPOST)
	router.GET("/verify/:token", handlers.VerifyGET)
	router.Run(":4000")
}
