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

	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var user User
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil {
		http.Error(w, "Unable to parse JSON", http.StatusBadRequest)
		return
	}

	insertQery := `SELECT 1 FROM users WHERE login = ? AND password = ? LIMIT = 1`
	_, err = db.Exec(insertQery, user.Login, user.Password)
	if err != nil {
		http.Error(w, "Failed to check if the user exists", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
