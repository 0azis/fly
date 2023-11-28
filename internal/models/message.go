package models

type Message struct {
	ID        int    `json:"messageID" db:"messageid"`
	ChatID    string `json:"chatID" db:"chatid"`
	From      int    `json:"from" db:"from_"`
	Message   string `json:"message" db:"msg"`
	CreatedAt string `json:"createdAt" db:"created_at"`
}
