package todos

import (
	"fmt"
	"main/helper"
	"main/internal/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	// "todo/internal/models"
)

type TodoController struct {
	model *models.TodoModel
}

func NewTodoController(m *models.TodoModel) *TodoController {
	return &TodoController{
		model: m,
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

func (tc *TodoController) AddTodo(c echo.Context) error {
	var input TodoRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
	}
	_, err = tc.model.AddTodo(ToModelTodo(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
}

// func (tc *TodoController) UpdateTodo(id uint) (bool, error) {
// 	var updateData models.Todo
// 	fmt.Print("Udpdate Aktivitas ")
// 	fmt.Scanln(&updateData.Activity)
// 	updateData.Owner = id
// 	_, err := tc.model.UpdateTodo(updateData)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

func (tc *TodoController) UpdateTodo(c echo.Context) error {
	var input TodoRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "bad requestr", nil))
	}
	_, err = tc.model.UpdateTodo(ToModelTodo(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(200, helper.ResponseFormat(200, "success update data", nil))
}

// func (tc *TodoController) DeleteTodo(id uint) (bool, error) {
// 	var deleteData models.Todo
// 	fmt.Print("Hapus Aktivitas ")
// 	fmt.Scanln(&deleteData.Activity)
// 	deleteData.Owner = id
// 	_, err := tc.model.DeleteTodo(deleteData)
// 	if err != nil {
// 		return false, err
// 	}

// 	return true, nil
// }

func (tc *TodoController) DeleteTodo(c echo.Context) error {
	var input TodoRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
	}
	_, err = tc.model.DeleteTodo(ToModelTodo(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(200, helper.ResponseFormat(200, "success Delete data", nil))
}

// func (tc *TodoController) ShowTodos(id uint) ([]models.Todo, error) {
// 	var showTodoData models.Todo
// 	fmt.Print("Menampilkan Aktivitas ")
// 	// fmt.Scanln(&showTodoData.Activity)
// 	showTodoData.Owner = id
// 	todo, err := tc.model.ShowTodo(showTodoData)
// 	if err != nil {
// 		return todo, err
// 	}
// 	// fmt.Println("type", reflect.TypeOf(showTodoData))
// 	return todo, nil
// }

func (tc *TodoController) ShowTodo(c echo.Context) error {
	id := c.Param("id")
	todoID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, helper.ResponseFormat(http.StatusBadRequest, "invalid ID", nil))
	}
	result, err := tc.model.ShowTodo(uint(todoID))

	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(http.StatusOK, helper.ResponseFormat(http.StatusOK, "success get data", ToTodoResponses(result)))

}

func (tc *TodoController) ShowParseData(tm []models.Todo) {
	for _, tmData := range tm {
		fmt.Printf("ID: %d, Activity: %s, Mark: %t, Owner: %d, CreatedAt: %s, UpdatedAt: %s\n",
			tmData.ID, tmData.Activity, tmData.Mark, tmData.Owner, tmData.CreatedAt, tmData.UpdatedAt)
	}
}
