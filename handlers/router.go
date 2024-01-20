package handlers

import "net/http"

func Router() {
	http.HandleFunc("/", IndexHandler) // Обработчик для корневого пути
	http.HandleFunc("/ascii-art", AsciiHandler)
}