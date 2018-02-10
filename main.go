package main

import (
	"github.com/BooookStore/echovue/controller"
	"github.com/BooookStore/echovue/model"
	"github.com/labstack/echo"
)

var (
	mockdb = []*model.User{
		{ID: 1, Name: "bookstore", Age: 24},
		{ID: 2, Name: "ryosuke", Age: 25},
		{ID: 3, Name: "yuki", Age: 26},
	}
)

func main() {
	var (
		e = echo.New()
		h = controller.Handler{mockdb}
	)

	// RESTful API
	e.GET("/users", h.RetrieveUsers)
	e.GET("/user/:id", h.RetrieveUser)

	// Static file
	e.Static("/", "assets")

	// Server startup
	e.Logger.Fatal(e.Start(":8080"))
}
