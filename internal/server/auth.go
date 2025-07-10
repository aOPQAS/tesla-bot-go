package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/tesla/tesla-bot-go/internal/auth"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate

// Simple -> login password (base64)

// jwt - access token и refresh token

// OAuth2 -> vk-trash, gmail, github

func (s *Server) Authenticate(c *fiber.Ctx) error {
	type user struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var input user
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.ErrUnauthorized.Message,
		})
	}

	foundUser, err := s.Deps.PG.GetUserByEmail(input.Email)
	if err != nil || foundUser == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fiber.ErrUnauthorized.Message,
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(input.Password))
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": fiber.ErrUnauthorized.Message,
		})
	}

	accessToken, err := auth.GenerateAccessToken(foundUser.ID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": fiber.ErrUnauthorized.Message,
		})
	}

	refreshToken, err := auth.GenerateRefreshToken(foundUser.ID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": fiber.ErrUnauthorized.Message,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}
