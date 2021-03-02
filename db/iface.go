package db

type DB interface {
	SetUpDatabase() error
	CloseConnection()
	TableUser
}

type TableUser interface {
	InsertUser(user *User) (id int64, err error)
}
