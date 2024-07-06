package repository

import (
	"be23/internal/features/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uint
	Name      string
	Password  string
	Email     string
	Phone     string
	BirthDate time.Time `gorm:"type:date"`
	// Todos     []Todo    `gorm:"foreignKey:Owner"`
}

func (u *User) toUserEntity() users.Users {
	return users.Users{
		ID:       u.ID,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Phone:    u.Phone,
	}
}

func toUserQuery(usersfromInputData users.Users) User {
	return User{
		Name:      usersfromInputData.Name,
		Email:     usersfromInputData.Email,
		Password:  usersfromInputData.Password,
		Phone:     usersfromInputData.Phone,
		BirthDate: usersfromInputData.BirthDate,
	}
}
