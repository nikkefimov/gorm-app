package main

import (
	"log"
	"net/http"
)

func main() {
	//ConnectDB()
	http.HandleFunc("/login", loginPage)
	http.HandleFunc("/create_user", createUserHandler)

	log.Println("Server is running on http://localhost:8087/")
	log.Fatal(http.ListenAndServe(":8087", nil))

}
