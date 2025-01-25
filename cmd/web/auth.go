package auth

import "github.com/gofiber/fiber/v2"

// GetSessionUserID is fetching user's ID from session.
func GetSessionUserID(c *fiber.Ctx) uint {
	sess, _ := session.Get(c)
	userID := sess.Get("user_id")
	if userID != nil {
		return userID.(uint)
	}
	return 0
}

/*
// Create a new cookie store for storing session data securely.
var store = sessions.NewCookieStore([]byte("secret-key"))

// createUser handler.
func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the form was submitted.
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		// Hash the password using bcrypt.
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Error hashing password",
				http.StatusInternalServerError)
			return
		}
		// Create a new user object.
		user := User{
			Username: username,
			Password: string(hashedPassword),
		}

		// Save the user to the database
		if err := DB.Create(&user).Error; err != nil {
			http.Error(w, "Creating user error!",
				http.StatusInternalServerError)
			return
		}

		// After successfully creation log in automatically.
		session, _ := store.Get(r, "session-name")
		session.Values["userID"] = user.ID

		// Store the user ID in the session.
		session.Save(r, w)

		// Redirect to the home page.
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}

	// If its a GET request, proces the registration form.
	tmpl, _ := template.ParseFiles("public/user_create.html")
	tmpl.Execute(w, nil)
}

// loginPage handler.
func loginPage(w http.ResponseWriter, r *http.Request) {
	// Check if the form was submitted.
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

		// Success message.
		fmt.Fprintln(w, "You are logged in!")

		// Redirect to movie_create page after login.
		http.Redirect(w, r, "/movie_create", http.StatusFound)
		return
	}
	// Parse login page.
	tmpl, _ := template.ParseFiles("public/home.html")
	tmpl.Execute(w, nil)
}

*/
