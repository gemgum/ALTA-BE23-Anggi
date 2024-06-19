package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Activity string
	Mark     bool
	Owner    uint
}

type TodoModel struct {
	db *gorm.DB
}

func NewTodoModel(connection *gorm.DB) *TodoModel {
	return &TodoModel{
		db: connection,
	}
}

func (tm *TodoModel) AddTodo(newData Todo) (Todo, error) {
	newData.Mark = false
	err := tm.db.Create(&newData).Error
	if err != nil {
		return Todo{}, err
	}

	return newData, nil
}

func (tm *TodoModel) UpdateTodo(updateData Todo) (Todo, error) {
	updateData.Mark = true
	updateData.UpdatedAt = time.Now()

	// err := tm.db.Debug().Model(&Todo{}).Where("owner = ? AND activity = ?", &updateData.Owner, &updateData.Activity).Updates(map[string]interface{}{
	// 	"mark":     updateData.Mark,
	// }).Error

	query := `UPDATE "be23"."todos" SET "mark"=?,"updated_at"= ?
	WHERE "owner" = ? AND "activity"= ? AND "mark"= false AND "todos"."deleted_at" IS NULL;`
	err := tm.db.Exec(query, &updateData.Mark, &updateData.UpdatedAt, &updateData.Owner, &updateData.Activity).Error
	fmt.Println(query)
	if err != nil {
		return Todo{}, err
	}

	return updateData, nil
}

func (tm *TodoModel) DeleteTodo(deleteData Todo) (Todo, error) {
	deleteData.Mark = true
	deleteData.UpdatedAt = time.Now()

	// err := tm.db.Debug().Where("owner = ? AND activity = ?", deleteData.Owner, deleteData.Activity).Delete(&Todo{}).Error

	query := ` UPDATE "be23"."todos" SET "deleted_at"= ? 
	WHERE (owner = ? AND activity = ?) AND "todos"."deleted_at" IS NULL `
	err := tm.db.Exec(query, &deleteData.UpdatedAt, &deleteData.Owner, &deleteData.Activity).Error
	fmt.Println(query)
	if err != nil {
		return Todo{}, err
	}

	return deleteData, nil
}

func (tm *TodoModel) ShowTodo(showTodoData Todo) ([]Todo, error) {
	todos := make([]Todo, 0)
	query := `SELECT * FROM "be23"."items"
	WHERE owner = ?  AND "items"."deleted_at" IS NULL;`
	err := tm.db.Debug().Raw(query, &showTodoData.Owner).Scan(&todos).Error
	fmt.Println(query)
	if err != nil {
		// return Todo{}, err
		return nil, err

	}
	return todos, nil
}
