package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "TheStoneSword"
)

type Users struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	PassHash string `json:"passHash"`
}

type UserEnter struct {
	Name     string `json:"name"`
	PassHash string `json:"passHash"`
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// urlExample := "postgres://postgres:1234@localhost:5432/TheStoneSword"
	// conn, err := pgx.Connect(context.Background(), urlExample)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	server := gin.Default()

	server.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "OK!",
		})
	})

	server.GET("/users", func(ctx *gin.Context) {

		users, err := allUsers(db)

		if err != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": users,
		})
	})

	server.POST("/users", func(ctx *gin.Context) {
		var user UserEnter
		ctx.BindJSON(&user)

		result, err := addUser(db, user.Name, user.PassHash)

		if err != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	})

	server.DELETE("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		result, err := deleteUser(db, id)

		if err != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": result,
		})

	})

	server.PUT("/users/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")

		var user UserEnter

		ctx.BindJSON(&user)

		result, err := updateUser(db, id, user)

		if err != nil {
			log.Fatal(err)
		}

		ctx.JSON(http.StatusOK, gin.H{
			"message": result,
		})
	})

	server.Run(":8080")

}

func allUsers(db *sql.DB) ([]Users, error) {

	rows, err := db.Query("select * from users;")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []Users

	for rows.Next() {
		var user Users

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

func addUser(db *sql.DB, login string, pass_hash string) (string, error) {
	sqlStatement := `insert into users (login_user, pass_hash) values ($1, $2);`
	_, err := db.Exec(sqlStatement, login, pass_hash)

	if err != nil {
		return "error", err
	}

	return "user is add", nil
}

func deleteUser(db *sql.DB, id string) (string, error) {
	sqlStatement := `delete from users where id_user = $1;`
	_, err := db.Exec(sqlStatement, id)

	if err != nil {
		return "error", err
	}

	return "user is deleted", nil
}

func updateUser(db *sql.DB, id string, user UserEnter) (string, error) {
	sqlStatement := `update users set login_user = $1, pass_hash = $2 where id_user = $3`
	_, err := db.Exec(sqlStatement, user.Name, user.PassHash, id)

	if err != nil {
		return "error", err
	}

	result := fmt.Sprintf("user id %s is updated", id)
	return result, nil
}
