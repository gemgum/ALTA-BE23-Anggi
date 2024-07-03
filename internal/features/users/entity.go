package users

import (
	"time"

	"github.com/labstack/echo/v4"
)

type Users struct {
	ID        uint
	Name      string
	Password  string
	Email     string
	Phone     string
	BirthDate time.Time `gorm:"type:date"`
}

type Handler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type Services interface {
	Register(newUser Users) error
	Login(email string, password string) (Users, string, error)
}

type Query interface {
	Register(newUser Users) error
	Login(email string) (Users, error)
}
