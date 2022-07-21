package controllers

import (
	"net/http"

	"github.com/tomotakashimizu/todoapp-go/config"
)

func StartMainServer() error {
	// URL /view/ の処理を関数で定義
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	http.HandleFunc("/todos/", getAllTodosHandler)

	// サーバを起動
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
