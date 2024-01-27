package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, r, "404")
		return
	}

	if r.Method != http.MethodGet {
		log.Println("404")
		ErrorPage(w, r, "404")
		return
	}

	tmpl, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		log.Printf("Error parsing template - index.html: %v", err)
		ErrorPage(w, r, "404")
		return
	}

	tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error execute template - index.html: %v", err)
		ErrorPage(w, r, "500")
	}
}
