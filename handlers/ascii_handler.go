package handlers

import (
	ascii "BugMakers/internal/ascii_art"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var (
	tmplAsciiArt  *template.Template
	asciiArtCache string
)

func init() {
	var err error
	tmplAsciiArt, err = template.ParseFiles("./web/templates/asciiArt.html")
	if err != nil {
		log.Printf("Error parsing template - ascii_art.html: %v", err)
	}
}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		log.Println("Method Not Allowed Ascii Handler")
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}

	text := r.FormValue("text")
	bannerName := r.FormValue("banner")
	bannerPath := fmt.Sprintf("./internal/banner/%s.txt", bannerName)

	banner, err := ascii.LoadBanner(bannerPath)
	if err != nil {
		log.Printf("Failed to load banner: %v", err)
		ErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	asciiArt, err := ascii.PrintAscii(banner, text)
	if err != nil {
		log.Printf("Failed print ascii: %v", err)
		ErrorPage(w, r, http.StatusInternalServerError)
		return
	}

	log.Printf("Received text: %s", text)
	log.Printf("Selected banner: %s", bannerName)
	log.Printf("Generated ASCII Art:\n%s \n", asciiArt)

	asciiArtCache = asciiArt
	data := map[string]string{
		"AsciiArt": asciiArtCache,
	}

	if err := tmplAsciiArt.Execute(w, data); err != nil {
		log.Printf("Error execute template - ascii_art.html: %v", err)
		ErrorPage(w, r, http.StatusInternalServerError)
	}
}
