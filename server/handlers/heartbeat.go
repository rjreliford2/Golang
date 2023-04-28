package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func HeartBeat(c echo.Context) error {
	return c.JSON(http.StatusOK, "still breathing")
}
