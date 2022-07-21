package controllers

import (
	"net/http"

	"github.com/tomotakashimizu/todoapp-go/config"
)

func StartMainServer() error {
	files := http.FileServer(http.Dir(config.Config.Static))
	http.Handle("/static/", http.StripPrefix("/static/", files))

	return http.ListenAndServe(":"+config.Config.Port, nil)
}
