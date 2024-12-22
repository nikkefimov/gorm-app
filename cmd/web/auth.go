package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

// Create a new cookie store for storing session data securely.
var store = sessions.NewCookieStore([]byte("secret-key"))

// loginPage handler.
func loginPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		var user User
		if err := DB.Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		// Create the session.
		session, _ := store.Get(r, "session")
		// Set session data.
		session.Values["userID"] = user.ID
		// Save the session data.
		session.Save(r, w)

		// Redirect to movie_create page after login.
		http.Redirect(w, r, "/movie_create", http.StatusFound)
		return
	}
	// Parse login page.
	tmpl, _ := template.ParseFiles("templates/user_login.html")
	tmpl.Execute(w, nil)
}
