package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nartvt/go-todo-list/app/entity"
	errhandler "github.com/nartvt/go-todo-list/app/error"
	"github.com/nartvt/go-todo-list/app/handler/request"
	"github.com/nartvt/go-todo-list/app/handler/response"
)

var User user

type user struct {
}

func (user) CreateUser(c *gin.Context) {
	userRequest, err := request.Bind(c)
	if err != nil {
		errhandler.HandlerErrorGin(c, err)
		return
	}

	userResp, err := entity.UserEntity.CreateUser(&userRequest)
	if err != nil {
		errhandler.HandlerErrorGin(c, err)
		return
	}
	c.JSON(http.StatusCreated, response.NewUserView(userResp))
}
