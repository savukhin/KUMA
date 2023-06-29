package api

import (
	"server/db/models"
	"server/db/query"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

func updateStatus(db *gorm.DB, validate *validator.Validate, onStopStatus func(string)) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &UpdateCncDTO{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		if err := validate.Struct(payload); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Incorrect data")
		}

		statusID, err := models.StringToStatusID(payload.Status)
		if err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, err.Error())
		}

		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(*UserClaims)
		cncID := claims.UserID

		cnc := query.Use(db).CncChecker
		cncChecker, err := cnc.
			Where(cnc.ID.Eq(int(cncID))).
			First()
		// Update(cnc.StatusID, statusID)

		if err != nil {
			return fiber.ErrForbidden
		}

		cncChecker.StatusID = statusID
		result, err := cnc.Updates(cncChecker)

		if result.Error != nil || err != nil {
			return fiber.ErrInternalServerError
		}

		if result.RowsAffected == 0 {
			c.Status(fiber.StatusForbidden)
			return c.SendString("No such string")
		}

		if statusID != models.WorkingStatus.ID && onStopStatus != nil {
			onStopStatus(cncChecker.Title)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}
