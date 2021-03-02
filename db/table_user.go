package db

import "fmt"

func (p *PostgresClient) InsertUser(user *User) (id int64, err error) {
	var ID_OF_NOT_INSERTED_USER int64 = -1
	sqlStatement := fmt.Sprintf(`INSERT INTO %v (first_name, last_name, nickname, password) VALUES ($1, $2, $3, $4);`, p.tableNames.user)

	result, err := p.conn.Exec(sqlStatement, user.FirstName, user.LastName, user.Nickname, user.Password)
	if err != nil {
		return ID_OF_NOT_INSERTED_USER, err
	}

	id, err = result.LastInsertId()
	return id, err
}
