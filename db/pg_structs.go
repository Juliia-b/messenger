package db

import "database/sql"

/* --- PostgresClient --- */

type PostgresClient struct {
	conn       *sql.DB
	dbInfo     *dbInfo
	tablesInfo *tablesInfo
}

type dbInfo struct {
	dbName string
}

type tablesInfo struct {
	tableUser       string
	tableMessage    string
	tableChat       string
	tableChatMember string
}

/* ------- Tables ------- */

type User struct {
	ID        int64  `json="user_id"`
	FirstName string `json="first_name"`
	LastName  string `json="last_name"`
	Nickname  string `json="nickname"`
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
