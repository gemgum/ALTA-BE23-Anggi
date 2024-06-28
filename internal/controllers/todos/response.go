package todos

import "main/internal/models"

type TodoResponse struct {
	ID       uint   `json:"id"`
	Activity string `json:"activity"`
	Mark     bool   `json:"mark"`
	Owner    uint   `json:"owner"`
}

func ToTodoResponse(input models.Todo) TodoResponse {
	return TodoResponse{
		ID:       input.ID,
		Activity: input.Activity,
		Mark:     input.Mark,
		Owner:    input.Owner,
	}
}

func ToTodoResponses(input []models.Todo) []TodoResponse {
	responses := make([]TodoResponse, len(input))
	for i, todo := range input {
		responses[i] = ToTodoResponse(todo)
	}
	return responses
}
