package main

import (
	"be23/configs"
	"be23/internal/routes"
	"be23/internal/utils"
	"fmt"
	"log"

	"be23/internal/features/users"
	userHandler "be23/internal/features/users/handler"
	userRepository "be23/internal/features/users/repository"
	userServices "be23/internal/features/users/services"

	"be23/internal/features/todos"
	todoHandler "be23/internal/features/todos/handler"
	todoRepository "be23/internal/features/todos/repository"
	todoServices "be23/internal/features/todos/services"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func InitUserRoute(db *gorm.DB) users.Handler {
	um := userRepository.NewUserModel(db)
	pu := utils.NewPasswordUtility()
	jwt := utils.NewJwtUtility()
	vldt := utils.NewAccountUtility(*validator.New())
	us := userServices.NewUserService(um, vldt, pu, jwt)
	uc := userHandler.NewUserController(us)
	return uc
}

func InitTodoRoute(db *gorm.DB) todos.Handler {
	tm := todoRepository.NewTodoModel(db)
	ts := todoServices.NewTodoService(tm)
	jwt := utils.NewJwtUtility()
	tc := todoHandler.NewTodoController(ts, jwt)
	return tc
}

func main() {

	setup := configs.ImportSetting()
	connection, err := configs.ConnectDB(setup)
	if err != nil {
		fmt.Println("Stop program, masalah pada database", err.Error())
		return
	}

	err = connection.AutoMigrate(&todoRepository.Todo{}, &userRepository.User{})

	if err != nil {
		log.Fatal("Stop program, masalah pada migrasi database", err.Error())
		return
	}
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORS()) // ini aja cukup

	ur := InitUserRoute(connection)
	tr := InitTodoRoute(connection)

	routes.InitRoute(e, tr, ur)
	e.Logger.Fatal(e.Start(":8000"))

}
