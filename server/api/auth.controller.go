package api

import (
	"server/db/query"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func login(db *gorm.DB, secretKey interface{}, validate *validator.Validate) fiber.Handler {
	return func(c *fiber.Ctx) error {
		payload := &LoginUserDTO{}

		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		if err := validate.Struct(payload); err != nil {
			return fiber.NewError(fiber.ErrBadRequest.Code, "Incorrect data")
		}

		cnc := query.Use(db).CncChecker
		checker, err := cnc.
			Where(cnc.Username.Eq(payload.Username)).
			First()

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		if err := bcrypt.CompareHashAndPassword([]byte(checker.PasswordHash), []byte(payload.Password)); err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		access, refresh, err := GenerateTokens(checker.ID, secretKey)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		return c.JSON(map[string]string{"access-token": access, "refresh-token": refresh})
		// return c.SendStatus(fiber.StatusOK)
	}
}

func restrictedUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(*UserClaims)
	userid := claims.UserID
	return c.SendString("Welcome " + strconv.FormatUint(uint64(userid), 10))
}
