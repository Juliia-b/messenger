package db

/*

user: {
	id         SERIAL
	first_name text  NOT NULL
	last_name  text  NOT NULL
	nickname   text  NOT NULL UNIQUE
    password   bytea NOT NULL
	PRIMARY KEY(id)
}

chat: {
	id         SERIAL
	title      text     - available only for not private chats
    is_private boolean NOT NULL
}

chat_member: {
    chat_id  integer NOT NULL
	user_id  integer NOT NULL
	UNIQUE (chat, user_id) - Такое ограничение указывает,
                                    что сочетание значений перечисленных столбцов должно быть уникально во  всей таблице,
                                    тогда как значения каждого столбца по отдельности не должны быть (и обычно не будут)
                                    уникальными.
}

message: {
    id     SERIAL
	chat_id        integer NOT NULL
    sender_id      integer NOT NULL - is user who sends message into the chat
    content_text   text
    content_photo  bytea            - as []byte in base64
	date           bigint  NOT NULL
    UNIQUE (id, chat_id)
}

*/
