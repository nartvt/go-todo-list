package main

import (
	"github.com/nartvt/go-todo-list/app/database"

	"github.com/gin-gonic/gin"
	"github.com/nartvt/go-todo-list/app/route"
)

func main() {
	setupDatabase()
	setupRoute()
}

func setupRoute() {
	r := gin.Default()
	route.SetupRoute(r)
}

func setupDatabase() {
	// database.InitMysql()
	database.InitPostgres()
}
