package controllers

import (
	"chat/internal/models"
	"chat/internal/store"
	"chat/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

type ChatsControllers struct {
	Store store.InterfaceStore
}

func (cc ChatsControllers) Create(c *gin.Context) {
	friend := struct {
		ID int `json:"ID"`
	}{}

	if err := c.BindJSON(&friend); err != nil {
		c.JSON(500, "Плохая форма")
		return
	}

	userToken := strings.SplitN(c.Request.Header["Authorization"][0], " ", 2)[1]
	ID, _, _ := pkg.GetIdentity(userToken)

	myChats, _ := cc.Store.Chats().GetMyChats(ID)
	for i := range myChats {
		if myChats[i].Companion == friend.ID {
			c.JSON(403, "Такой чат уже есть")
			return
		}
	}

	chatInfo := models.Chat{
		ID:      pkg.GenerateUUID(),
		UserOne: ID,
		UserTwo: friend.ID,
	}

	err := cc.Store.Chats().InsertOne(chatInfo)
	if err != nil {
		c.JSON(500, "Ошибка при создании чата")
		return
	}
	c.JSON(200, "Чат успешно создан")
}

func (cc ChatsControllers) Delete(c *gin.Context) {
	chatID := c.Query("chat_id")

	err := cc.Store.Chats().DeleteOne(chatID)

	if err != nil {
		c.JSON(500, "Ошибка при удалении чата")
		return
	}
	c.JSON(200, "Чат успешно удален")
}

func (cc ChatsControllers) MyChats(c *gin.Context) {
	userToken := strings.SplitN(c.Request.Header["Authorization"][0], " ", 2)[1]
	ID, _, _ := pkg.GetIdentity(userToken)

	myChats, err := cc.Store.Chats().GetMyChats(ID)
	if err != nil {
		c.JSON(500, "Bad")
		return
	}
	
	c.JSON(200, myChats)
}

func (cc ChatsControllers) ChatHistory(c *gin.Context) {
	chatID := c.Param("chat_id")

	chatHistory, err := cc.Store.Chats().ChatHistory(chatID)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, "Ошибка при ототбражении сообщений")
		return
	}
	c.JSON(200, chatHistory)
}

func GetChatControllers(store store.InterfaceStore) ChatsControllers {
	return ChatsControllers{
		Store: store,
	}
}
