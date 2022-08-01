package entity

import (
	"time"

	"github.com/nartvt/go-todo-list/app/handler/request"
	"github.com/nartvt/go-todo-list/app/model"
	"github.com/nartvt/go-todo-list/app/orm"
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

func (todoItem) UpdateById(item *request.TodoItem, itemId int) (model.TodoItem, error) {
	if item == nil {
		return model.TodoItem{}, nil
	}

	now := time.Now()
	itemModel := model.TodoItem{
		Id:        itemId,
		Title:     item.Title,
		Status:    item.Status,
		Content:   item.Content,
		UpdatedAt: now,
	}
	err := orm.TodoITem.UpdateById(&itemModel, itemId)
	if err != nil {
		return model.TodoItem{}, err
	}
	return itemModel, nil
}

func (todoItem) DeleteById(itemId int) (model.TodoItem, error) {
	if itemId <= 0 {
		return model.TodoItem{}, nil
	}
	item, err := orm.TodoITem.DeleteById(itemId)
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
