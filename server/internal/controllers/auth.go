package controllers

import (
	"chat/internal/models"
	"chat/internal/store"
	"chat/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
)

type AuthControllers struct {
	Store store.InterfaceStore
}

func (a AuthControllers) SignUp(c *gin.Context) {
	var credentials models.Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(500, "Плохая форма")
		return
	}

	hashedPassword, err := pkg.Encode([]byte(credentials.Password))
	if err != nil {
		c.JSON(500, "Ошибка при хешировании пароля")
		return
	}

	newUser := models.User{
		Nick:     credentials.Nick,
		Password: string(hashedPassword),
	}

	newUserID, err := a.Store.Users().InsertOne(newUser)
	if err != nil {

		c.JSON(500, "Ошибка при создании пользователя")
		return
	}

	accessToken, err := pkg.CreateAccessToken(newUserID)
	if err != nil {
		c.JSON(500, "Ошибка при создании пользовательского токена")
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
	})
}

func (a AuthControllers) SignIn(c *gin.Context) {
	var credentials models.Credentials

	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(500, "Плохая форма")
		return
	}

	authUser, err := a.Store.Users().GetByNick(credentials.Nick)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, "Ошибка при чтении данных")
		return
	}

	if err = pkg.Decode([]byte(authUser.Password), []byte(credentials.Password)); err != nil {
		c.JSON(409, "Неправильные данные пользователя")
		return
	}

	accessToken, err := pkg.CreateAccessToken(authUser.ID)
	if err != nil {
		c.JSON(500, "Ошибка при создании пользовательского токена")
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
	})
}

func GetAuthControllers(store store.InterfaceStore) AuthControllers {
	return AuthControllers{
		Store: store,
	}
}
