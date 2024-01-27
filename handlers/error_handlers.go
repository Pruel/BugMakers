package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func ErrorPage(w http.ResponseWriter, r *http.Request, errorCode string) {
	status, message := getErrorDetails(errorCode)
	// устанавливаем http-статусный код ответа
	w.WriteHeader(status)

	err := renderErrorPage(w, message)
	if err != nil {
		log.Printf("Error rendering the error page: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}

func getErrorDetails(errorCode string) (int, string) {
	switch errorCode {
	case "400":
		return http.StatusBadRequest, "Bad Request"
	case "404":
		return http.StatusNotFound, "Not Found"
	case "500":
		return http.StatusInternalServerError, "Internal Server Error"
	default:
		return http.StatusBadRequest, "Bad Request"
	}
}

// рендерит html-страницу ошибки с использованием шаблона
func renderErrorPage(w http.ResponseWriter, message string) error {

	tmpl, err := template.ParseFiles("web/templates/error.html")
	if err != nil {
		return err
	}

	context := map[string]string{"Message": message}
	return tmpl.Execute(w, context)
}
