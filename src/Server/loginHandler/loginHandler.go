package loginHandler

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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler login called")

	// Open the database
	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user User
	// Decode JSON from the request body
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	// SQL query to check if user exists
	insertQuery := `SELECT 1 FROM users WHERE login = ? AND password = ? LIMIT 1`

	var result int
	err = db.QueryRow(insertQuery, user.Login, user.Password).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			// No matching user found
			http.Error(w, "Invalid login or password", http.StatusUnauthorized)
			return
		}
		// Some other error occurred
		http.Error(w, "Failed to check user existence", http.StatusInternalServerError)
		return
	}

	// If user exists, send a success response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}
