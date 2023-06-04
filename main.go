package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

const config string = "root:==PaSsWoRd==@tcp(127.0.0.1:3306)/main_database"

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

func create(response http.ResponseWriter, request *http.Request) {
	template, err := template.ParseFiles(
		"templates/create.html",
		"templates/header.html",
		"templates/footer.html")

	if err != nil {
		fmt.Fprint(response, err.Error())
	}

	template.ExecuteTemplate(response, "create", nil)
}

func save_article(response http.ResponseWriter, request *http.Request) {
	defer http.Redirect(response, request, "/", http.StatusSeeOther)

	title := request.FormValue("title")
	anons := request.FormValue("anons")
	full_text := request.FormValue("full_text")

	db, err := sql.Open("mysql", config)
	assert(err)
	defer db.Close()

	insert, err := db.Query(fmt.Sprintf("INSERT INTO `articles` (`title`, `anons`, `full_text`) VALUES ('%s', '%s', '%s')",
		title, anons, full_text))
	assert(err)
	defer insert.Close()
}

func handler() {
	setupDataBase()

	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/create", create)
	http.HandleFunc("/save_article", save_article)

	http.ListenAndServe("localhost:80", nil)
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func setupDataBase() {
	db, err := sql.Open("mysql", config)
	assert(err)
	defer db.Close()

	createArticles, err := db.Query(`CREATE TABLE IF NOT EXISTS articles (
        id INT UNSIGNED NOT NULL AUTO_INCREMENT,
        title VARCHAR(255) NOT NULL,
        anons VARCHAR(255) NOT NULL,
		full_text TEXT NOT NULL,
		PRIMARY KEY (id)
    ) ENGINE = InnoDB`)
	assert(err)
	defer createArticles.Close()
}
