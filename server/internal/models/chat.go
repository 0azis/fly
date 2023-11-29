package models

type Chat struct {
	ID      string `json:"chatID" db:"chatid"`
	UserOne int    `json:"userOne" db:"user_one"` // userOne who first create a chat 
	UserTwo int    `json:"userTwo" db:"user_two"`
}

type ChatResponse struct {
	ID string `json:"chatID"`
	Companion int `json:"companion"` // the name of owner the chat
	LastMessage string `json:"lastMessage"`
}
