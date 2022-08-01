package route

import (
	"github.com/gin-gonic/gin"
	"github.com/nartvt/go-todo-list/app/handler"
	"github.com/nartvt/go-todo-list/app/middleware"
)

func SetupRoute(r *gin.Engine) {
	r.Use(middleware.CORS())

	/*
		r.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})

		r.POST("/users", handler.User.CreateUser)
	*/
	todoRoute(r)
	r.Run()
}

func todoRoute(r *gin.Engine) {
	t := r.Group("todo_items")
	t.POST("items", handler.TodoItem.Create)
	t.PUT("items/:itemId", handler.TodoItem.Update)
	t.DELETE("items/:itemId", handler.TodoItem.Delete)
	t.GET("items/:itemId", handler.TodoItem.GetById)
	t.GET("items", handler.TodoItem.GetItems)
}
