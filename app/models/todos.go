package models

import (
	"fmt"
	"log"
	"time"
)

type Todo struct {
	ID        int
	Title     string
	Content   string
	Completed string //trueなら"on"、falseなら"" （初期値は""）
	UserID    int
	CreatedAt time.Time
}

func (t *Todo) CreateTodo() error {
	cmd := `INSERT INTO todos(
		title,
		content,
		completed,
		user_id,
		created_at)
		values (?, ?, ?, ?, ?)
		`

	_, err := Db.Exec(cmd, t.Title, t.Content, "", t.UserID, time.Now())
	return err
}

func GetTodo(id int) (Todo, error) {
	cmd := `SELECT id, title, content, completed, user_id, created_at from todos where id = ?`
	row := Db.QueryRow(cmd, id)
	var todo Todo
	err := row.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Completed, &todo.UserID, &todo.CreatedAt)
	return todo, err
}

func GetAllTodos() ([]Todo, error) {
	cmd := `SELECT id, title, content, completed, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		return nil, err
	}
	var todos []Todo
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Content, &todo.Completed, &todo.UserID, &todo.CreatedAt)
		if err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("GetAllTodos: %w", err)
	}

	return todos, nil
}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `select id, title, content, completed, user_id, created_at from todos
	where user_id = ?`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Content,
			&todo.Completed,
			&todo.UserID,
			&todo.CreatedAt)

		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

func (t *Todo) UpdateTodo() error {
	cmd := `update todos set title = ?, content = ?, completed = ? 
	where id = ?`
	_, err := Db.Exec(cmd, t.Title, t.Content, t.Completed, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

func (t *Todo) DeleteTodo() error {
	cmd := `delete from todos where id = ?`
	_, err := Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
