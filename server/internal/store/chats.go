package store

import (
	"chat/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ChatService interface {
	InsertOne(chat models.Chat) error
	DeleteOne(chatID string) error
	GetMyChats(userID int) ([]models.ChatResponse, error)
	ChatHistory(chatID string) ([]models.Message, error)
	whoCompanion(chatID string, userID int) (int, error)
	lastMessage(chatID string) (string, error)
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

func (c chat) GetMyChats(userID int) ([]models.ChatResponse, error) {
	var myChats []models.Chat
	var myChatsResponse []models.ChatResponse

	err := c.db.Select(&myChats, fmt.Sprintf("select * from chats where user_one = %d or user_two = %d", userID, userID))

	for i := range myChats {
		companion, _ := c.whoCompanion(myChats[i].ID, userID)
		lastMessage, _ := c.lastMessage(myChats[i].ID)
		myChatsResponse = append(myChatsResponse, models.ChatResponse{
			ID: myChats[i].ID,
			Companion: companion,
			LastMessage: lastMessage,
		})
	}

	return myChatsResponse, err
}

func (c chat) ChatHistory(chatID string) ([]models.Message, error) {
	var chatHistory []models.Message

	err := c.db.Select(&chatHistory, fmt.Sprintf("select * from messages where chatid = '%s'", chatID))

	return chatHistory, err
}

func (c chat) whoCompanion(chatID string, userID int) (int, error) {
	var getChat models.Chat

	err := c.db.Get(&getChat, fmt.Sprintf("select * from chats where chatid = '%s'", chatID))

	if err != nil {
		return 0, err
	}

	if getChat.UserOne == userID {
		return getChat.UserTwo, nil
	}

	return getChat.UserOne, nil
}

func (c chat) lastMessage(chatID string) (string, error) {
	var lastMessage []models.Message

	err := c.db.Select(&lastMessage, fmt.Sprintf("select * from messages where chatid = '%s'", chatID))

	if lastMessage == nil {
		return "", nil
	}
	return lastMessage[len(lastMessage)-1].Message, err
}

