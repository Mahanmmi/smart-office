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
