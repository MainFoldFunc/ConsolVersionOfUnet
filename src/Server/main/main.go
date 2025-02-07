package main

import (
	"fmt"
	"github.com/MainFoldFunc/ConsolVersionOfUnet/tree/main/src/Server/loginHandler"
	"github.com/MainFoldFunc/ConsolVersionOfUnet/tree/main/src/Server/signupHandler"
	"net/http"
)

func main() {

	fmt.Println("Starting server on port 8080...")
	http.HandleFunc("/signup", signupHandler.SignupHandler)
	http.HandleFunc("/login", loginHandler.LoginHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error while running server: ", err)
	}
}
