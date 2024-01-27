package handlers

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	ascii "web/internal/ascii_art"
)

var asciiArtCache string

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Printf("400")
		ErrorPage(w, r, "400")
		return
	}

	// Извлекаем данные из POST запроса
	text := r.FormValue("text")
	bannerName := r.FormValue("banner")

	bannerPath := fmt.Sprintf("./internal/banner/%s.txt", bannerName)

	banner, err := loadBanner(bannerPath)
	if err != nil {
		log.Printf("Failed to load banner: %v", err)
		ErrorPage(w, r, "500")
		return
	}

	log.Printf("Received text: %s", text)
	log.Printf("Selected banner: %s", bannerName)

	asciiArt, err := ascii.PrintAscii(banner, text)
	if err != nil {
		log.Printf("Failed print ascii: %v", err)
		ErrorPage(w, r, "500")
		return
	}
	log.Printf("Generated ASCII Art:\n%s", asciiArt)
	asciiArtCache = asciiArt

	data := map[string]string{
		"AsciiArt": asciiArtCache,
	}

	tmpl, err := template.ParseFiles("./web/templates/asciiArt.html")
	if err != nil {
		log.Printf("Error parsing template - ascii_art.html: %v", err)
		ErrorPage(w, r, "404")
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error execute template - ascii_art.html: %v", err)
		ErrorPage(w, r, "500")
	}
}

func loadBanner(banner string) (string, error) {
	file, err := os.Open(banner)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var builder strings.Builder
	for scanner.Scan() {
		builder.WriteString(scanner.Text())
		builder.WriteString("\n")
	}
	return builder.String(), nil
}
