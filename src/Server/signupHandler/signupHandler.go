package signupHandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
)

type User struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	// Log when the handler is called
	fmt.Println("Handler signup Called")

	// Open SQLite database (this will create the file if it doesn't exist)
	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Check if 'users' table exists and create it if it doesn't
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			login TEXT NOT NULL,
			password TEXT NOT NULL
		);
	`
	_, err = db.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}

	// Parse JSON from the HTTP request body
	var user User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	// Insert the login and password into the users table
	insertQuery := `INSERT INTO users (login, password) VALUES (?, ?)`
	_, err = db.Exec(insertQuery, user.Login, user.Password)
	if err != nil {
		http.Error(w, "Failed to insert user into database", http.StatusInternalServerError)
		return
	}

	// Send a response back to the client
	w.WriteHeader(http.StatusOK)
}
