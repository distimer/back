package authctrl

import (
	"context"

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
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// @Summary Google Oauth Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body authctrl.oauthLoginReq true "oauthLoginReq"
// @Success 200 {object} loginRes
// @Router /auth/oauth/google [post]
func GoogleOauthLogin(c *fiber.Ctx) error {
	data := new(oauthLoginReq)
	if err := dto.Bind(c, data); err != nil {
		return err
	}

	claims, err := crypt.VerifyGoogleToken(data.Token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid Google Token",
		})
	}

	dbConn := db.GetDBClient()
	var findUser *ent.User
	findUser, err = dbConn.User.Query().Where(user.And(user.OauthID(claims.SUB), user.OauthProvider(1))).Only(context.Background())
	if err != nil {
		if ent.IsNotFound(err) {
			// Create User
			userID := uuid.New()
			findUser, err = dbConn.User.Create().
				SetID(userID).
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
				SetColor(0).
				SetCategory(categoryObj).
				Save(context.Background())
			if err != nil {
				logger.Error(c, err)
				return c.Status(500).JSON(fiber.Map{
					"error": "Internal server error",
				})
			}
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
	newAccessToken := crypt.NewJWT(findUser.ID)

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
// @Router /auth/oauth/apple [post]
func AppleOauthLogin(c *fiber.Ctx) error {
	data := new(oauthLoginReq)
	if err := dto.Bind(c, data); err != nil {
		return err
	}

	claims, err := crypt.VerifyAppleToken(data.Token)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid Apple Token",
		})
	}
	dbConn := db.GetDBClient()
	var findUser *ent.User
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
	newAccessToken := crypt.NewJWT(findUser.ID)

	return c.JSON(
		loginRes{
			UserID:       findUser.ID.String(),
			Name:         findUser.Name,
			AccessToken:  newAccessToken,
			RefreshToken: newRefreshToken.String(),
		},
	)
}
