package database

import (
	"database/sql"
	"fmt"

	"github.com/nartvt/go-todo-list/app/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Postgres *gorm.DB
var sqlPosgres *sql.DB

func InitPostgres() {
	conf := config.Config
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		conf.Posgres.UserName,
		conf.Posgres.Password,
		conf.Posgres.Host,
		conf.Posgres.Port,
		conf.Posgres.Database,
	)
	var err error
	Postgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func ClosePostgres() {
	if db, _ := Postgres.DB(); db != nil {
		if err := db.Close(); err != nil {
			fmt.Println("[ERROR] Cannot close mysql connection, err:", err)
		}
	}
}
