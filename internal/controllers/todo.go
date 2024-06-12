package controllers

import (
	"fmt"
	"main/internal/models"
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

func (tc *TodoController) AddTodo(id uint) (bool, error) {
	var newData models.Todo
	fmt.Print("Masukkan Aktivitas ")
	fmt.Scanln(&newData.Activity)
	newData.Owner = id
	_, err := tc.model.AddTodo(newData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (tc *TodoController) UpdateTodo(id uint) (bool, error) {
	var updateData models.Todo
	fmt.Print("Udpdate Aktivitas ")
	fmt.Scanln(&updateData.Activity)
	updateData.Owner = id
	_, err := tc.model.UpdateTodo(updateData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (tc *TodoController) DeleteTodo(id uint) (bool, error) {
	var deleteData models.Todo
	fmt.Print("Hapus Aktivitas ")
	fmt.Scanln(&deleteData.Activity)
	deleteData.Owner = id
	_, err := tc.model.DeleteTodo(deleteData)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (tc *TodoController) ShowTodo(id uint) ([]models.Todo, error) {
	var showTodoData models.Todo
	fmt.Print("Menampilkan Aktivitas ")
	// fmt.Scanln(&showTodoData.Activity)
	showTodoData.Owner = id
	todo, err := tc.model.ShowTodo(showTodoData)
	if err != nil {
		return todo, err
	}
	// fmt.Println("type", reflect.TypeOf(showTodoData))
	return todo, nil
}

func (tc *TodoController) ShowParseData(tm []models.Todo) {
	for _, tmData := range tm {
		fmt.Printf("ID: %d, Activity: %s, Mark: %t, Owner: %d, CreatedAt: %s, UpdatedAt: %s\n",
			tmData.ID, tmData.Activity, tmData.Mark, tmData.Owner, tmData.CreatedAt, tmData.UpdatedAt)
	}
}
