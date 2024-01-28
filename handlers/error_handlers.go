package handlers

import (
	"html/template"
	"log"
	"net/http"
)

var errorTemplate *template.Template

func init() {
	var err error
	errorTemplate, err = template.ParseFiles("./web/templates/error.html")
	if err != nil {
		log.Printf("Error parsing error template: %v", err)
	}
}
func ErrorPage(w http.ResponseWriter, r *http.Request, status int) {
     message := http.StatusText(status)
	w.WriteHeader(status)

	err := renderErrorPage(w, message)
	if err != nil {
		log.Printf("Error rendering the error page: %v", err)
	}
}

func renderErrorPage(w http.ResponseWriter, message string) error {
	context := map[string]string{"Message": message}
	return errorTemplate.Execute(w, context)
}
