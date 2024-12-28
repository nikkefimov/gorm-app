package main

import (
	"html/template"
	"net/http"
	"strconv"
)

// Movie create handler.
func createMoviePage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userID, ok := session.Values["userID"].(uint)
	if !ok {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	if r.Method == "POST" {
		title := r.FormValue("title")
		year, _ := strconv.Atoi(r.FormValue("year"))
		genre := r.FormValue("genre")
		rating := r.FormValue("rating")

		movie := Movie{Title: title, Year: year, Genre: genre, Rating: rating, UserID: userID}
		if err := DB.Create(&movie).Error; err != nil {
			http.Error(w, "Error movie create",
				http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/movie_page", http.StatusFound)
		return
	}

	tmpl, _ := template.ParseFiles("public/movie_create.html")
	tmpl.Execute(w, nil)
}

// Movie page handler.
func moviePage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	userID, ok := session.Values["userID"].(uint)
	if !ok {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	// Looking movie in the list.
	var movies []Movie
	if err := DB.Where("user_id = ?", userID).Find(&movies).Error; err != nil {
		http.Error(w, "Error movies list",
			http.StatusInternalServerError)
		return
	}

	tmpl, _ := template.ParseFiles("template/movie_page.html")
	tmpl.Execute(w, movies)
}

// Find movie handler.
func findMoviePage(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, _ := strconv.Atoi(r.FormValue("id"))
		var movie Movie
		if err := DB.First(&movie, id).Error; err != nil {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}

		tmpl, _ := template.ParseFiles("public/movie_find.html")
		tmpl.Execute(w, movie)
		return
	}

	tmpl, _ := template.ParseFiles("public/movie_find.html")
	tmpl.Execute(w, nil)
}

// Logout handler.
func logoutPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	delete(session.Values, "userID")
	session.Save(r, w)

	http.Redirect(w, r, "/", http.StatusFound)
}
