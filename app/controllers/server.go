package controllers

import (
	"net/http"
	"todo_app/config"
)

/** サーバーの立ち上げ */
func StartMainServer() error {
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}