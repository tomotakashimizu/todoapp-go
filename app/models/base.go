package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tomotakashimizu/todoapp-go/config"
)

var Db *sql.DB

const (
	tableNameUser = "users"
	tableNameTodo = "todos"
)

func init() {
	var err error
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)

	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)

	Db.Exec(cmdT)
}

func createUUID() uuid.UUID {
	uuidobj, _ := uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) string {
	cryptext := fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
