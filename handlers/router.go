package handlers

import (
	"log"
	"net/http"
)

func Router() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ascii-art", AsciiHandler)
	http.HandleFunc("/download", DownloadAsciiArtHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

	log.Println("Routes are configurated.")
}


