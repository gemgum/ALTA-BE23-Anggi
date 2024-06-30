package routes

import (
	// "main/configs"
	"main/configs"
	"main/internal/controllers/todos"
	"main/internal/controllers/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, ctlUser users.UserController, ctlTodo todos.TodoController) {
	c.POST("/users", ctlUser.Register) // register -> umum (boleh diakses semua orang)
	c.POST("/login", ctlUser.Login)
	c.GET("/todo", ctlTodo.ShowTodo, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.ImportSetting().JWTSECRET),
	}))
	c.DELETE("/todo/:id", ctlTodo.DeleteTodo, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.ImportSetting().JWTSECRET),
	}))

	c.PUT("/todo", ctlTodo.UpdateTodo, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.ImportSetting().JWTSECRET),
	}))

	c.POST("/todo", ctlTodo.AddTodo, echojwt.WithConfig(echojwt.Config{
		SigningKey: []byte(configs.ImportSetting().JWTSECRET),
	}))
}
