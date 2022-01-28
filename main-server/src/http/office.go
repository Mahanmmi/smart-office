package http

import (
	"github.com/Mahanmmi/fuzzy-lamp/main-server/db/tables"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type OfficeRegisterRequest struct {
	ID           int16     `json:"id"`
	LightOnTime  time.Time `json:"light_on_time"`
	LightOffTime time.Time `json:"light_off_time"`
	APIKey       string    `json:"api_key"`
}

func (s *Server) OfficeRegister(c echo.Context) error {
	req := new(OfficeRegisterRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	id, err := s.databases.Offices.Insert(tables.OfficesTableRecord{
		ID:           req.ID,
		LightOnTime:  req.LightOnTime,
		LightOffTime: req.LightOffTime,
		APIKey:       req.APIKey,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"id": id,
	})
}

func (s *Server) GetLightTimes(c echo.Context) error {
	office, err := s.databases.Offices.GetByAPIKey(c.Request().Header.Get("Authorization"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"light_on_time":  office.LightOnTime.Format("15:04:05"),
		"light_off_time": office.LightOffTime.Format("15:04:05"),
	})
}

func (s *Server) CheckIn(c echo.Context) error {
	office, err := s.databases.Offices.GetByAPIKey(c.Request().Header.Get("Authorization"))
	if err != nil {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	cardId := c.QueryParam("cardid")

	user, err := s.databases.Users.GetByCardID(cardId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if user.Office != office.ID {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "user is not in this office",
		})
	}

	activity := tables.ActivitiesTableRecord{
		UserID:   user.ID,
		Office:   office.ID,
		Datetime: time.Now(),
	}

	if c.QueryParam("in") == "true" {
		activity.Type = tables.ActivityType_CheckIn
	} else if c.QueryParam("in") == "false" {
		activity.Type = tables.ActivityType_CheckOut
	} else {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "\"in\" query param must be true or false",
		})
	}

	_, err = s.databases.Activities.Insert(activity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, echo.Map{
		"user_id": user.ID,
		"light":   user.Light,
		"room":    user.Room,
	})
}
