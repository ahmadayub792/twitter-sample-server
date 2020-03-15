package handler

import (
	"fmt"
	"net/http"

	"github.com/ahmadayub792/twitter-sample-server/app"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Login(*gin.Context)
}

type Handler struct {
	App *app.App
}

type loginCreds struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewHandler(app *app.App) Controller {
	return &Handler{app}
}

func (h *Handler) Login(c *gin.Context) {
	var creds loginCreds
	c.BindJSON(&creds)
	tokenStr, err := h.App.GenerateToken(creds.Email, creds.Password)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": fmt.Sprintf("%v", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": tokenStr})
}
