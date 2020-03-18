package handler

import (
	"fmt"
	"net/http"

	"github.com/ahmadayub792/twitter-sample-server/app"
	"github.com/ahmadayub792/twitter-sample-server/model"
	"github.com/gin-gonic/gin"
)

// type loginCreds struct {
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }

func Login(c *gin.Context) {
	myapp := c.MustGet("app").(*app.App)

	creds := struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}{}

	if err := c.BindJSON(&creds); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("email and password required")})
		return
	}
	tokenStr, err := myapp.GenerateToken(creds.Email, creds.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func ListUsers(c *gin.Context) {
	myapp := c.MustGet("app").(*app.App)
	if myapp.User.Role != model.RoleAdmin {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Permission denied"})
		return
	}

	users, err := myapp.UserStore.FindAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, users)
}
