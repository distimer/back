package studylogctrl

// import (
// 	"time"

// 	"github.com/gofiber/fiber/v2"
// )

// func DailyGroupMemberLog(c *fiber.Ctx) error {

// 	dateStr := c.Params("date")
// 	date, err := time.Parse("2006-01-02", dateStr)
// 	if err != nil {
// 		return c.Status(400).JSON(fiber.Map{
// 			"error": "Invalid date format",
// 		})
// 	}

// 	return nil
// }
