package controllers

import (
	"net/http"

	"github.com/tomotakashimizu/todoapp-go/config"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	// URL /view/ の処理を関数で定義
	http.HandleFunc("/view/", viewHandler)

	// サーバを起動
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
