package db

import (
	"database/sql"
	"fmt"
	"os"
)

// ConnectToPostgres opens a connection to PostgreSQL.
func ConnectToPostgres() (*PostgresClient, error) {
	var dbName = "messenger"
	dbConn, err := openDbByName(dbName) // psql: error: FATAL:  database "*" does not exist
	if err != nil {
		// handling a database connection error

		if !isErrorDbNotExist(err, dbName) {
			return nil, err
		}

		if tryCreateMissingDatabase(dbName) != nil {
			return nil, err
		}

		dbConn, err = openDbByName(dbName)
		if err != nil {
			return nil, err
		}
	}

	err = dbConn.Ping()
	if err != nil {
		return nil, err
	}

	return &PostgresClient{
		conn:       dbConn,
		dbInfo:     get_dbInfo(dbName),
		tablesInfo: get_tablesInfo(),
	}, nil
}

// SetUpDatabase adds the necessary tables to the database if they do not already exist.
// Sets the maximum number of open conn to the database.
// Sets the maximum number of conn in the idle connection pool.
func (p *PostgresClient) SetUpDatabase() error {
	p.conn.SetMaxOpenConns(25)
	p.conn.SetMaxIdleConns(25)

	err := createTablesIfNotExist(p.conn)
	return err
}

// CloseConnection closes the connection to the PostgreSQL.
func (p *PostgresClient) CloseConnection() {
	p.conn.Close()
}

func get_dbInfo(dbName string) *dbInfo {
	return &dbInfo{
		dbName: dbName,
	}
}

func get_tablesInfo() *tablesInfo {
	return &tablesInfo{
		tableUser:       "user",
		tableMessage:    "message",
		tableChat:       "chat",
		tableChatMember: "chat_member",
	}
}

// createTablesIfNotExist creates all necessary missing tables in the database.
func createTablesIfNotExist(dbConn *sql.DB) error {
	if err := createTableUser(dbConn); err != nil {
		return err
	}

	if err := createTableChat(dbConn); err != nil {
		return err
	}

	if err := createTableChatMembers(dbConn); err != nil {
		return err
	}

	if err := createTableMessage(dbConn); err != nil {
		return err
	}

	return nil
}

func createTableUser(dbConn *sql.DB) error {
	_, err := dbConn.Exec(`CREATE TABLE IF NOT EXISTS user ( id SERIAL, first_name text NOT NULL, last_name text NOT NULL, nickname text NOT NULL,  PRIMARY KEY(id) );`)
	return err
}

func createTableChat(dbConn *sql.DB) error {
	_, err := dbConn.Exec(`CREATE TABLE IF NOT EXISTS chat ( id SERIAL, title text NOT NULL, is_private boolean NOT NULL );`)
	return err
}

func createTableChatMembers(dbConn *sql.DB) error {
	_, err := dbConn.Exec(`CREATE TABLE IF NOT EXISTS chat_member ( chat_id  integer NOT NULL, user_id  integer NOT NULL, UNIQUE (chat_id, user_id) );`)
	return err
}

func createTableMessage(dbConn *sql.DB) error {
	_, err := dbConn.Exec(`CREATE TABLE IF NOT EXISTS message (id SERIAL, chat_id integer NOT NULL, sender_id integer NOT NULL, content_text text, content_photo text, date bigint  NOT NULL, UNIQUE (id, chat_id) );`)
	return err
}

//////////////////////////////////////////////////

func tryCreateMissingDatabase(dbName string) error {
	dbConnDefault, err := openDbByDefaultName()
	if err != nil {
		return err
	}

	if err := createDatabaseByName(dbConnDefault, dbName); err != nil {
		return err
	}

	return nil
}

func openDbByName(dbName string) (*sql.DB, error) {
	var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), dbName)

	return sql.Open("postgres", dataSourceName)
}

func openDbByDefaultName() (*sql.DB, error) {
	var dataSourceName = fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"))

	return sql.Open("postgres", dataSourceName)
}

// createDatabaseByName
func createDatabaseByName(dbConn *sql.DB, dbName string) error {
	var sqlStatement = fmt.Sprintf("SELECT 'CREATE DATABASE %v'\nWHERE NOT EXISTS (SELECT FROM pg_database WHERE datname = '%v')\\gexec", dbName, dbName)

	_, err := dbConn.Exec(sqlStatement)
	return err
}

// isErrorDbNotExist
func isErrorDbNotExist(openingError error, dbName string) (dbNotExist bool) {
	if openingError.Error() == fmt.Sprintf(`psql: error: FATAL:  database "%v" does not exist`, dbName) {
		return true
	}

	return false
}
