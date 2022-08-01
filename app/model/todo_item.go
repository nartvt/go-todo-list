package model

import "time"

type TodoItem struct {
	tableName struct{} `sql:"todo_items,alias:todo_items" pg:",discard_unknown_columns"`
	Id        int
	Title     string
	Content   string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
