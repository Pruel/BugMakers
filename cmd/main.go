package main

import (
	"fmt"
	"log"
	"net/http"
	"web/handlers"
)

func main() {

	handlers.Router()

	fmt.Println("Open http://localhost:8080/\nExit Сtrl + C\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error load server", err)
	}
}
