package todos

import (
	"fmt"
	"main/helper"
	"main/internal/models"
	"main/middlewares"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	// "todo/internal/models"
)

type TodoController struct {
	Model models.TodoModel
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

func (tc *TodoController) AddTodo(c echo.Context) error {
	var input TodoRequest
	err := c.Bind(&input)

	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
	}

	var idFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))

	if input.Owner != uint(idFromToken) {
		return c.JSON(401, helper.ResponseFormat(400, "Unauthorized", nil))
	}

	_, err = tc.Model.AddTodo(ToModelTodo(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
}

func (tc *TodoController) UpdateTodo(c echo.Context) error {
	var input TodoRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "bad requestr", nil))
	}

	var idFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))

	if input.Owner != uint(idFromToken) {
		return c.JSON(401, helper.ResponseFormat(401, "Unauthorized", nil))
	}

	_, err = tc.Model.UpdateTodo(ToModelTodo(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(200, helper.ResponseFormat(200, "success update data", nil))
}

func (tc *TodoController) DeleteTodo(c echo.Context) error {
	var input TodoRequest
	var idFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))
	input.Owner = uint(idFromToken)

	id := c.Param("id")
	todoID, err := strconv.Atoi(id)
	input.ID = uint(todoID)

	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
	}
	_, err = tc.Model.DeleteTodo(ToModelTodo(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(200, helper.ResponseFormat(200, "success Delete data", nil))
}

func (tc *TodoController) ShowTodo(c echo.Context) error {
	var idFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))
	result, err := tc.Model.ShowTodo(uint(idFromToken))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "success get data", ToTodoResponses(result)))

}

// func (tc *TodoController) ShowTodo() echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		var idFromToken = middlewares.DecodeToken(c.Get("user").(*jwt.Token))

// 		// id := c.Param("id")
// 		// 	id := idFromToken
// 		result, err := tc.Model.ShowTodo(uint(idFromToken))

// 		if err != nil {
// 			if strings.Contains(err.Error(), "not found") {
// 				return c.JSON(http.StatusNotFound, map[string]any{
// 					"code":    http.StatusNotFound,
// 					"message": "data tidak ditemukan",
// 				})
// 			}
// 			return c.JSON(http.StatusInternalServerError, map[string]any{
// 				"code":    http.StatusInternalServerError,
// 				"message": "terjadi kesalahan pada sistem",
// 			})
// 		}
// 		return c.JSON(http.StatusOK, map[string]any{
// 			"code":    http.StatusOK,
// 			"message": "berhasil mendapatkan data",
// 			"data":    result,
// 		})
// 	}
// }

func (tc *TodoController) ShowParseData(tm []models.Todo) {
	for _, tmData := range tm {
		fmt.Printf("ID: %d, Activity: %s, Mark: %t, Owner: %d, CreatedAt: %s, UpdatedAt: %s\n",
			tmData.ID, tmData.Activity, tmData.Mark, tmData.Owner, tmData.CreatedAt, tmData.UpdatedAt)
	}
}
