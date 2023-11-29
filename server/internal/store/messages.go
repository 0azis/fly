package store

import (
	"chat/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type MessageService interface {
	InsertOne(msg models.Message) error
	DeleteOne(msgID int) error
}

type message struct {
	db *sqlx.DB
}

func (m message) InsertOne(msg models.Message) error {
	_, err := m.db.Query(fmt.Sprintf("insert into messages (chatid, from_, msg) values ('%s', %d, '%s')", msg.ChatID, msg.From, msg.Message))
	return err
}

func (m message) DeleteOne(msgID int) error {
	_, err := m.db.Query(fmt.Sprintf("delete from messages where messageid = %d", msgID))
	return err
}

