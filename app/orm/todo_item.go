package orm

import (
	"github.com/nartvt/go-todo-list/app/database"
	"github.com/nartvt/go-todo-list/app/model"
)

var TodoITem ITodoItem

type todoItem struct{}

func init() {
	TodoITem = todoItem{}
}

type ITodoItem interface {
	Create(item *model.TodoItem) error
	UpdateById(item *model.TodoItem, itemId int) error
	DeleteById(itemId int) (model.TodoItem, error)
	GetItemById(itemId int) (model.TodoItem, error)
	GetItems(limit, offset int) ([]model.TodoItem, int64, error)
}

func (todoItem) Create(item *model.TodoItem) error {
	return database.MySQL.Create(item).Error
}

func (todoItem) UpdateById(item *model.TodoItem, itemId int) error {
	return database.MySQL.Model(item).
		Where("id = ?", itemId).
		Updates(item).
		Error
}

func (todoItem) DeleteById(itemId int) (model.TodoItem, error) {
	item := model.TodoItem{}
	err := database.MySQL.Where("id = ?", itemId).
		Delete(&item).
		Error
	return item, err
}

func (todoItem) GetItemById(itemId int) (model.TodoItem, error) {
	item := model.TodoItem{}
	err := database.MySQL.Where("id = ?", itemId).
		First(&item).
		Error
	return item, err
}

func (todoItem) GetItems(limit, offset int) ([]model.TodoItem, int64, error) {
	var items []model.TodoItem
	total := int64(0)
	err := database.MySQL.
		Count(&total).
		Offset(offset).
		Limit(limit).
		Order("id ASC").
		Select(&items).
		Error
	return items, total, err
}
