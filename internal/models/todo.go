package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID       uint
	Activity string
	Mark     bool
	Owner    uint
}

type TodoModel struct {
	Connection *gorm.DB
}

func (tm *TodoModel) AddTodo(newData Todo) (Todo, error) {
	newData.Mark = false
	err := tm.Connection.Create(&newData).Error
	if err != nil {
		return Todo{}, err
	}

	return newData, nil
}

func (tm *TodoModel) UpdateTodo(updateData Todo) (Todo, error) {
	updateData.Mark = true
	updateData.UpdatedAt = time.Now()
	// query := `UPDATE "be23"."todos" SET "mark"=?,"updated_at"= ?
	// WHERE "owner" = ? AND "activity"= ? AND "mark"= false AND "todos"."deleted_at" IS NULL;`
	query := `UPDATE "be23"."todos" SET "mark"= true,"updated_at"= ?
	WHERE ID = ? AND "mark"= false AND "todos"."deleted_at" IS NULL;`
	err := tm.Connection.Exec(query, &updateData.UpdatedAt, &updateData.ID).Error
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
	WHERE ID = ? AND owner = ? AND "todos"."deleted_at" IS NULL;`
	err := tm.Connection.Debug().Exec(query, &deleteData.UpdatedAt, &deleteData.ID, &deleteData.Owner).Error
	fmt.Println(query)
	if err != nil {
		return Todo{}, err
	}

	return deleteData, nil
}

func (tm *TodoModel) ShowTodo(id uint) ([]Todo, error) {
	todos := make([]Todo, 0)
	// var todos Todo

	query := `SELECT * FROM "be23"."todos"
	WHERE owner = ?  AND "todos"."deleted_at" IS NULL;`
	err := tm.Connection.Debug().Raw(query, id).Scan(&todos).Error
	fmt.Println(query)
	if err != nil {
		// return Todo{}, err
		return todos, err

	}
	return todos, nil
}
