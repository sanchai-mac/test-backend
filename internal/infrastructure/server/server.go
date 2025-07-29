package server

import (
	"fmt"
	"log"
	"test-backend/internal/config"
	"test-backend/internal/controller"

	"github.com/gin-gonic/gin"
)

// Server ...
type Server struct {
	Gateway controller.Gateway
	config  *config.Configuration
	App     *gin.Engine
}

// StartRestful ...
func (s *Server) StartRestful() error {
	log.Println("[Server:StartRestful] Server listening on restful port: ", s.config.Port)
	addr := fmt.Sprintf("0.0.0.0:%s", s.config.Port)
	return s.App.Run(addr) // Run Gin server
}

// NewServer ...
func NewServer(
	g controller.Gateway,
	c *config.Configuration,
) *Server {
	engine := gin.New()
	engine.Use(gin.Logger(), gin.Recovery())

	s := &Server{
		App:     engine,
		Gateway: g,
		config:  c,
	}
	s.configRoute()
	return s
}
