package services

import (
	"be23/internal/features/todos"
)

type TodoServices struct {
	qry todos.Query
}

func NewTodoService(q todos.Query) todos.Services {
	return &TodoServices{
		qry: q,
	}
}

func (ts *TodoServices) AddTodo(newData todos.Todo) (todos.Todo, error) {
	newData, err := ts.qry.AddTodo(newData)
	if err != nil {
		return todos.Todo{}, err
	}

	return newData, nil
}

func (ts *TodoServices) UpdateTodo(updateData todos.Todo) (todos.Todo, error) {
	updateData.Mark = true
	// err := ts.qry.Exec(query, &updateData.UpdatedAt, &updateData.ID).Error
	updateData, err := ts.qry.UpdateTodo(updateData)

	if err != nil {
		return todos.Todo{}, err
	}

	return updateData, nil
}

func (ts *TodoServices) DeleteTodo(deleteData todos.Todo) (todos.Todo, error) {
	err := ts.qry.DeleteTodo(deleteData)

	if err != nil {
		return todos.Todo{}, err
	}

	return deleteData, nil
}

func (ts *TodoServices) ShowTodo(id uint) ([]todos.Todo, error) {
	todos := make([]todos.Todo, 0)
	// var todos Todo

	// query := `SELECT * FROM "be23"."todos"
	// WHERE owner = ?  AND "todos"."deleted_at" IS NULL;`
	// err := ts.qry.Debug().Raw(query, id).Scan(&todos).Error
	// fmt.Println(query)
	todos, err := ts.qry.ShowTodo(id)

	if err != nil {
		// return Todo{}, err
		return todos, err

	}
	return todos, nil
}
