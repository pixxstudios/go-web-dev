package main

import (
	"fmt"
	"net/http"

	"github.com/pixxstudios/go-web-dev/pkg/handlers"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/home", handlers.Home)

	// http.HandleFunc("/divide", divide)

	http.ListenAndServe(":8080", nil)
}
