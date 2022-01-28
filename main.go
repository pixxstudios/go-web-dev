package main

import (
	"errors"
	"fmt"
	"net/http"
)

func divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(200.0, 0.0)

	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("%s", err))
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("%f", f))
}

func divideValues(x, y float32) (float32, error) {
	if y <= 0 {
		err := errors.New("Cannot divide by zero or negative number")
		return 0, err
	}

	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/divide", divide)

	http.ListenAndServe(":8080", nil)
}
