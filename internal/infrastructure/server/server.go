package server

import (
	"fmt"
	"log"
	"test-backend/internal/config"
	"test-backend/internal/controller"

	fiber "github.com/gofiber/fiber/v3"
)

// Server ...
type Server struct {
	Gateway controller.Gateway
	config  *config.Configuration
	App     *fiber.App
}

// StartRestful ...
func (s *Server) StartRestful() error {
	log.Println("[Server:StartRestful] Server listening on restful port: ", s.config.Port)
	addr := fmt.Sprintf("0.0.0.0:%s", s.config.Port)
	return s.App.Listen(addr)
}

// NewServer ...
func NewServer(
	g controller.Gateway,
	c *config.Configuration,
) *Server {
	app := fiber.New()

	s := &Server{
		App:     app,
		Gateway: g,
		config:  c,
	}
	s.configRoute()
	return s
}
