package main

import (
	"fmt"
	"main/configs"
	"main/internal/controllers/todos"
	"main/internal/controllers/users"
	"main/internal/models"

	"github.com/labstack/echo/v4"
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
	um := models.NewUserModel(connection)
	uc := users.NewUserController(um)

	tu := models.NewTodoModel(connection)
	tc := todos.NewTodoController(tu)

	e := echo.New()

	e.POST("/users", uc.Register)
	e.POST("/login", uc.Login)
	e.POST("/todo", tc.AddTodo)
	e.PUT("/todo", tc.UpdateTodo)
	e.DELETE("/todo", tc.DeleteTodo)
	e.GET("/todo/:id", tc.ShowTodo)

	e.Start(":8000")

}
