package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		log.Printf("404")
		Error(w, "Method is not supported.", http.StatusNotFound) // Мы явно объявляем метод тем самым разделяя логику обработки HTTP
		// Запросов, тем самым помогая предотвратить неожиданное поведение и делая наш код более читаемым.
		// Разумеется мы можем не указывать метод вообще но в таком случае наш IndexHandler ответит на любой HTTP request
		// Что не соответствует практике разработки web приложений.
	}

	tmpl, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		log.Printf("Error parsing template - index.html: %v", err)
		Error(w, "Error parse template - index.html", http.StatusNotFound)
		return
	}

	tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error execute template - index.html: %v", err)
		Error(w, "Error execute template index.html", http.StatusInternalServerError)
	}
}
