package api

import (
	"fmt"

	"github.com/go-playground/validator"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	fiber_logger "github.com/gofiber/fiber/v2/middleware/logger"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func JWTErrorChecker(c *fiber.Ctx, err error) error {
	fmt.Println("error = ", err)
	return err
}

func SetupRouter(db *gorm.DB, secretKey interface{}, logger *zap.Logger) *fiber.App {
	app := fiber.New()
	validate := validator.New()

	app.Use(fiber_logger.New())

	app.Get("/ping", func(c *fiber.Ctx) error { return c.SendString("Pong") })

	api := app.Group("/api")
	v1 := api.Group("/v1")

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.HS256,
			Key:    secretKey,
		},
		ErrorHandler: JWTErrorChecker,
		Claims:       &UserClaims{},
	})

	auth := v1.Group("/auth")
	auth.Get("/login", login(db, secretKey, validate))
	auth.Get("/restricted", jwtMiddleware, restrictedUser)

	return app
}
