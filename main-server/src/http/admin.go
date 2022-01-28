package http

import (
	"github.com/Mahanmmi/fuzzy-lamp/main-server/db/tables"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type AdminRegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Office   int16  `json:"office"`
}

func (s *Server) AdminRegister(c echo.Context) error {
	req := new(AdminRegisterRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	// create hash of the password
	hash, err := s.hashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	username, err := s.databases.Admins.Insert(tables.AdminsTableRecord{
		Username: req.Username,
		Password: hash,
		Office:   req.Office,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"username": username,
	})
}

type AdminLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (s *Server) AdminLogin(c echo.Context) error {
	req := new(AdminLoginRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	hash, err := s.hashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	admin, err := s.databases.Admins.GetByUsername(req.Username)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if admin.Password != hash {
		return c.JSON(http.StatusUnauthorized, "Invalid username or password")
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, &AdminClaims{
		Username: admin.Username,
		Office:   admin.Office,
	}).SignedString(s.conf.JWTSecret)

	return c.JSON(http.StatusOK, echo.Map{
		"token":    token,
		"username": admin.Username,
		"office":   admin.Office,
	})
}

type UserRegisterRequest struct {
	CardID   string `json:"card_id"`
	Password string `json:"password"`
	Office   int16  `json:"office"`
	Room     int16  `json:"room"`
	Light    int16  `json:"light"`
}

func (s *Server) UserRegister(c echo.Context) error {
	req := new(UserRegisterRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	admin := c.Get("user").(*jwt.Token)
	claims := admin.Claims.(*AdminClaims)

	if claims.Office != req.Office {
		return c.JSON(http.StatusUnauthorized, "You are not allowed to register users in this office")
	}

	// create hash of the password
	hash, err := s.hashPassword(req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	username, err := s.databases.Users.Insert(tables.UsersTableRecord{
		CardID:   req.CardID,
		Password: hash,
		Office:   req.Office,
		Room:     req.Room,
		Light:    req.Light,
	})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"username": username,
	})
}

func (s *Server) GetActivities(c echo.Context) error {
	admin := c.Get("user").(*jwt.Token)
	claims := admin.Claims.(*AdminClaims)

	activities, err := s.databases.Activities.GetByOfficeID(claims.Office)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, activities)
}

type SetOfficeLightTimesRequest struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
}

func (s *Server) SetOfficeLightTimes(c echo.Context) error {
	admin := c.Get("user").(*jwt.Token)
	claims := admin.Claims.(*AdminClaims)

	req := new(SetOfficeLightTimesRequest)
	if err := c.Bind(req); err != nil {
		return err
	}

	err := s.databases.Offices.Update(tables.OfficesTableRecord{
		ID:           claims.Office,
		LightOnTime:  req.Start,
		LightOffTime: req.End,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"start": req.Start,
		"end":   req.End,
	})
}
