package auth

import (
	"Knoxiaes/fairesults/helpers"
	"context"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}
var allowedPaths = []string{"login", "signup","verfiy"}
var alwaysAllowed = []string{"assets","query"}
		


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fullPath := c.Request.URL.Path
		splitedPath := strings.Split(fullPath,"/")[1]

		if slices.Contains(alwaysAllowed,splitedPath){
			c.Next()
			return
		}

		tokenString, err := c.Cookie("jwt_token")

		if slices.Contains(allowedPaths,splitedPath){
			if tokenString == ""{
				c.Next()
				return
			}else{
				c.Redirect(http.StatusSeeOther,"/results")
				return
			}
		}

		// Allow unauthenticated users in
		if err != nil{
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}


		username, err := helpers.ParseToken(tokenString)
		if err != nil {
			c.Redirect(http.StatusSeeOther, "/login")
			return
		}
		// put it in context
		ctx := context.WithValue(c.Request.Context(), userCtxKey, username)

		// and call the next with our new context
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}

}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(userCtxKey).(string)
	return raw
}
