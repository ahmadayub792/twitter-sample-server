package middleware

import (
	"fmt"
	"net/http"

	"github.com/ahmadayub792/twitter-sample-server/app"
	"github.com/gin-gonic/gin"
)

func SetAppSession(session *app.App) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("app", session)
		c.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		myapp := c.MustGet("app").(*app.App)
		tokenStr := c.Request.Header.Get("token")
		if err := myapp.VerifyToken(tokenStr); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("%v", err)})
		}
		c.Next()
	}
}
