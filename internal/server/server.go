package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tesla/tesla-bot-go/internal/deps"
)

type Server struct {
	App *fiber.App

	Deps *deps.Deps
}

func New(deps *deps.Deps) *Server {
	s := &Server{
		App: fiber.New(fiber.Config{}),

		Deps: deps,
	}

	s.App.Use(cors.New())

	// panic recovery
	// s.App.Use(recover.New(
	// 	recover.Config{
	// 		Next:             nil,
	// 		EnableStackTrace: true,
	// 	},
	// ))

	// status checks
	s.App.Get("/healthz", s.healthzHandler)

	// api := s.App.Group("/api") // localhost:8080/api !!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!

	return s
}

func (s *Server) healthzHandler(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

type response struct {
	Message string `json:"message"`
}

func (s *Server) ResponseOK(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(response{"ok"})
}

func (s *Server) InternalServerError(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusInternalServerError).
		JSON(response{err.Error()})
}

func (s *Server) BadRequest(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusBadRequest).
		JSON(response{err.Error()})
}

func (s *Server) Unauthorized(c *fiber.Ctx) error {
	return c.Status(fiber.StatusUnauthorized).JSON(response{fiber.ErrUnauthorized.Message})
}
