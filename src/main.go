package main

import (
	"./handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.Index)

	log.Fatalln(http.ListenAndServe(":7070", nil))
}
