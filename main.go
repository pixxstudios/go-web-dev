package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// func divide(w http.ResponseWriter, r *http.Request) {
// 	f, err := divideValues(200.0, 0.0)

// 	if err != nil {
// 		fmt.Fprintf(w, fmt.Sprintf("%s", err))
// 		return
// 	}

// 	fmt.Fprintf(w, fmt.Sprintf("%f", f))
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y <= 0 {
// 		err := errors.New("Cannot divide by zero or negative number")
// 		return 0, err
// 	}

// 	result := x / y
// 	return result, nil
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	http.HandleFunc("/home", Home)

	// http.HandleFunc("/divide", divide)

	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home.page.tmpl")
}

func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)

	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("Error parsing template ", err)
		return
	}
}
