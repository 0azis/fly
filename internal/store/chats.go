package store

import (
	"chat/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ChatService interface {
	InsertOne(chat models.Chat) error
	DeleteOne(chatID string) error
	GetMyChats(userID int) ([]models.Chat, error)
	ChatHistory(chatID string) ([]models.Message, error)
}

type chat struct {
	db *sqlx.DB
}

func (c chat) InsertOne(chat models.Chat) error {
	_, err := c.db.Query(fmt.Sprintf("insert into chats values ('%s', %d, %d)", chat.ID, chat.UserOne, chat.UserTwo))
	return err
}

func (c chat) DeleteOne(chatID string) error {
	_, err := c.db.Query(fmt.Sprintf("delete from chats where chatid = '%s'", chatID))
	return err
}

func (c chat) GetMyChats(userID int) ([]models.Chat, error) {
	var myChats []models.Chat

	err := c.db.Select(&myChats, fmt.Sprintf("select * from chats where user_one = %d or user_two = %d", userID, userID))

	return myChats, err
}

func (c chat) ChatHistory(chatID string) ([]models.Message, error) {
	var chatHistory []models.Message

	err := c.db.Select(&chatHistory, fmt.Sprintf("select * from messages where chatid = '%s'", chatID))

	return chatHistory, err
}
