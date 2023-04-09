package middlewares

import (
	"net/http"

	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
)

func FirebaseAuth(app *firebase.App) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get Firebase ID Token from Authorization header
		idToken := c.Get("Authorization")[7:] // Remove "Bearer " prefix

		// Verify ID Token
		fb, err := app.Auth(c.Context())
		token, err := fb.VerifyIDToken(c.Context(), idToken)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		// Add User to Context
		c.Locals("user", token)

		// Call next middleware
		return c.Next()
	}
}
