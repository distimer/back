package authctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/deleteduser"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/utils/crypt"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type oauthLoginReq struct {
	Token      string `json:"token" validate:"required"`
	DeviceType *int8  `json:"device_type" validate:"required,min=0,max=1"`
}

func GoogleOauthLogin(c *fiber.Ctx) error {
	data := new(oauthLoginReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	claims, err := crypt.VerifyGoogleToken(data.Token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid Google Token",
		})
	}
	userName := claims.FamilyName + claims.GivenName
	if utf8.RuneCountInString(userName) > 20 {
		userName = truncateToRuneCharacters(userName, 20)
	}

	dbConn := db.GetDBClient()
	var findUser *ent.User
	new := false
	findUser, err = dbConn.User.Query().Where(user.And(user.OauthID(claims.SUB), user.OauthProvider(1))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			// check deleted user
			existUser, err := dbConn.DeletedUser.Query().
				Where(deleteduser.And(deleteduser.OauthProvider(1), deleteduser.OauthID(claims.SUB))).
				Exist(context.Background())
			if err != nil {
				logger.CtxError(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
			if existUser {
				return c.Status(409).JSON(fiber.Map{
					"error": "Quit user can re-register after 1 week",
				})
			}
			// Create User
			userID := uuid.New()
			findUser, err = dbConn.User.Create().
				SetID(userID).
				SetName(userName).
				SetOauthID(claims.SUB).
				SetOauthProvider(1).
				Save(context.Background())
			if err != nil {
				logger.CtxError(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}

			// add Default Category
			categoryObj, err := dbConn.Category.Create().
				SetName("미분류").
				SetUserID(userID).
				Save(context.Background())
			if err != nil {
				logger.CtxError(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}

			// add Default Subject
			_, err = dbConn.Subject.Create().
				SetName("미분류").
				SetColor("#000000").
				SetCategory(categoryObj).
				Save(context.Background())
			if err != nil {
				logger.CtxError(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
			new = true
		} else {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	// create new refresh token
	newRefreshToken := uuid.New()
	sessionObj, err := dbConn.Session.Create().
		SetRefreshToken(newRefreshToken).
		SetDeviceType(*data.DeviceType).
		SetUserID(findUser.ID).
		Save(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// create new access token
	newAccessToken := crypt.NewJWT(findUser.ID, findUser.TermsAgreed)

	if new {
		return c.Status(201).JSON(
			accessInfoDTO{
				SessionID:    sessionObj.ID.String(),
				UserID:       findUser.ID.String(),
				Name:         findUser.Name,
				AccessToken:  newAccessToken,
				RefreshToken: newRefreshToken.String(),
			},
		)
	}
	return c.JSON(
		accessInfoDTO{
			SessionID:    sessionObj.ID.String(),
			UserID:       findUser.ID.String(),
			Name:         findUser.Name,
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken.String(),
		},
	)
}

// @Summary Apple Oauth Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body authctrl.oauthLoginReq true "oauthLoginReq"
// @Success 200 {object} accessInfoDTO
// @Success 201 {object} accessInfoDTO
// @Failure 400
// @Failure 401
// @Failure 409
// @Failure 500
// @Router /auth/oauth/apple [post]
func AppleOauthLogin(c *fiber.Ctx) error {
	data := new(oauthLoginReq)
	if err := dto.Bind(c, data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err,
		})
	}

	claims, err := crypt.VerifyAppleToken(data.Token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid Apple Token",
		})
	}
	dbConn := db.GetDBClient()
	var findUser *ent.User
	new := false
	findUser, err = dbConn.User.Query().Where(user.And(user.OauthID(claims.Sub), user.OauthProvider(0))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			// check deleted user
			existUser, err := dbConn.DeletedUser.Query().
				Where(deleteduser.And(deleteduser.OauthProvider(0), deleteduser.OauthID(claims.Sub))).
				Exist(context.Background())
			if err != nil {
				logger.CtxError(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
			if existUser {
				return c.Status(409).JSON(fiber.Map{
					"error": "Quit user can re-register after 1 week",
				})
			}
			// Create User
			userID := uuid.New()
			findUser, err = dbConn.User.Create().
				SetID(userID).
				SetOauthID(claims.Sub).
				SetOauthProvider(0).
				Save(context.Background())
			if err != nil {
				logger.CtxError(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
			// add Default Category
			categoryObj, err := dbConn.Category.Create().
				SetName("미분류").
				SetUserID(userID).
				Save(context.Background())
			if err != nil {
				logger.CtxError(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}

			// add Default Subject
			_, err = dbConn.Subject.Create().
				SetName("미분류").
				SetColor("#000000").
				SetCategory(categoryObj).
				Save(context.Background())
			if err != nil {
				logger.CtxError(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
			new = true
		} else {
			logger.CtxError(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	// create new refresh token
	newRefreshToken := uuid.New()
	sessionObj, err := dbConn.Session.Create().
		SetRefreshToken(newRefreshToken).
		SetDeviceType(*data.DeviceType).
		SetUserID(findUser.ID).
		Save(context.Background())
	if err != nil {
		logger.CtxError(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// create new access token
	newAccessToken := crypt.NewJWT(findUser.ID, findUser.TermsAgreed)
	if new {
		return c.Status(201).JSON(
			accessInfoDTO{
				SessionID:    sessionObj.ID.String(),
				UserID:       findUser.ID.String(),
				Name:         findUser.Name,
				AccessToken:  newAccessToken,
				RefreshToken: newRefreshToken.String(),
			},
		)
	}
	return c.JSON(
		accessInfoDTO{
			SessionID:    sessionObj.ID.String(),
			UserID:       findUser.ID.String(),
			Name:         findUser.Name,
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken.String(),
		},
	)
}

func truncateToRuneCharacters(s string, limit int) string {
	if limit <= 0 {
		return ""
	}

	count := 0
	for i := range s {
		if count == limit {
			return s[:i]
		}
		count++
	}

	return s
}
