package response

import (
	"time"

	"github.com/nartvt/go-todo-list/app/model"
)

type TodoItem struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

func NewViewTodoItems(items []model.TodoItem) []TodoItem {
	todoItems := make([]TodoItem, len(items))
	for i := range items {
		todoItems[i] = NewViewTodoItem(items[i])
	}
	return todoItems
}
