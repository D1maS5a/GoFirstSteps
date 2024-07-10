package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello web world!\n"))
}

func main() {
	http.HandleFunc("/", Handler)

	log.Println("Start HTTP server on 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
