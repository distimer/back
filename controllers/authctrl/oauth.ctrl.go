package authctrl

import (
	"context"
	"unicode/utf8"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"pentag.kr/distimer/db"
	"pentag.kr/distimer/ent"
	"pentag.kr/distimer/ent/user"
	"pentag.kr/distimer/utils/crypt"
	"pentag.kr/distimer/utils/dto"
	"pentag.kr/distimer/utils/logger"
)

type oauthLoginReq struct {
	Token string `json:"token" validate:"required"`
}

type loginRes struct {
	UserID       string `json:"user_id" validate:"required,uuid"`
	Name         string `json:"name" validate:"required"`
	AccessToken  string `json:"access_token" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required,uuid"`
}

// @Summary Google Oauth Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body authctrl.oauthLoginReq true "oauthLoginReq"
// @Success 200 {object} loginRes
// @Success 201 {object} loginRes
// @Failure 400
// @Failure 401
// @Failure 500
// @Router /auth/oauth/google [post]
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
			// Create User
			userID := uuid.New()
			findUser, err = dbConn.User.Create().
				SetID(userID).
				SetName(userName).
				SetOauthID(claims.SUB).
				SetOauthProvider(1).
				Save(context.Background())
			if err != nil {
				logger.Error(c, err)
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
				logger.Error(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}

			// add Default Subject
			_, err = dbConn.Subject.Create().
				SetName("미분류").
				SetColor("000000").
				SetCategory(categoryObj).
				Save(context.Background())
			if err != nil {
				logger.Error(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
			new = true
		} else {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	// create new refresh token
	newRefreshToken := uuid.New()
	_, err = dbConn.RefreshToken.Create().
		SetID(newRefreshToken).
		SetUserID(findUser.ID).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// create new access token
	newAccessToken := crypt.NewJWT(findUser.ID, false)

	if new {
		return c.Status(201).JSON(
			loginRes{
				UserID:       findUser.ID.String(),
				Name:         findUser.Name,
				AccessToken:  newAccessToken,
				RefreshToken: newRefreshToken.String(),
			},
		)
	}
	return c.JSON(
		loginRes{
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
// @Success 200 {object} loginRes
// @Success 201 {object} loginRes
// @Failure 400
// @Failure 401
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
			// Create User
			userID := uuid.New()
			findUser, err = dbConn.User.Create().
				SetID(userID).
				SetOauthID(claims.Sub).
				SetOauthProvider(0).
				Save(context.Background())
			if err != nil {
				logger.Error(c, err)
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
				logger.Error(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}

			// add Default Subject
			_, err = dbConn.Subject.Create().
				SetName("미분류").
				SetColor("000000").
				SetCategory(categoryObj).
				Save(context.Background())
			if err != nil {
				logger.Error(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
			new = true
		} else {
			logger.Error(c, err)
			return c.Status(500).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}
	}

	// create new refresh token
	newRefreshToken := uuid.New()
	_, err = dbConn.RefreshToken.Create().
		SetID(newRefreshToken).
		SetUserID(findUser.ID).
		Save(context.Background())
	if err != nil {
		logger.Error(c, err)
		return c.Status(500).JSON(fiber.Map{
			"error": "Internal server error",
		})
	}

	// create new access token
	newAccessToken := crypt.NewJWT(findUser.ID, false)
	if new {
		return c.Status(201).JSON(
			loginRes{
				UserID:       findUser.ID.String(),
				Name:         findUser.Name,
				AccessToken:  newAccessToken,
				RefreshToken: newRefreshToken.String(),
			},
		)
	}
	return c.JSON(
		loginRes{
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
