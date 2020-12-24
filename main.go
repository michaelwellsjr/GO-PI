package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	tmpl := template.Must(template.ParseFiles("layout.html"))
	mux.HandleFunc("/plsdnthackme", func(w http.ResponseWriter, r *http.Request) {

		tmpl.Execute(w, "")
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
