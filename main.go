package main

import (
	"fmt"
	"main/configs"
	"main/routes"

	// "main/internal/controllers/todos"
	"main/internal/controllers/todos"
	"main/internal/controllers/users"
	"main/internal/models"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}

	connection.AutoMigrate(&models.User{}, &models.Todo{})

	// var inputMenu int
	um := models.UserModel{Connection: connection}
	uc := users.UserController{Model: um}

	tu := models.TodoModel{Connection: connection}
	tc := todos.TodoController{Model: tu}

	e := echo.New()

	// e.POST("/users", uc.Register)
	// e.POST("/login", uc.Login)
	// e.POST("/todo", tc.AddTodo)
	// e.PUT("/todo", tc.UpdateTodo)
	// e.DELETE("/todo", tc.DeleteTodo)
	// e.GET("/todo/:id", tc.ShowTodo)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup
	routes.InitRoute(e, uc, tc)
	e.Logger.Fatal(e.Start(":8000"))

}
