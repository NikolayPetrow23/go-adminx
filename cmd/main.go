package main

import (
	"log"
	"net/http"
	
	"go-adminx/internal/router"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", static))
	mux.HandleFunc("/", router.IndexRouter)
	mux.HandleFunc("/orderuid", router.OrderRouter)
	log.Println("Запуск сервера на http://127.0.0.1:8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
