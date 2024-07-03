package repository

import (
	"main/internal/features/todos"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	ID       uint
	Activity string
	Mark     bool
	Owner    uint
}

func ToTodoQuery(todofromInputData todos.Todo) Todo {
	return Todo{
		ID:       todofromInputData.ID,
		Activity: todofromInputData.Activity,
		Mark:     todofromInputData.Mark,
		Owner:    todofromInputData.Owner,
	}
}
