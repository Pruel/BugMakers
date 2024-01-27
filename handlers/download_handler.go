package handlers

import (
	"fmt"
	"net/http"
)

func DownloadAsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	// Здесь должен быть ваш код для генерации ASCII-арта или получения его из хранилища
	asciiArt := asciiArtCache

	if asciiArt == "" {
		http.Error(w, "ASCII art is empty or not found", http.StatusNotFound)
		return
	}

	// Установка необходимых заголовков для файла
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.txt\"", "ascii-art"))
	w.Header().Set("Content-Length", fmt.Sprint(len(asciiArt)))

	if _, err := w.Write([]byte(asciiArt)); err != nil {
		http.Error(w, "Error writing ASCII art to response", http.StatusInternalServerError)
	}
}
