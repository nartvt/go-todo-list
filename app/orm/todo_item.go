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
	DeleteById(itemId int) error
	GetItemById(itemId int) (model.TodoItem, error)
	GetItems(limit, offset int) ([]model.TodoItem, int64, error)
}

func (todoItem) Create(item *model.TodoItem) error {
	return database.Postgres.Create(item).Error
}

func (todoItem) UpdateById(item *model.TodoItem, itemId int) error {
	return database.Postgres.
		Model(&item).
		Where("id = ?", itemId).
		Updates(item).
		Error
}

func (todoItem) DeleteById(itemId int) error {
	item := model.TodoItem{}
	return database.Postgres.
		Model(&model.TodoItem{}).
		Where("id = ?", itemId).
		Delete(&item).
		Error
}

func (todoItem) GetItemById(itemId int) (model.TodoItem, error) {
	item := model.TodoItem{}
	err := database.Postgres.
		Model(&item).
		Where("id = ?", itemId).
		First(&item).
		Error
	return item, err
}

func (todoItem) GetItems(limit, offset int) ([]model.TodoItem, int64, error) {
	var items []model.TodoItem
	total := int64(0)
	err := database.Postgres.Model(&items).
		Count(&total).
		Offset(offset).
		Limit(limit).
		Order("id ASC").
		Find(&items).
		Error
	return items, total, err
}
