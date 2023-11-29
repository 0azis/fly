package store

import (
	"chat/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserService interface {
	InsertOne(user models.User) (int, error)
	GetByLogin(login string) (models.User, error)
}

type user struct {
	db *sqlx.DB
}

func (u *user) InsertOne(user models.User) (int, error) {
	var newUserID int

	err := u.db.Get(&newUserID, fmt.Sprintf("insert into chat_users (nick, email, password) values ('%s', '%s', '%s') RETURNING userid", user.Nick, user.Email, user.Password))

	return newUserID, err
}

func (u *user) GetByLogin(login string) (models.User, error) {
	var resultUser models.User

	err := u.db.Get(&resultUser, fmt.Sprintf("select * from chat_users where nick = '%s' or email = '%s'", login, login))

	return resultUser, err
}

func (u *user) GetNameByID(userID int) (string, error) {
	var resultUser models.User

	err := u.db.Get(&resultUser, fmt.Sprintf("select * from chat_users where userid = %d", userID))

	return resultUser.Nick, err
}