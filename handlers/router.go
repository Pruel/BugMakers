package handlers

import "net/http"

func Router() {
	http.HandleFunc("/", IndexHandler) // Обработчик для корневого пути
	http.HandleFunc("/ascii-art", AsciiHandler)
	http.HandleFunc("/download", DownloadAsciiArtHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
}
