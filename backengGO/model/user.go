package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	PassHash string `json:"passHash"`
}

func AllUsers(db *sql.DB) ([]User, error) {

	rows, err := db.Query("select * from users;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		if err := rows.Scan(&user.Id, &user.Name, &user.PassHash); err != nil {
			return users, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return users, err
	}

	return users, nil
}

func AddUser(db *sql.DB, login string, pass_hash string) (string, error) {
	sqlStatement := `insert into users (login_user, pass_hash) values ($1, $2);`
	_, err := db.Exec(sqlStatement, login, pass_hash)

	if err != nil {
		return "error", err
	}

	return "user is add", nil
}

func DeleteUser(db *sql.DB, id string) (string, error) {
	sqlStatement := `delete from users where id_user = $1;`
	_, err := db.Exec(sqlStatement, id)

	if err != nil {
		return "error", err
	}

	return "user is deleted", nil
}

func UpdateUser(db *sql.DB, id string, user User) (string, error) {
	sqlStatement := `update users set login_user = $1, pass_hash = $2 where id_user = $3`
	_, err := db.Exec(sqlStatement, user.Name, user.PassHash, id)

	if err != nil {
		return "error", err
	}

	result := fmt.Sprintf("user id %s is updated", id)
	return result, nil
}
