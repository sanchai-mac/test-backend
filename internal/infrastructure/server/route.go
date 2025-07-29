package server

import "github.com/gin-gonic/gin"

func (s *Server) configRoute() {
	api := s.App.Group("/api")
	s.init(api)
}

func (s *Server) init(rg *gin.RouterGroup) {
	g := rg.Group("/v1/user")
	g.GET("/:user_id", s.Gateway.IUserController.GetUser)
	g.POST("/create", s.Gateway.IUserController.CreateUser)
	g.POST("/update/:user_id", s.Gateway.IUserController.UpdateUser)
	g.POST("/delete/:user_id", s.Gateway.IUserController.DeleteUser)
}
