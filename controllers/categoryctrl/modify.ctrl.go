package categoryctrl

import "github.com/gofiber/fiber/v2"

type modifyCategoryReq struct {
	Name string `json:"name" validate:"required" example:"name between 1 and 20"`
}

func ModifyCategory(c *fiber.Ctx) error {
	return nil
}
