package handlers

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	ascii "web/internal/ascii_art"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("400")
		Error(w, "Method is not supported", http.StatusBadRequest)
	}

	// Извлекаем данные из POST запроса
	text := r.FormValue("text")
	bannerName := r.FormValue("banner")

	banner, err := loadBanner(bannerName)
	if err != nil {
		log.Printf("Failed to load banner")
		Error(w, "Failed to load banner", http.StatusInternalServerError)
	}

	asciiArt, err := ascii.PrintAscii(banner, text)
	if err != nil {
		log.Printf("Failed print ascii")
		Error(w, "Failed print ascii", http.StatusInternalServerError)
	}

	// Send our result
	w.Header().Set("Content-Type", "text-plan")
	w.Write([]byte(asciiArt))
}

func loadBanner(banner string) (string, error) {
	filePath := filepath.Join("./internal/banner", banner+".txt")

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
