package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
)

func main() {
	// Initialize Fiber template engine.
	engine := html.New("./public", ".html")

	// Create Fiber app with template engine.
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// Set Fiber middleware.
	app.Use(cors.New())
	app.Use(FlashMiddleware)

	// Initialize database.
	ConnectDB()

	// Initialize session management.
	InitSession()

	// Connect all static files to Fiber app.
	app.Static("/", "./public")

	// Public routes (access without login).
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{})
	})
	app.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("login", fiber.Map{})
	})
	app.Get("/register", func(c *fiber.Ctx) error {
		return c.Render("register", fiber.Map{})
	})

	// Authentication routes (access after login).
	app.Post("/register", RegisterHandler)
	app.Post("/login", LoginHandler)
	app.Get("/logout", LogoutHandler)

	// Routes protection.
	protected := app.Group("/", AuthMiddleware)
	protected.Get("/home", func(c *fiber.Ctx) error {
		sess, _ := SessionStore.Get(c)
		return c.Render("home", fiber.Map{
			"username": sess.Get("username"),
		})
	})
	protected.Get("/create-movie", func(c *fiber.Ctx) error {
		return c.Render("create-movie", fiber.Map{})
	})
	protected.Post("/create-movie", CreateMovieHandler)
	protected.Get("/find-movie", func(c *fiber.Ctx) error {
		return c.Render("find-movie", fiber.Map{})
	})
	protected.Get("/search-movie", FindMovieHandler)
	protected.Get("/movies", ListUserMoviesHandler)

	// Start server (set port your port if need).
	log.Fatal(app.Listen(":8087"))
}
