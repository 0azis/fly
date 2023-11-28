package store

import (
	"chat/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserService interface {
	InsertOne(user models.User) (int, error)
	GetByNick(nick string) (models.User, error)
}

type user struct {
	db *sqlx.DB
}

func (u *user) InsertOne(user models.User) (int, error) {
	var newUserID int

	err := u.db.Get(&newUserID, fmt.Sprintf("insert into chat_users (nick, password) values ('%s', '%s') RETURNING userid", user.Nick, user.Password))

	return newUserID, err
}

func (u *user) GetByNick(nick string) (models.User, error) {
	var resultUser models.User

	err := u.db.Get(&resultUser, fmt.Sprintf("select * from chat_users where nick = '%s'", nick))

	return resultUser, err
}
