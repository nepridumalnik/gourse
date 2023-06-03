package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type User struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	Age  uint   `json:"age"`
}

func main() {
	// handleRequests()
	testDB()
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

	http.ListenAndServe(":80", mux)
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func testDB() {
	fmt.Println("Test")

	const config string = "root:==PaSsWoRd==@tcp(127.0.0.1:3306)/main_database"

	db, err := sql.Open("mysql", config)
	assert(err)
	defer db.Close()

	create, err := db.Query(`CREATE TABLE IF NOT EXISTS users (
        id INT UNSIGNED NOT NULL AUTO_INCREMENT,
        name VARCHAR(255) NOT NULL,
		age INT NOT NULL,
		PRIMARY KEY (id)
    ) ENGINE = InnoDB`)
	assert(err)
	defer create.Close()

	// insert, err := db.Query(`INSERT INTO users (name, age) VALUES ("Иван", 23)`)
	// assert(err)
	// defer insert.Close()

	res, err := db.Query(`SELECT * FROM users`)
	assert(err)

	for res.Next() {
		var user User

		err = res.Scan(&user.Id, &user.Name, &user.Age)
		assert(err)

		fmt.Printf("Name: %s, age: %d, id %d\n", user.Name, user.Age, user.Id)
	}

	fmt.Println("Подключено к MySQL!!!")
}
