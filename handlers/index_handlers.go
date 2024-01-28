package handlers

import (
	"log"
	"net/http"
	"text/template"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		log.Println("Method Not Allowed Index Handler")
		ErrorPage(w, r, http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		log.Printf("Error parsing template - index.html: %v", err)
		ErrorPage(w, r, http.StatusNotFound)
		return
	}

	
	if err := tmpl.Execute(w, nil); err != nil {
		log.Printf("Error execute template - index.html: %v", err)
		ErrorPage(w, r, http.StatusInternalServerError)
	}
}
