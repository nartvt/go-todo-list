package request

import "github.com/gin-gonic/gin"

type TodoItem struct {
	Title   string
	Content string
	Status  string
}

func BindTodoItem(c *gin.Context) (TodoItem, error) {
	var item TodoItem
	if err := c.ShouldBind(&item); err != nil {
		return TodoItem{}, err
	}
	return item, nil
}
