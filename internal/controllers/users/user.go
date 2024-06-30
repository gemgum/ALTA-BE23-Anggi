package users

import (
	"main/helper"
	"main/internal/models"
	"main/middlewares"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	Model models.UserModel
}

// func NewUserController(m *models.UserModel) *UserController {
// 	return &UserController{
// 		model: m,
// 	}
// }

func (uc *UserController) Register(c echo.Context) error {
	var input RegisterRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
	}
	_, err = uc.Model.Register(ToModelUsers(input))
	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}
	return c.JSON(201, helper.ResponseFormat(201, "success insert data", nil))
}

func (uc *UserController) Login(c echo.Context) error {
	var input LoginRequest
	err := c.Bind(&input)
	if err != nil {
		return c.JSON(400, helper.ResponseFormat(400, "bad request", nil))
	}
	result, err := uc.Model.Login(input.Email, input.Password)

	if err != nil {
		return c.JSON(500, helper.ResponseFormat(500, "server error", nil))
	}

	token, err := middlewares.GenerateJWT(result.ID, result.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]any{
			"code":    http.StatusInternalServerError,
			"message": "terjadi kesalahan pada sistem, gagal memproses data",
		})
	}

	// return c.JSON(200, helper.ResponseFormat(200, "success login", ToLoginReponse(result)))
	return c.JSON(http.StatusOK, map[string]any{"code": http.StatusOK, "message": "selama anda berhasil login", "data": ToLoginReponse(result), "token": token})

}
