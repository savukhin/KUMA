package api

import (
	"server/db/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AuthorizationHeaderName = "Authorization"
	AccessTokenHeaderName   = "access-token"
	RefreshTokenHeaderName  = "refresh-token"
)

type UserClaims struct {
	UserID    models.IDType `json:"user_id,omitempty"`
	ExpiresAt time.Time     `json:"exp,omitempty"`
	jwt.RegisteredClaims
}

type UserType models.CncChecker

func generateToken(userID models.IDType, expirationTime time.Time, secret interface{}) (string, time.Time, error) {
	// Create the JWT claims, which includes the username and expiry time
	claims := &UserClaims{
		UserID:    userID,
		ExpiresAt: expirationTime,
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Create the JWT string
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func generateAccessToken(userID models.IDType, secretKey interface{}) (string, time.Time, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(1 * time.Hour)

	return generateToken(userID, expirationTime, secretKey)
}

func generateRefreshToken(userID models.IDType, secretKey interface{}) (string, time.Time, error) {
	// Declare the expiration time of the token
	expirationTime := time.Now().Add(24 * time.Hour)

	return generateToken(userID, expirationTime, secretKey)
}

func GenerateTokens(userID models.IDType, secretKey interface{}) (access, refresh string, e error) {
	accessToken, _, err := generateAccessToken(userID, secretKey)
	if err != nil {
		e = err
		return
	}

	refreshToken, _, err := generateRefreshToken(userID, secretKey)
	if err != nil {
		e = err
		return
	}

	access = accessToken
	refresh = refreshToken
	return
}

func GenerateTokensAndSetHeaders(userID models.IDType, secretKey interface{}, c *fiber.Ctx) (e error) {
	access, refresh, err := GenerateTokens(userID, secretKey)

	if err != nil {
		return err
	}

	c.Response().Header.Add(AccessTokenHeaderName, access)
	c.Response().Header.Add(RefreshTokenHeaderName, refresh)

	return nil
}

func TokenRefresherMiddleware(secretKey interface{}) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if c.Locals("user") == nil {
			return c.Next()
		}

		// Gets user token from the context.
		u := c.Locals("user").(*jwt.Token)
		claims := u.Claims.(*UserClaims)

		// We ensure that a new token is not issued until enough time has elapsed
		// In this case, a new token will only be issued if the old token is within
		// 15 mins of expiry.
		if time.Until(claims.ExpiresAt) < 15*time.Minute {
			// Gets the refresh token from the cookie.
			refresh, ok := c.GetRespHeaders()[RefreshTokenHeaderName]
			// rc, err := c.Cookie(refreshTokenCookieName)
			if !ok {
				return c.Next()
			}

			// Parses token and checks if it valid.
			tkn, err := jwt.ParseWithClaims(refresh, claims, func(token *jwt.Token) (interface{}, error) {
				return secretKey, nil
			})
			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					return fiber.ErrUnauthorized
				}
			}

			if tkn != nil && tkn.Valid {
				// If everything is good, update tokens.
				_ = GenerateTokensAndSetHeaders(claims.UserID, secretKey, c)
			}

		}

		return c.Next()
	}
}
