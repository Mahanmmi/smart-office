package http

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserLoginRequest struct {
	UserID   int16  `json:"user_id"`
	Password string `json:"password"`
}

func (s *Server) UserLogin(c echo.Context) error {
	req := new(UserLoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	hash, err := s.hashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	user, err := s.databases.Users.GetByID(req.UserID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if user.Password != hash {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		UserID: req.UserID,
	}).SignedString(s.conf.JWTSecret)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"token":   token,
		"user_id": user.ID,
		"office":  user.Office,
		"light":   user.Light,
		"room":    user.Room,
	})
}

type UpdateRoomLightRequest struct {
	LightVal int16 `json:"lights"`
}

func (s *Server) UpdateRoomLight(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*UserClaims)

	req := new(UpdateRoomLightRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	if fmt.Sprintf("%d", claims.UserID) != c.Param("userid") {
		return c.JSON(http.StatusUnauthorized, "Unauthorized")
	}

	err := s.databases.Users.UpdateLightByID(claims.UserID, req.LightVal)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "")
}
