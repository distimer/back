package pushctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/apnstoken"
	"pentag.kr/distimer/ent/session"
	"pentag.kr/distimer/utils/dto"
)

type pushUpdateTokenDTO struct {
	SessionID   string `json:"session_id" validate:"required,uuid"`
	UpdateToken string `json:"update_token" validate:"required"`
}

func UpdateTokenUpsert(c *fiber.Ctx) error {
	data := new(pushUpdateTokenDTO)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	dbConn := db.GetDBClient()

	// check session exist
	sessionObj, err := dbConn.Session.Query().Where(session.ID(uuid.MustParse(data.SessionID))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			return c.Status(404).JSON(fiber.Map{
				"error": "Session not found",
			})
		}
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	err = dbConn.APNsToken.Update().Where(apnstoken.HasSessionWith(session.ID(sessionObj.ID))).SetUpdateToken(data.UpdateToken).Exec(context.Background())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	return c.SendStatus(200)
}
