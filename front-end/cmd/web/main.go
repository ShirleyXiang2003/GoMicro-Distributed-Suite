package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", renderTemplates)

	fmt.Println("Starting front end service on port 80")
	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Panic(err)
	}

}

func renderTemplates(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"./cmd/web/templates/test.page.gohtml",
		"./cmd/web/templates/base.layout.gohtml",
	}

	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
