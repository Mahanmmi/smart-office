package http

import (
	"github.com/Mahanmmi/fuzzy-lamp/main-server/db/tables"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
)

type OfficeRegisterRequest struct {
	ID           int16     `json:"id"`
	LightOnTime  time.Time `json:"light_on_time"`
	LightOffTime time.Time `json:"light_off_time"`
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
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"id": id,
	})
}

func (s *Server) GetLightTimes(c echo.Context) error {
	stringId, ok := s.conf.OfficeKeyIDMap[c.Request().Header.Get("Authorization")]
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	id, err := strconv.ParseInt(stringId, 10, 16)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	office, err := s.databases.Offices.GetByID(int16(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"light_on_time":  office.LightOnTime.Format("15:04:05"),
		"light_off_time": office.LightOffTime.Format("15:04:05"),
	})
}

func (s *Server) CheckIn(c echo.Context) error {
	stringId, ok := s.conf.OfficeKeyIDMap[c.Request().Header.Get("Authorization")]
	if !ok {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "unauthorized",
		})
	}

	id, err := strconv.ParseInt(stringId, 10, 16)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userId, err := strconv.ParseInt(c.Param("userid"), 10, 16)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := s.databases.Users.GetByID(int16(userId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	if user.Office != int16(id) {
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "user is not in this office",
		})
	}

	activity := tables.ActivitiesTableRecord{
		UserID:   int16(userId),
		Office:   int16(id),
		Datetime: time.Now(),
	}

	if c.Param("in") == "true" {
		activity.Type = tables.ActivityType_CheckIn
	} else if c.Param("in") == "false" {
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
		"user_id": userId,
		"light":   user.Light,
		"room":    user.Room,
	})
}
