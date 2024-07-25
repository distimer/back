package dto

import (
	"github.com/gofiber/fiber/v2"
)

type BaseDTO struct{}

// Bind binds the request body to the
func Bind(c *fiber.Ctx, b interface{}) error {
	if err := c.BodyParser(b); err != nil {
		return c.Status(400).JSON(fiber.Map{"errors": []string{err.Error()}})
	}
	if errs := myValidator.Validate(b); len(errs) > 0 {
		return c.Status(400).JSON(fiber.Map{"errors": errs})
	}
	return nil
}
