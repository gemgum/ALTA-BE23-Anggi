package handler

import "main/internal/features/todos"

type TodoRequest struct {
	ID       uint   `json:"id"`
	Activity string `json:"activity"`
	Mark     bool   `json:"mark"`
	Owner    uint   `json:"owner"`
}

func ToModelTodo(tr TodoRequest) todos.Todo {
	return todos.Todo{
		ID:       tr.ID,
		Activity: tr.Activity,
		Mark:     tr.Mark,
		Owner:    tr.Owner,
	}
}
