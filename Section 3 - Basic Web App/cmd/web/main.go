package main

import (
	"github.com/drewdunne/go-course/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"


// main is the main application function
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/About", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)
}