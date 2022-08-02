package handler

import (
	"fmt"
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
	c.JSON(http.StatusCreated, response.NewViewTodoItem(itemResp))
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
	c.JSON(http.StatusOK, response.NewViewTodoItem(itemResp))
}

func (todoItem) Delete(c *gin.Context) {
	itemIdStr := c.Param("itemId")
	fmt.Println("AHIIII - " + itemIdStr)
	if len(itemIdStr) <= 0 {
		errHandler.HandlerErrorGin(c, errHandler.NotFoundError(fmt.Errorf("not found item")))
		return
	}
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return

	}

	itemResp, err := entity.TodoItemEntity.DeleteById(itemId)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}
	c.JSON(http.StatusOK, response.NewViewTodoItem(itemResp))
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
	c.JSON(http.StatusOK, response.NewViewTodoItem(itemResp))
}

func (todoItem) GetItems(c *gin.Context) {
	param := request.GetParams(c)
	items, total, err := entity.TodoItemEntity.GetItems(param.Limit, param.Offset)
	if err != nil {
		errHandler.HandlerErrorGin(c, err)
		return
	}

	if len(items) == param.Limit {
		pagination := response.Pagination{
			Data:  response.NewViewTodoItems(items),
			Total: int(total),
			Page:  param.Page + 1,
		}
		c.JSON(http.StatusCreated, pagination)
		return
	}
	c.JSON(http.StatusCreated, response.NewViewTodoItems(items))
}
