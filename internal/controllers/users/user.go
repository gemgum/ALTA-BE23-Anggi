package users

import (
	"main/helper"
	"main/internal/models"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	model *models.UserModel
}

func NewUserController(m *models.UserModel) *UserController {
	return &UserController{
		model: m,
	}
}

func (uc *UserController) Register(c echo.Context) error {
	var input RegisterRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	_, err = uc.model.Register(ToModelUsers(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
}

func (uc *UserController) Login(c echo.Context) error {
	var input LoginRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "input error", nil))
	}
	result, err := uc.model.Login(input.Email, input.Password)

	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}

	return c.JSON(200, helper.ResponseFormat(200, "success login", ToLoginReponse(result)))
}
