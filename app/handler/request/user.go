package request

import (
	"time"

	"github.com/gin-gonic/gin"
)

type UserRequest struct {
	UserName    string    `json:"user_name"`
	Password    string    `json:"password"`
	AccessToken string    `json:"access_token"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
}

func Bind(c *gin.Context) (UserRequest, error) {
	var user UserRequest
	if err := c.ShouldBind(&user); err != nil {
		return UserRequest{}, err
	}
	return user, nil
}
