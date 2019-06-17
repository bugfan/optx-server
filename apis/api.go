package apis

import (
	"errors"
	"net/http"
	"optx-server/apis/option"
	"optx-server/apis/user"
	"optx-server/utils"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Register(g *gin.RouterGroup)
}

var controllers = make([]Controller, 0)

func RegisterController(c Controller) {
	controllers = append(controllers, c)
}
func Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		attr, err := utils.GetJWTDataFromCookie(c.Request)
		if err != nil || attr["Username"] == "" {
			// 过期或错cookie
			c.JSON(http.StatusUnauthorized, errors.New("身份验证失败"))
			c.Abort()
		}
	}
}

func NewAPIServer() *APIServer {
	s := &APIServer{
		G: gin.Default(),
	}

	s.G.Use(gin.Recovery())
	s.G.Use(gin.ErrorLogger())
	s.G.Static("/admin", "./panel/dist")
	api := s.G.Group("/api")
	{
		// 登录
		api.POST("/login", user.Login)
		api.POST("/logout", user.Logout)
		api.POST("/logon", user.Logon)
	}
	sms := api.Group("/sms")
	{
		// 小程序接口
		sms.GET("", option.GetSMS)
	}
	admin := api.Group("/admin")
	{
		admin.POST("/login", user.AdminLogin)
		admin.POST("/logout", user.AdminLogout)
		admin.POST("/logon", user.AdminLogon)
	}
	fun := s.G.Group("/api")
	fun.Use(Validate())
	{
		// 首页
		fun.GET("/index", user.Login)

		// 后台题接口
		fun.POST("/option", option.Create)
		fun.POST("/option/all", option.CreateLot) //批量
		fun.GET("/option", option.All)
		fun.DELETE("/option/:id", option.Delete)
		fun.PUT("/option/:id", option.Update)

	}

	return s
}

type APIServer struct {
	G *gin.Engine
}
