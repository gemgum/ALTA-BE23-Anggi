package handler

import (
	"fmt"
	"main/internal/features/todos"
	"main/internal/helper"
	"main/internal/utils"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	// "todo/internal/models"
)

type TodoController struct {
	tserv todos.Services
	jwt   utils.JwtUtilityInterface
}

func NewTodoController(t todos.Services, j utils.JwtUtilityInterface) todos.Handler {
	return &TodoController{
		tserv: t,
		jwt:   j,
	}
}

// func (tc *TodoController) AddTodo(id uint) (bool, error) {
// 	var newData models.Todo
// 	fmt.Print("Masukkan Aktivitas ")
// 	fmt.Scanln(&newData.Activity)
// 	newData.Owner = id
// 	_, err := tc.model.AddTodo(newData)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

func (tc *TodoController) AddTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input TodoRequest
		err := c.Bind(&input)

		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
		}

		var idFromToken = tc.jwt.DecodeToken(c.Get("user").(*jwt.Token))

		if input.Owner != uint(idFromToken) {
			return c.JSON(401, helper.ResponseFormat(400, "Unauthorized", nil))
		}

		_, err = tc.tserv.AddTodo(ToModelTodo(input))
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
	}
}
func (tc *TodoController) UpdateTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input TodoRequest
		err := c.Bind(&input)
		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "bad requestr", nil))
		}

		var idFromToken = tc.jwt.DecodeToken(c.Get("user").(*jwt.Token))

		if input.Owner != uint(idFromToken) {
			return c.JSON(401, helper.ResponseFormat(401, "Unauthorized", nil))
		}

		_, err = tc.tserv.UpdateTodo(ToModelTodo(input))
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success update data", nil))
	}
}
func (tc *TodoController) DeleteTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var input TodoRequest
		var idFromToken = tc.jwt.DecodeToken(c.Get("user").(*jwt.Token))
		input.Owner = uint(idFromToken)

		id := c.Param("id")
		todoID, err := strconv.Atoi(id)
		input.ID = uint(todoID)

		if err != nil {
			return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
		}
		_, err = tc.tserv.DeleteTodo(ToModelTodo(input))
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(200, helper.ResponseFormat(200, "success Delete data", nil))
	}
}

func (tc *TodoController) ShowTodo() echo.HandlerFunc {
	return func(c echo.Context) error {
		var idFromToken = tc.jwt.DecodeToken(c.Get("user").(*jwt.Token))
		result, err := tc.tserv.ShowTodo(uint(idFromToken))
		if err != nil {
			return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
		}
		return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "success get data", ToTodoResponses(result)))

	}
}

func (tc *TodoController) ShowParseData(tm []todos.Todo) {
	for _, tmData := range tm {
		fmt.Printf("ID: %d, Activity: %s, Mark: %t, Owner: %d\n",
			tmData.ID, tmData.Activity, tmData.Mark, tmData.Owner)
	}
}
