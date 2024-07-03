package handler

import "main/internal/features/todos"

type TodoResponse struct {
	ID       uint   `json:"id"`
	Activity string `json:"activity"`
	Mark     bool   `json:"mark"`
	Owner    uint   `json:"owner"`
}

func ToTodoResponse(tInp todos.Todo) TodoResponse {
	return TodoResponse{
		ID:       tInp.ID,
		Activity: tInp.Activity,
		Mark:     tInp.Mark,
		Owner:    tInp.Owner,
	}
}

func ToTodoResponses(tInp []todos.Todo) []TodoResponse {
	responses := make([]TodoResponse, len(tInp))
	for i, todo := range tInp {
		responses[i] = ToTodoResponse(todo)
	}
	return responses
}
