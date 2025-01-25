package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

// SessionStore manages user's sessions.
var SessionStore *session.Store

// InitSession configures session management.
func InitSession() {
	SessionStore = session.New(session.Config{
		Expiration:     15 * time.Minute, // Session expires after 15 minutes of inactivity.
		CookiePath:     "/",
		CookieHTTPOnly: true,
		CookieSecure:   true, // Recommended for production by Fiber.
	})
}

// Checks user auth status.
func AuthMiddleware(c *fiber.Ctx) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return c.Redirect("/login")
	}

	// Check if user is logged in.
	if sess.Get("user_id") == nil {
		return c.Redirect("/login")
	}

	// Extend session on activity.
	sess.Set("last_activity", time.Now())
	sess.Save()

	return c.Next()
}
