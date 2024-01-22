package handlers

import (
	ascii "BugMakers/internal/ascii_art"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		//log.Printf("400")
		//Error(w, "Method is not supported", http.StatusBadRequest)
		ErrorPage(w, r, "400")
		return
	}

	// Извлекаем данные из POST запроса
	text := r.FormValue("text")
	bannerName := r.FormValue("banner")

	banner, err := loadBanner(bannerName)
	if err != nil {
		log.Printf("Failed to load banner: %v", err)
		//Error(w, "Failed to load banner", http.StatusInternalServerError)
		ErrorPage(w, r, "500")
		return
	}

	asciiArt, err := ascii.PrintAscii(banner, text)
	if err != nil {
		log.Printf("Failed print ascii: %v", err)
		//Error(w, "Failed print ascii", http.StatusInternalServerError)
		ErrorPage(w, r, "500")
		return
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
