package controllers

import (
	"context"
	"time"

	"github.com/millbj92/synctl/internal/cache"
	"github.com/millbj92/synctl/internal/database"
	models "github.com/millbj92/synctl/pkg/models/auth"
	"github.com/millbj92/synctl/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateUser method: to create a new user (only accessible to admins)
// @Description Create a new user
// @Summary Create a new user
// @Tags user
// @Accept json
// @Produce json
// @Param email body string true "Email"
// @Param password body string true "Password"
// @Param first_name body string true "First Name"
// @Param last_name body string true "Last Name"
// @Param role body string true "Role"
// @Success 200 {object} auth.User
// @Router /v1/users/create/ [post]
func CreateUser(c *fiber.Ctx) error {
	create := &models.UserForCreate{}

	if err := c.BodyParser(create); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	validate := utils.NewValidator()

	if err := validate.Struct(create); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	role, err := utils.VerifyRole(create.Role)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	user := &models.User{}

	user.ID = uuid.New()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Email = create.Email
	user.Password = utils.GeneratePassword(create.Password)
	user.FirstName = create.FirstName
	user.LastName = create.LastName
	user.Role = role

	if err := validate.Struct(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	if _, err := db.GetUserByEmail(user.Email); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	user.Password = ""

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"user":  user,
	})
}

func UserLogin(c *fiber.Ctx) error {
	login := &models.UserForLogin{}

	if err := c.BodyParser(login); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	db, err := database.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	found, err := db.GetUserByEmail(login.Email)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error": true,
				"msg":   "Email or password is incorrect",
			},
		)
	}

	compareUserPassword := utils.ComparePasswords(found.Password, login.Password)
	if !compareUserPassword {
		return c.Status(fiber.StatusUnauthorized).JSON(
			fiber.Map{
				"error": true,
				"msg":   "Email or password is incorrect",
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

	tokens, err := utils.GenerateNewTokens(found.ID.String(), credentials)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	userID := found.ID.String()

	connRedis, err := cache.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	errSaveToRedis := connRedis.Set(context.Background(), userID, tokens.Refresh, 0).Err()
	if errSaveToRedis != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   errSaveToRedis.Error(),
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
}

func UserLogout(c *fiber.Ctx) error {
	claims, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	userID := claims.UserID.String()

	connRedis, err := cache.Connect()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   err.Error(),
			},
		)
	}

	errDelFromRedis := connRedis.Del(context.Background(), userID).Err()
	if errDelFromRedis != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{
				"error": true,
				"msg":   errDelFromRedis.Error(),
			},
		)
	}

	return c.SendStatus(fiber.StatusNoContent)
}
