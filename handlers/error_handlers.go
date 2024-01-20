package handlers

import "net/http"

func Error(w http.ResponseWriter, errMsg string, status int) {
	http.Error(w, errMsg, status)
}