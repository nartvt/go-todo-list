package response

import (
	"time"

	"github.com/nartvt/go-todo-list/app/model"
)

type UserView struct {
	Id          int       `json:"id"`
	UserName    string    `json:"user_name"`
	AccessToken string    `json:"access_token"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

func NewUserView(user model.User) UserView {
	return UserView{
		Id:          user.Id,
		UserName:    user.UserName,
		AccessToken: user.AccessToken,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		CreatedAt:   user.CreatedAt,
	}
}
