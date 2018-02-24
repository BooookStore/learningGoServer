package main

import (
	"github.com/BooookStore/learningGoServer/controller"
	"github.com/BooookStore/learningGoServer/model"
	"github.com/labstack/echo"
)

var (
	mockdb = model.UserList{
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
	e.GET("/user", h.GetAllUser)
	e.GET("/user/:id", h.GetUser)
	e.DELETE("/user/:id", h.DeleteUser)
	e.POST("/user/", h.CreateUser)
	e.PUT("/user/:id", h.UpdateUser)

	// Static file
	e.Static("/", "assets")

	// Server startup
	e.Logger.Fatal(e.Start(":8080"))
}
