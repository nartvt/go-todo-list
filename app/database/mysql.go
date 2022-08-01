package database

import (
	"database/sql"
	"fmt"

	"github.com/nartvt/go-todo-list/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var MySQL *gorm.DB
var sqlDB *sql.DB

func InitMysql() {
	conf := config.Config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.MySQL.UserName,
		conf.MySQL.Password,
		conf.MySQL.Host,
		conf.MySQL.Port,
		conf.MySQL.Database,
	)
	var err error
	MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

}

func CloseMysql() {
	if db, _ := MySQL.DB(); db != nil {
		if err := db.Close(); err != nil {
			fmt.Println("[ERROR] Cannot close mysql connection, err:", err)
		}
	}
}
