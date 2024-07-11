package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// LoginPage serves the login page
func LoginPage(c *fiber.Ctx) error {
	return c.SendFile("login.html")
}

// Login handles the session key submission
func Login(c *fiber.Ctx) error {
	sessionKey := c.FormValue("sessionKey")
	c.Cookie(&fiber.Cookie{
		Name:  "sessionKey",
		Value: sessionKey,
		Path:  "/",
	})
	return c.Redirect("/")
}
