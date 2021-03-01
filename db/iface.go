package db

type DB interface {
	SetUpDatabase() error
	CloseConnection()
}

