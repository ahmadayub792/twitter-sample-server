package main

import (
	"fmt"

	"github.com/ahmadayub792/twitter-sample-server/app"
	"github.com/ahmadayub792/twitter-sample-server/handler"
	"github.com/ahmadayub792/twitter-sample-server/middleware"
	"github.com/ahmadayub792/twitter-sample-server/model"
	"github.com/ahmadayub792/twitter-sample-server/store"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := model.Setup("localhost", "5432", "postgres", "postgres")
	if err != nil {
		panic(err)
	}

	userStore := store.NewUserStore(db)
	clientStore := store.NewClientStore(db)
	targetStore := store.NewTargetStore(db)
	bcryptHasher := app.NewBcryptHasher(10)

	// Seed Data
	{
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
			{
				Email:    "ahmadayub792@gmail.com",
				Password: hashPass,
				Role:     model.RoleClient,
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

		var Rtargets []*model.Target
		Rtargets, err = targetStore.FindAll()
		if err != nil {
			panic(err)
		}
		spew.Dump(Rtargets)
	}

	// Application
	myapp := &app.App{
		UserStore:      userStore,
		ClientStore:    clientStore,
		TargetStore:    targetStore,
		PasswordHasher: bcryptHasher,
		TokenSecret:    []byte("Hello World"),
	}

	token, err := myapp.GenerateToken("admin1", "HelloWorld")
	if err != nil {
		panic(err)
	}

	fmt.Println(token)
	err = myapp.VerifyToken(token)
	if err != nil {
		panic(err)
	}
	spew.Dump(myapp.User)

	r := gin.Default()

	r.Use(middleware.SetAppSession(&app.App{
		UserStore:      userStore,
		ClientStore:    clientStore,
		TargetStore:    targetStore,
		PasswordHasher: bcryptHasher,
		TokenSecret:    []byte("Hello World"),
	}))

	r.POST("/users/authenticate", handler.Login)
	r.GET("/users", middleware.Authenticate(), handler.ListUsers) //implemented auth middleware
	r.Run()
}
