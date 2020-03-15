package main

import (
	"fmt"

	"github.com/ahmadayub792/twitter-sample-server/app"
	"github.com/ahmadayub792/twitter-sample-server/config"
	"github.com/ahmadayub792/twitter-sample-server/handler"
	"github.com/ahmadayub792/twitter-sample-server/model"
	"github.com/ahmadayub792/twitter-sample-server/store"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func main() {
	err := config.ConnectDB()
	if err != nil {
		panic(err)
	}

	if err := model.Setup(); err != nil {
		panic(err)
	}

	userStore := store.NewUserStore(config.DB)
	clientStore := store.NewClientStore(config.DB)
	targetStore := store.NewTargetStore(config.DB)
	bcryptHasher := app.NewBcryptHasher(10)

	// Seed Data

	hashPass, err := bcryptHasher.GenerateHash("HelloWorld")
	if err != nil {
		panic(err)
	}

	users := []*model.User{
		{
			Email:    "admin1",
			Password: hashPass,
			Role:     model.RoleAdmin,
		},
	}

	for _, u := range users {
		if err := userStore.Create(u); err != nil {
			panic(err)
		}
	}

	targets := []*model.Target{
		{
			Type:    "api",
			Handler: "Barack Obama",
		},
		{
			Type:    "scrapper",
			Handler: "Donald Trump",
		},
	}

	for _, t := range targets {
		if err := targetStore.Create(t); err != nil {
			panic(err)
		}
	}

	// Application
	app := &app.App{
		UserStore:      userStore,
		ClientStore:    clientStore,
		TargetStore:    targetStore,
		PasswordHasher: bcryptHasher,
		TokenSecret:    []byte("Hello World"),
	}

	token, err := app.GenerateToken("admin1", "HelloWorld")
	if err != nil {
		panic(err)
	}

	fmt.Println(token)
	err = app.VerifyToken(token)
	if err != nil {
		panic(err)
	}
	spew.Dump(app.User)

	handle := handler.NewHandler(app)
	r := gin.Default()
	r.POST("/users/authenticate", handle.Login)
	r.Run()
}
