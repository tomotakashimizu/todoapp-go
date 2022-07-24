package controllers

import (
	"html/template"
	"net/http"

	"github.com/tomotakashimizu/todoapp-go/config"
)

var templates = template.Must(template.ParseFiles("app/views/todos.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

// func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		m := validPath.FindStringSubmatch(r.URL.Path)
// 		if m == nil {
// 			http.NotFound(w, r)
// 			fmt.Println("NotFound")
// 			return
// 		}
// 		fn(w, r, m[2])
// 	}
// }

func StartMainServer() error {
	http.HandleFunc("/todos/", getAllTodosHandler)
	http.HandleFunc("/create/", createTodoHandler)

	// サーバを起動
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
