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
	id, err := strconv.ParseInt(c.QueryParam("id"), 10, 16)
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
