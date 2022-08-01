package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	errHandler "github.com/nartvt/go-todo-list/app/error"
	"github.com/nartvt/go-todo-list/app/orm"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var Auth auth

type auth struct{}

const (
	secretKey = "secret"
	oneDay    = time.Hour * 24
)

func (au auth) Login(email string, password string) (string, error) {
	user, err := orm.User.GetUserByEmail(email)
	if err != nil && err == gorm.ErrRecordNotFound {
		return "", errHandler.NotFoundError(err)
	}
	if err != nil {
		return "", errHandler.InternalError(err)
	}
	return au.GenerateToken(user.Id, []byte(user.Password), []byte(password))
}

func (au auth) GenerateToken(userId int, hashPassword, password []byte) (string, error) {
	return au.generateToken(userId, hashPassword, password)
}

func (auth) generateToken(userId int, hashPassword, password []byte) (string, error) {
	if userId <= 0 || len(hashPassword) <= 0 || len(password) <= 0 {
		return "", errHandler.BadRequestErr(fmt.Errorf("input info token generate invalid"))
	}
	if err := bcrypt.CompareHashAndPassword(hashPassword, password); err != nil {
		return "", errHandler.BadRequestErr(err)
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userId)),
		ExpiresAt: time.Now().Add(oneDay).Unix(),
	})

	token, err := claims.SignedString([]byte(secretKey))
	return token, err
}
