package middlewares

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/utils/crypt"
)

func JWTMiddleware(c *fiber.Ctx) error {
	// get Bearer authHeader
	authHeader := strings.Split(c.Get("Authorization"), " ")
	if len(authHeader) != 2 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	token := authHeader[1]

	usreID, err := crypt.ParseIDJWT(token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Unauthorized",
		})
	}
	c.Locals("user-id", usreID) // type: uuid.UUID
	return c.Next()
}

func GetUserIDFromMiddleware(c *fiber.Ctx) uuid.UUID {
	return c.Locals("user-id").(uuid.UUID)
}
