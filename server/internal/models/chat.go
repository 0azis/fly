package models

type Chat struct {
	ID      string `json:"chatID" db:"chatid"`
	UserOne int    `json:"userOne" db:"user_one"`
	UserTwo int    `json:"userTwo" db:"user_two"`
}
