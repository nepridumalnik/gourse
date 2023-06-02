package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	handleRequests()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	user := struct {
		Name    string
		Age     uint
		Hobbies []string
	}{
		Name:    "Ivan",
		Age:     23,
		Hobbies: []string{"Ping-Pong", "Programming"},
	}

	t, err := template.ParseFiles("templates/index.html")

	if err != nil {
		fmt.Print(err)
	}

	err = t.Execute(w, user)

	if err != nil {
		fmt.Print(err)
	}
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1><b>Контакты</h1></b>")
}

func handleRequests() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/contacts", contactsPage)

	http.ListenAndServe(":8080", mux)
}
