package apis

import (
	"optx-server/apis/user"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Register(g *gin.RouterGroup)
}

var controllers = make([]Controller, 0)

func RegisterController(c Controller) {
	controllers = append(controllers, c)
}

func NewAPIServer() *APIServer {
	s := &APIServer{
		G: gin.Default(),
	}
	s.G.Use(gin.Recovery())
	s.G.Use(gin.ErrorLogger())
	api := s.G.Group("/api")
	{
		api.POST("/login", user.Login)
		api.POST("/logout", user.Logout)
		api.POST("/logon", user.Logon)
	}
	admin := api.Group("/admin")
	{
		admin.POST("/login", user.AdminLogin)
		admin.POST("/logout", user.AdminLogout)
		admin.POST("/logon", user.AdminLogon)
	}
	return s
}

type APIServer struct {
	G *gin.Engine
}
