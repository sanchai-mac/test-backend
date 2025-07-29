package server

import (
	"github.com/gofiber/fiber/v3"
)

func (s *Server) configRoute() {
	api := s.App.Group("/api")
	s.init(api)
}

func (s *Server) init(rg fiber.Router) {
	g := rg.Group("/v1/user")
	g.Get("/:id", s.Gateway.IUserController.GetUser)
}
