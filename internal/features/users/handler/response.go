package handler

import user "main/internal/features/users"

type LoginResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"hp"`
}

func ToLoginReponse(input user.Users) LoginResponse {
	return LoginResponse{
		ID:    input.ID,
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}
}
