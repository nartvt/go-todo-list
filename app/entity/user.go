package entity

import (
	"fmt"
	"time"

	"github.com/nartvt/go-todo-list/app/auth"
	errhandler "github.com/nartvt/go-todo-list/app/error"
	"github.com/nartvt/go-todo-list/app/handler/request"
	"github.com/nartvt/go-todo-list/app/model"
	"github.com/nartvt/go-todo-list/app/orm"
)

var UserEntity IUser

func init() {
	UserEntity = user{}
}

type user struct{}

type IUser interface {
	CreateUser(user *request.UserRequest) (model.User, error)
}

func (user) CreateUser(input *request.UserRequest) (model.User, error) {
	if input == nil {
		return model.User{}, errhandler.BadRequestErr(fmt.Errorf("error must not be null"))
	}

	now := time.Now()
	user := model.User{
		UserName:    input.UserName,
		Password:    input.Password,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		AccessToken: input.AccessToken,
		CreatedAt:   now,
	}

	err := orm.User.CreateUser(&user)
	if err != nil {
		return model.User{}, errhandler.InternalError(err)
	}
	accessToken, err := auth.Auth.GenerateToken(user.Id, []byte(user.Password), []byte(user.Password))
	if err != nil {
		return model.User{}, err
	}
	user.AccessToken = accessToken
	err = orm.User.UpdateUser(&user)
	if err != nil {
		return model.User{}, errhandler.InternalError(err)
	}
	return user, nil
}
