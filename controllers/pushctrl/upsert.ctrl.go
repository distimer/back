package pushctrl

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/apnstoken"
	"pentag.kr/distimer/ent/fcmtoken"
	"pentag.kr/distimer/ent/session"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type pushStartTokenDTO struct {
	SessionID  string `json:"session_id" validate:"required,uuid"`
	StartToken string `json:"start_token" validate:"required"`
	DeviceType *int8  `json:"device_type" validate:"required,min=0,max=1"`
}

func UpsertStartToken(c *fiber.Ctx) error {
	data := new(pushStartTokenDTO)
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
	_, err = dbConn.APNsToken.Delete().Where(apnstoken.HasSessionWith(session.ID(sessionObj.ID))).Exec(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}
	_, err = dbConn.FCMToken.Delete().Where(fcmtoken.HasSessionWith(session.ID(sessionObj.ID))).Exec(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	if *data.DeviceType == 0 {
		_, err = dbConn.APNsToken.Create().
			SetStartToken(data.StartToken).
			SetUpdateToken("").
			SetSession(sessionObj).
			Save(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	} else {
		_, err = dbConn.FCMToken.Create().
			SetPushToken(data.StartToken).
			SetSession(sessionObj).
			Save(context.Background())
		if err != nil {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	return c.SendStatus(200)
}
