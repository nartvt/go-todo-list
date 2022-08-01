package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nartvt/go-todo-list/app/entity"
	errHandler "github.com/nartvt/go-todo-list/app/error"
	"github.com/nartvt/go-todo-list/app/handler/request"
	"github.com/nartvt/go-todo-list/app/handler/response"
)

type todoItem struct {
}

var TodoItem todoItem

func (todoItem) Create(c *gin.Context) {
	item, err := request.BindTodoItem(c)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}

	itemResp, err := entity.TodoItemEntity.Create(&item)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}
	c.JSON(http.StatusCreated, itemResp)
}

func (todoItem) Update(c *gin.Context) {
	item, err := request.BindTodoItem(c)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}
	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return

	}

	itemResp, err := entity.TodoItemEntity.UpdateById(&item, itemId)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}
	c.JSON(http.StatusCreated, itemResp)
}

func (todoItem) Delete(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return

	}

	itemResp, err := entity.TodoItemEntity.DeleteById(itemId)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}
	c.JSON(http.StatusCreated, itemResp)
}

func (todoItem) GetById(c *gin.Context) {
	itemId, err := strconv.Atoi(c.Param("itemId"))
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}

	itemResp, err := entity.TodoItemEntity.GetItemById(itemId)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}
	c.JSON(http.StatusCreated, itemResp)
}

func (todoItem) GetItems(c *gin.Context) {
	pagination := response.Pagination{}
	err := c.ShouldBind(&pagination)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}

	if pagination.Limit <= 0 {
		pagination.Limit = 1
	}
	if pagination.Page <= 0 {
		pagination.Page = 1
	}
	offset := (pagination.Page - 1) * pagination.Limit
	items, total, err := entity.TodoItemEntity.GetItems(pagination.Limit, offset)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}
	pagination.Data = items
	pagination.Total = int(total)
	if len(items) == pagination.Limit {
		pagination.Page += 1
	}
	c.JSON(http.StatusCreated, pagination)
}
