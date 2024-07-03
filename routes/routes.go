package routes

import (
	"main/configs"
	"main/internal/features/todos"
	"main/internal/features/users"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoute(c *echo.Echo, tr todos.Handler, ur users.Handler) {

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
