package todos

import "main/internal/models"

type TodoRequest struct {
	ID       uint   `json:"ID"`
	Activity string `json:"activity"`
	Mark     bool   `json:"mark"`
	Owner    uint   `json:"owner"`
}

func ToModelTodo(tr TodoRequest) models.Todo {
	return models.Todo{
		ID:       tr.ID,
		Activity: tr.Activity,
		Mark:     tr.Mark,
		Owner:    tr.Owner,
	}
}
