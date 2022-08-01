package orm

import (
	"github.com/nartvt/go-todo-list/app/model"

	"github.com/nartvt/go-todo-list/app/database"
)

type user struct{}

var User IUser

func init() {
	User = user{}
}

type IUser interface {
	GetUserByEmail(email string) (model.User, error)
}

func (user) GetUserByEmail(email string) (model.User, error) {
	if len(email) < 0 {
		return model.User{}, nil
	}
	var user model.User
	err := database.MySQL.Model(&model.User{}).
		Where("email = ?", email).
		Select(&user).Error
	return user, err
}
