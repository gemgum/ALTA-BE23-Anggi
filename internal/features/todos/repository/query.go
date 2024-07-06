package repository

import (
	"be23/internal/features/todos"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) todos.Query {
	return &TodoModel{
		db: connection,
	}
}

func (tm *TodoModel) AddTodo(newData todos.Todo) (todos.Todo, error) {
	resultData := ToTodoQuery(newData)
	err := tm.db.Create(&resultData).Error
	if err != nil {
		return todos.Todo{}, err
	}

	return newData, nil
}

func (tm *TodoModel) UpdateTodo(updateData todos.Todo) (todos.Todo, error) {
	query := `UPDATE "be23"."todos" SET "mark"= true,"updated_at"= ?
	WHERE ID = ? AND "mark"= false AND "todos"."deleted_at" IS NULL;`
	err := tm.db.Exec(query, time.Now(), &updateData.ID).Error
	fmt.Println(query)
	if err != nil {
		return todos.Todo{}, err
	}
	return updateData, nil
}

func (tm *TodoModel) DeleteTodo(deleteData todos.Todo) error {
	deleteData.Mark = true
	// deleteData.UpdatedAt = time.Now()
	query := ` UPDATE "be23"."todos" SET "deleted_at"= ?
	WHERE ID = ? AND owner = ? AND "todos"."deleted_at" IS NULL;`
	err := tm.db.Debug().Exec(query, time.Now(), &deleteData.ID, &deleteData.Owner).Error
	fmt.Println(query)
	if err != nil {
		return err
	}

	return nil
}

func (tm *TodoModel) ShowTodo(id uint) ([]todos.Todo, error) {
	todos := make([]todos.Todo, 0)
	query := `SELECT * FROM "be23"."todos"
	WHERE owner = ?  AND "todos"."deleted_at" IS NULL;`
	err := tm.db.Debug().Raw(query, id).Scan(&todos).Error
	fmt.Println(query)
	if err != nil {
		return todos, err

	}
	return todos, nil
}
