package serializers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/mamachengcheng/12306/app/middlewares"
	"github.com/mamachengcheng/12306/app/models"
	"github.com/mamachengcheng/12306/app/utils"
	"gorm.io/gorm"
)

type Login struct {
	Username string
	Password string
}

type Register struct {
	Username string
	Password string
	IDCard   string
	Name     string
	Phone    string
	Type     string
	Email    string
}

func (l Login) ValidCheck(c *gin.Context) error {
	return nil
}

func (r Register) ValidCheck(c *gin.Context) error {
	return nil
}

func (l Login) Login(c *gin.Context) (string, error) {
	l.Username = c.PostForm("username")
	l.Password = c.PostForm("password")

	var user models.User
	result := utils.MysqlDB.Where("username = ? AND password = ?", l.Username, l.Password).First(&user)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		token, err := middlewares.GenerateToken(l.Username)
		return token, err
	} else {
		return "", result.Error
	}
}

func (r Register) Register(ctx *gin.Context) error {
	return nil
}

