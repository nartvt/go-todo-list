package request

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoItem struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Status  string `json:"status"`
}

type QueryParams struct {
	Limit  int
	Page   int
	Offset int
}

const (
	defaultLimit = 8
	defaultPage  = 1
)

func BindTodoItem(c *gin.Context) (TodoItem, error) {
	var item TodoItem
	if err := c.ShouldBind(&item); err != nil {
		return TodoItem{}, err
	}
	return item, nil
}

func GetParams(c *gin.Context) QueryParams {
	queryParam := QueryParams{}
	limitStr := c.Query("limit")
	if len(limitStr) > 0 {
		limit, err := strconv.Atoi(limitStr)
		if err == nil {
			queryParam.Limit = limit
		}
	}

	pageStr := c.Query("page")
	if len(pageStr) > 0 {
		page, err := strconv.Atoi(pageStr)
		if err == nil {
			queryParam.Page = page
		}
	}

	if queryParam.Limit <= 0 {
		queryParam.Limit = defaultLimit
	}

	if queryParam.Page <= 0 {
		queryParam.Page = defaultPage
	}
	queryParam.Offset = (queryParam.Page - 1) * queryParam.Limit
	return queryParam
}
