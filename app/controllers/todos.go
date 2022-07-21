package controllers

import (
	"html/template"
	"net/http"

	"github.com/tomotakashimizu/todoapp-go/app/models"
)

var todoTemplates = template.Must(template.ParseFiles("app/views/todos.html"))

func renderTodoTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := todoTemplates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getAllTodosHandler(w http.ResponseWriter, _ *http.Request) {
	todos, err := models.GetAllTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTodoTemplate(w, "todos", todos)
}
