package todos

import (
	"github.com/labstack/echo/v4"
)

type Todo struct {
	ID       uint
	Activity string
	Mark     bool
	Owner    uint
}

type Handler interface {
	AddTodo() echo.HandlerFunc
	DeleteTodo() echo.HandlerFunc
	UpdateTodo() echo.HandlerFunc
	ShowTodo() echo.HandlerFunc
}

type Services interface {
	AddTodo(newData Todo) (Todo, error)
	DeleteTodo(dataToDelete Todo) (Todo, error)
	UpdateTodo(dataToUpdate Todo) (Todo, error)
	ShowTodo(ID uint) ([]Todo, error)
}

type Query interface {
	AddTodo(newData Todo) (Todo, error)
	DeleteTodo(dataToDelete Todo) error
	UpdateTodo(dataToUpdate Todo) (Todo, error)
	ShowTodo(ID uint) ([]Todo, error)
}
