package server

import (
	"net/http"

	"github.com/BackendTigr/Backend2/internal/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Authenticate

// Simple -> login password (base64)

// jwt - access token и refresh token

// OAuth2 -> vk-trash, gmail, github

func (s *Server) Authenticate(c *fiber.Ctx) error {
	type user struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	var input user
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map {
			"error": fiber.ErrUnauthorized.Message,
		})
	}

	student, err := s.Deps.PG.//cruds не сделан, сначала надо сделать cruds а потом писать auth
}
