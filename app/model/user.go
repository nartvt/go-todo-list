package model

import "time"

type User struct {
	tableName   struct{} `sql:"users,alias:users" pg:",discard_unknown_columns"`
	Id          int
	UserName    string
	Password    string
	AccessToken string
	Email       string
	PhoneNumber string
	CreatedAt   time.Time
	CreatedBy   int
}
