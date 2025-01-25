package main

import (
	"github.com/gofiber/fiber/v2"
)

// FlashMiddleware manages flash messages across requests.
func FlashMiddleware(c *fiber.Ctx) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return c.Next()
	}

	// Retrieve and clear flash message.
	flashMessage := sess.Get("flash_message")
	if flashMessage != nil {
		c.Locals("flash_message", flashMessage)
		sess.Delete("flash_message")
		sess.Save()
	}

	return c.Next()
}

// SetFlashMessage stores a flash message for the next request.
func SetFlashMessage(c *fiber.Ctx, message string) error {
	sess, err := SessionStore.Get(c)
	if err != nil {
		return err
	}
	sess.Set("flash_message", message)
	return sess.Save()
}
