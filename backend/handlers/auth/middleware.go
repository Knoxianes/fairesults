package auth

import (
	"Knoxiaes/fairesults/helpers"
	"context"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}
var allowedPaths = ["login", "signup","verfiy","query"]
		


func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		fullPath := c.Request.URL.Path
		splitedPath := strings.Split(fullPath,"/")[1]
		if  

		tokenString, err := c.Cookie("jwt_token")

		// Allow unauthenticated users in
		if err != nil || tokenString == "" {
			c.Redirect(http.StatusForbidden, "/login")
			return
		}

		username, err := helpers.ParseToken(tokenString)
		if err != nil {
			c.Redirect(http.StatusForbidden, "/login")
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
