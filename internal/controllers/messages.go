package controllers

import (
	"chat/internal/models"
	"chat/internal/store"
	"chat/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

type MessageInterface struct {
	Store store.InterfaceStore
}

func (mi MessageInterface) WriteMsg(c *gin.Context) {
	var newMsg models.Message

	if err := c.BindJSON(&newMsg); err != nil {
		c.JSON(500, "Плохая форма")
		return
	}

	userToken := strings.SplitN(c.Request.Header["Authorization"][0], " ", 2)[1]
	ID, _, _ := pkg.GetIdentity(userToken)

	newMsg.From = ID

	err := mi.Store.Messages().InsertOne(newMsg)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, "Ошибка при отправлении сообщения")
		return
	}

	c.JSON(200, "Сообщение отправлено")
}

func (mi MessageInterface) DeleteMsg(c *gin.Context) {
	msgID := c.Query("msg_id")
	msgIdInt, _ := strconv.Atoi(msgID)

	err := mi.Store.Messages().DeleteOne(msgIdInt)
	if err != nil {
		c.JSON(500, "Ошибка при удалении сообщения")
		return
	}
	c.JSON(200, "Сообщение удалено")
}

func GetMessageControllers(store store.InterfaceStore) MessageInterface {
	return MessageInterface{
		Store: store,
	}
}
