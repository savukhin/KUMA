package api

import (
	"crypto/rsa"

	"github.com/go-playground/validator"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB, privateKey *rsa.PrivateKey, logger *zap.Logger) *fiber.App {
	app := fiber.New()
	validate := validator.New()
	secretKey := privateKey.Public()

	api := app.Group("/api")
	v1 := api.Group("/v1")

	jwtMiddleware := jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{
			JWTAlg: jwtware.RS256,
			Key:    privateKey.Public(),
		},
		Claims: &UserClaims{},
	})

	auth := v1.Group("/auth")
	auth.Get("/login", login(db, secretKey, validate))
	auth.Get("/restricted", jwtMiddleware, restrictedUser)

	return app
}
