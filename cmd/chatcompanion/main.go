package main

import (
	"log"
	"net/http"

	"queuecompanion.com/chatcompanion/internal/server"
)

func main() {
	http.HandleFunc("/", server.Handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
