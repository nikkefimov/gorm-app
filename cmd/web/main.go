package main

import (
	"log"
	"net/http"
)

func main() {
	ConnectDB()
	http.HandleFunc("/home", loginPage)
	http.HandleFunc("/user_create", createUserHandler)
	http.HandleFunc("/movie_create", createMoviePage)

	log.Println("Server is running on http://localhost:8087/")
	log.Fatal(http.ListenAndServe(":8087", nil))

}
