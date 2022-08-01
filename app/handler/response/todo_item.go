package response

import (
	"time"

	"github.com/nartvt/go-todo-list/app/model"
)

type TodoItem struct {
	Id        int
	Title     string
	Content   string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewViewTodoItem(item model.TodoItem) TodoItem {
	return TodoItem{
		Id:        item.Id,
		Title:     item.Title,
		Content:   item.Content,
		Status:    item.Status,
		CreatedAt: item.CreatedAt,
		UpdatedAt: item.UpdatedAt,
	}
}
