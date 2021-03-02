package db

import "database/sql"

/* --- PostgresClient --- */

type PostgresClient struct {
	conn       *sql.DB
	dbName     string
	tableNames *tableNames
}

type tableNames struct {
	user       string
	message    string
	chat       string
	chatMember string
}

/* ------- Tables ------- */

type User struct {
	ID        int64
	FirstName string
	LastName  string
	Nickname  string
	Password  []byte
}

type Chat struct {
	ID        int64  `json="chat_id"`
	Title     string `json="title"`
	IsPrivate bool   `json="is_private"`
}

type ChatMember struct {
	ChatID int64 `json="chat_id"`
	UserID int64 `json="user_id"`
}

type Message struct {
	ID           int64  `json="message_id"`
	ChatID       int64  `json="chat_id"`
	SenderID     int64  `json="sender_id"`
	ContentText  string `json="content_text"`
	ContentPhoto string `json="content_photo"`
	Date         int64  `json="date"`
}
