package main

import (
	"fmt"
	"main/configs"
	"main/routes"

	"main/internal/features/users"
	userHandler "main/internal/features/users/handler"
	userRepository "main/internal/features/users/repository"
	userServices "main/internal/features/users/services"

	"main/internal/features/todos"
	todoHandler "main/internal/features/todos/handler"
	todoRepository "main/internal/features/todos/repository"
	todoServices "main/internal/features/todos/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitUserRoute(db *gorm.DB) users.Handler {
	um := userRepository.NewUserModel(db)
	us := userServices.NewUserService(um)
	uc := userHandler.NewUserController(us)
	return uc
}

func InitTodoRoute(db *gorm.DB) todos.Handler {
	tm := todoRepository.NewTodoModel(db)
	ts := todoServices.NewTodoService(tm)
	tc := todoHandler.NewTodoController(ts)
	return tc
}

func main() {

	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}

	// connection.AutoMigrate(&models.Todo{})

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup

	ur := InitUserRoute(connection)
	tr := InitTodoRoute(connection)

	routes.InitRoute(e, tr, ur)
	e.Logger.Fatal(e.Start(":8000"))

}
