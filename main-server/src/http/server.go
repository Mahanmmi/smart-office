package http

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/Mahanmmi/fuzzy-lamp/main-server/config"
	"github.com/Mahanmmi/fuzzy-lamp/main-server/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	conf       *config.MainServerConfig
	databases  *db.MainServerDatabase
	echoServer *echo.Echo
}

func NewServer(conf *config.MainServerConfig, databases *db.MainServerDatabase) *Server {
	server := Server{
		conf:       conf,
		databases:  databases,
		echoServer: echo.New(),
	}
	server.echoServer.Use(middleware.Logger())
	server.echoServer.Use(middleware.Recover())
	server.echoServer.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	server.echoServer.POST("/api/office/register", server.OfficeRegister)
	server.echoServer.GET("/api/office/checkin", server.CheckIn)
	server.echoServer.GET("/api/office/lights", server.GetLightTimes)

	adminGroup := server.echoServer.Group("/api/admin")
	adminGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &AdminClaims{},
		SigningKey: server.conf.JWTSecret,
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/admin/register" ||
				c.Path() == "/api/admin/login" {
				return true
			}
			return false
		},
	}))
	adminGroup.POST("/register", server.AdminRegister)
	adminGroup.POST("/login", server.AdminLogin)
	adminGroup.POST("/user/register", server.UserRegister)
	adminGroup.GET("/activities", server.GetActivities)
	adminGroup.POST("/setlights", server.SetOfficeLightTimes)

	userGroup := server.echoServer.Group("/api/user")
	userGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &UserClaims{},
		SigningKey: server.conf.JWTSecret,
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/api/user/login" {
				return true
			}
			return false
		},
	}))
	userGroup.POST("/login", server.UserLogin)
	userGroup.POST("/:userid", server.UpdateRoomLight)

	return &server
}

func (s *Server) Start() {
	s.echoServer.Logger.Fatal(s.echoServer.Start(fmt.Sprintf(":%s", s.conf.HTTPServerPort)))
}

func (s *Server) hashPassword(password string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(h.Sum(nil)), nil
}
