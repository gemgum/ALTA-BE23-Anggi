package routes

import (
	// "main/configs"

	"main/configs"
	"main/internal/features/users"
	userHandler "main/internal/features/users/handler"
	userRepository "main/internal/features/users/repository"
	userServices "main/internal/features/users/services"

	"main/internal/features/todos"
	todoHandler "main/internal/features/todos/handler"
	todoRepository "main/internal/features/todos/repository"
	todoServices "main/internal/features/todos/services"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
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

// func InitUserRoute(db *gorm.DB) users.Handler {
// 	tm := repository.NewUserModel(db)
// 	ts := services.NewUserService(tm)
// 	tc := handler.NewUserController(ts)
// 	return tc
// }

func InitRoute(c *echo.Echo, db *gorm.DB) {

	ur := InitUserRoute(db)
	tr := InitTodoRoute(db)
	// tm := repository.NewTodoModel(db)
	// ts := services.NewTodoService(tm)
	// tc := handler.NewTodoController(ts)

	c.POST("/users", ur.Register()) // register -> umum (boleh diakses semua orang)
	c.POST("/login", ur.Login())

	c.GET("/todo", tr.ShowTodo(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.ImportSetting().JWTSECRET),
	}))
	c.DELETE("/todo/:id", tr.DeleteTodo(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.ImportSetting().JWTSECRET),
	}))

	c.PUT("/todo", tr.UpdateTodo(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.ImportSetting().JWTSECRET),
	}))

	c.POST("/todo", tr.AddTodo(), echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.ImportSetting().JWTSECRET),
	}))
}
