package models

import (
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

func (t *Todo) CreateTodo() error {
	cmd := `INSERT INTO todos(
		content,
		user_id,
		created_at)
		values (?, ?, ?)
		`

	_, err := Db.Exec(cmd, t.Content, t.UserID, time.Now())
	return err
}

func GetTodo(id int) (Todo, error) {
	cmd := `SELECT id, content, user_id, created_at from todos where id = ?`
	row := Db.QueryRow(cmd, id)
	var todo Todo
	err := row.Scan(&todo.ID, &todo.Content, &todo.UserID, &todo.CreatedAt)
	return todo, err
}
