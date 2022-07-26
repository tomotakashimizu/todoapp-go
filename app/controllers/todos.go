package controllers

import (
	"fmt"
	"net/http"

	"github.com/tomotakashimizu/todoapp-go/app/models"
)

func getAllTodosHandler(w http.ResponseWriter, _ *http.Request) {
	todos, err := models.GetAllTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "todos", todos)
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("createTodoHandler", r.Method)
	if r.Method != "POST" {
		http.Error(w, "r.Method Error", http.StatusInternalServerError)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	todo := &models.Todo{Title: title, Content: content, UserID: 1}
	if err := todo.CreateTodo(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todos/", http.StatusFound)
}

func editTodoHandler(w http.ResponseWriter, r *http.Request, id int) {
	todo, err := models.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "edit", todo)
}

func updateTodoHandler(w http.ResponseWriter, r *http.Request, id int) {
	fmt.Println("updateTodoHandler", r.Method)
	if r.Method != "POST" {
		http.Error(w, "r.Method Error", http.StatusInternalServerError)
		return
	}
	title := r.FormValue("title")
	content := r.FormValue("content")
	fmt.Println("completed", r.FormValue("completed"))
	var completed int
	if r.FormValue("completed") == "on" {
		completed = 1
	} else {
		completed = 0
	}
	todo := models.Todo{ID: id, Title: title, Content: content, Completed: completed}
	if err := todo.UpdateTodo(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todos/", http.StatusFound)
}

func deleteTodoHandler(w http.ResponseWriter, r *http.Request, id int) {
	fmt.Println("deleteTodoHandler", r.Method)
	todo := models.Todo{ID: id}
	if err := todo.DeleteTodo(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/todos/", http.StatusFound)
}
