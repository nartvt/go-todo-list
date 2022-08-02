package entity

import (
	"time"

	errHandler "github.com/nartvt/go-todo-list/app/error"
	"github.com/nartvt/go-todo-list/app/handler/request"
	"github.com/nartvt/go-todo-list/app/model"
	"github.com/nartvt/go-todo-list/app/orm"
	"gorm.io/gorm"
)

var TodoItemEntity ITodoItem

func init() {
	TodoItemEntity = todoItem{}
}

type todoItem struct{}

type ITodoItem interface {
	Create(item *request.TodoItem) (model.TodoItem, error)
	UpdateById(item *request.TodoItem, itemId int) (model.TodoItem, error)
	DeleteById(itemId int) (model.TodoItem, error)
	GetItemById(itemId int) (model.TodoItem, error)
	GetItems(limit, offset int) ([]model.TodoItem, int64, error)
}

func (todoItem) Create(item *request.TodoItem) (model.TodoItem, error) {
	if item == nil {
		return model.TodoItem{}, nil
	}

	now := time.Now()
	itemModel := model.TodoItem{
		Title:     item.Title,
		Status:    item.Status,
		Content:   item.Content,
		CreatedAt: now,
		UpdatedAt: now,
	}
	err := orm.TodoITem.Create(&itemModel)
	if err != nil {
		return model.TodoItem{}, err
	}
	return itemModel, nil
}

func (todoItem) UpdateById(input *request.TodoItem, itemId int) (model.TodoItem, error) {
	if input == nil {
		return model.TodoItem{}, nil
	}
	item, err := orm.TodoITem.GetItemById(itemId)
	if err == gorm.ErrRecordNotFound {
		return model.TodoItem{}, errHandler.NotFoundError(err)
	}

	item.Title = input.Title
	item.Status = input.Status
	item.Content = input.Content
	item.UpdatedAt = time.Now()
	err = orm.TodoITem.UpdateById(&item, itemId)
	if err != nil {
		return model.TodoItem{}, err
	}
	return item, nil
}

func (todoItem) DeleteById(itemId int) (model.TodoItem, error) {
	if itemId <= 0 {
		return model.TodoItem{}, nil
	}

	item, err := orm.TodoITem.GetItemById(itemId)
	if err == gorm.ErrRecordNotFound {
		return model.TodoItem{}, errHandler.NotFoundError(err)
	}
	err = orm.TodoITem.DeleteById(itemId)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (todoItem) GetItemById(itemId int) (model.TodoItem, error) {
	if itemId <= 0 {
		return model.TodoItem{}, nil
	}
	item, err := orm.TodoITem.GetItemById(itemId)
	if err == gorm.ErrRecordNotFound {
		return model.TodoItem{}, errHandler.NotFoundError(err)
	}
	if err != nil {
		return model.TodoItem{}, err
	}
	return item, nil
}

func (todoItem) GetItems(limit, offset int) ([]model.TodoItem, int64, error) {
	if limit <= 0 && offset <= 0 {
		return []model.TodoItem{}, 0, nil
	}
	items, total, err := orm.TodoITem.GetItems(limit, offset)
	if err != nil {
		return []model.TodoItem{}, 0, err
	}
	return items, total, nil
}
