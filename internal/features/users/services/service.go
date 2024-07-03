package services

import (
	"main/internal/features/users"
	"main/internal/utils"
)

type UserServices struct {
	qry users.Query
}

func NewUserService(q users.Query) users.Services {
	return &UserServices{
		qry: q,
	}
}

func (us *UserServices) Register(newData users.Users) error {
	processPw, err := utils.GeneratePassword(newData.Password)

	if err != nil {
		return err
	}

	newData.Password = string(processPw)

	err = us.qry.Register(newData)

	if err != nil {
		return err
	}

	return nil
}

func (us *UserServices) Login(email string, password string) (users.Users, string, error) {
	result, err := us.qry.Login(email)
	if err != nil {
		return users.Users{}, "", err
	}

	err = utils.CheckPassword([]byte(password), []byte(result.Password))

	if err != nil {
		return users.Users{}, "", err
	}

	token, err := utils.GenerateJWT(result.ID, result.Email)
	if err != nil {
		return users.Users{}, "", err
	}

	return result, token, nil
}
