package controllers

import (
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
