package main

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// User registration handler.
func RegisterHandler(c *fiber.Ctx) error {
	var user User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Validate user input.
	if !user.Validate() {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid username or password"})
	}

	// Check if username already exists.
	var existingUser User
	if result := db.Where("username = ?", user.Username).First(&existingUser); result.Error == nil {
		return c.Status(400).JSON(fiber.Map{"error": "Username already exists"})
	}

	// Hash password by bcrypt.
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not hash password"})
	}
	user.Password = string(hashedPassword)

	// Create user in DB.
	if result := db.Create(&user); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create user"})
	}

	// Set flash message after registration.
	SetFlashMessage(c, "Registration is successful. Please enter OK")

	return c.Redirect("/login")
}

// User login handler.
func LoginHandler(c *fiber.Ctx) error {
	var loginUser User
	if err := c.BodyParser(&loginUser); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	var user User
	// Search by username.
	if err := db.Where("username = ?", loginUser.Username).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "User not found"})
	}

	// Verify password.
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginUser.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Create user's session.
	sess, err := SessionStore.Get(c)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Session error"})
	}
	sess.Set("user_id", user.ID)
	sess.Set("username", user.Username)
	if err := sess.Save(); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not save session"})
	}

	return c.Redirect("/home")
}

// Logout handler ends user session.
func LogoutHandler(c *fiber.Ctx) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return c.Redirect("/login")
	}
	sess.Destroy()
	return c.Redirect("/login")
}

// Movie creation handler.
func CreateMovieHandler(c *fiber.Ctx) error {
	var movie Movie
	if err := c.BodyParser(&movie); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	// Get user ID from session.
	sess, err := SessionStore.Get(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	// Set user ID and set by genre.
	movie.UserID = userID.(uint)
	movie.Genre = processGenres(c.FormValue("genre"))

	// Validate movie function.
	if !movie.Validate() {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid movie data"})
	}

	// Create movie function.
	if result := db.Create(&movie); result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not create movie"})
	}

	// Set flash message after movie creation.
	SetFlashMessage(c, "Movie successfully created!")

	return c.Redirect("/home")
}

// Search for movies by title handler.
func FindMovieHandler(c *fiber.Ctx) error {
	title := c.Query("title")
	var movies []Movie

	// Search movies by title (case-insensitive matching).
	result := db.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%").Find(&movies)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not find movies"})
	}

	if len(movies) == 0 {
		SetFlashMessage(c, "Unfortunately, movie is not found")
		return c.Redirect("/find_movie")
	}

	return c.JSON(movies)
}

// Retrieves user's movies handler.
func ListUserMoviesHandler(c *fiber.Ctx) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	userID := sess.Get("user_id")
	if userID == nil {
		return c.Status(401).JSON(fiber.Map{"error": "Unauthorized"})
	}

	var movies []Movie
	result := db.Where("user_id = ?", userID).Find(&movies)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Could not retrieve movies"})
	}

	return c.JSON(movies)
}

// Process genre selection function.
func processGenres(genres string) string {
	// Limit by 3 genres and split.
	genreList := strings.Split(genres, ",")
	if len(genreList) > 3 {
		genreList = genreList[:3]
	}
	return strings.Join(genreList, ", ")
}
