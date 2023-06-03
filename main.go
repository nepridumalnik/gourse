package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	handler()
}

func index(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles(
		"templates/index.html",
		"templates/header.html",
		"templates/footer.html")

	if err != nil {
		fmt.Fprint(response, err.Error())
	}

	template.ExecuteTemplate(response, "index", nil)
}

func handler() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)

	http.ListenAndServe("localhost:80", nil)
}
