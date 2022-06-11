package controllers

import (
	"context"
	"time"

	"github.com/millbj92/synctl/internal/cache"
	"github.com/millbj92/synctl/internal/database"
	"github.com/millbj92/synctl/pkg/models/auth"
	"github.com/millbj92/synctl/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// RenewTokens method: to renew expired tokens
// @Description Renew expired tokens
// @Summary Renew expired tokens
// @Tags token
// @Accept json
// @Produce json
// @Param refresh_token body string true "Refresh Token"
// @Success 200 {object} auth.Renew
// @Security ApiKeyAuth
// @Router /v1/tokens/renew/ [post]
func RenewTokens(c *fiber.Ctx) error {
	now := time.Now().Unix()

	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	expiresAccessToken := claims.Expires

	if now > expiresAccessToken {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error": true,
				"msg":   "Access token has expired",
			},
		)
	}

	renew := &auth.Renew{}

	if err := c.BodyParser(renew); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	expiresRefreshToken, err := utils.ParseRefreshToken(renew.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	if now < expiresRefreshToken {
		userID := claims.UserID

		db, err := database.Connect()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": true,
					"msg":   err.Error(),
				},
			)
		}

		found, err := db.GetUserByID(userID)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": true,
					"msg":   err.Error(),
				},
			)
		}

		credentials, err := utils.GetCredentialsByRole(found.Role)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				fiber.Map{
					"error": true,
					"msg":   err.Error(),
				},
			)
		}

		tokens, err := utils.GenerateNewTokens(userID.String(), credentials)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": true,
					"msg":   err.Error(),
				},
			)
		}

		connRedis, err := cache.Connect()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": true,
					"msg":   err.Error(),
				},
			)
		}

		errRedis := connRedis.Set(context.Background(), userID.String(), tokens.Refresh, 0).Err()
		if errRedis != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(
				fiber.Map{
					"error": true,
					"msg":   errRedis.Error(),
				},
			)
		}

		return c.JSON(fiber.Map{
			"error": false,
			"msg":   nil,
			"tokens": fiber.Map{
				"access":  tokens.Access,
				"refresh": tokens.Refresh,
			},
		})
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error": true,
				"msg":   "Session ended",
			},
		)
	}
}
