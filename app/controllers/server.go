package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/tomotakashimizu/todoapp-go/config"
)

var templates = template.Must(template.ParseFiles("app/views/todos.html", "app/views/edit.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var validPath = regexp.MustCompile("^/(edit|update|delete)/([0-9]+)$")

func makeHandler(fn func(http.ResponseWriter, *http.Request, int)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			fmt.Println("NotFound")
			return
		}
		id, err := strconv.Atoi(m[2])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fn(w, r, id)
	}
}

func StartMainServer() error {
	http.HandleFunc("/todos/", getAllTodosHandler)
	http.HandleFunc("/create/", createTodoHandler)
	http.HandleFunc("/edit/", makeHandler(editTodoHandler))
	http.HandleFunc("/update/", makeHandler(updateTodoHandler))
	http.HandleFunc("/delete/", makeHandler(deleteTodoHandler))

	// サーバを起動
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
