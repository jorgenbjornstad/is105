package main

import (
	"net/http"
	"log"
)

func main() {
	// Http-pakken kjører index-filen på port 8080.
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("web"))))
}

