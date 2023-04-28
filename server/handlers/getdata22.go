package handlers

import (
	"log"
	"net/http"
	"server/models"
	"time"

	"github.com/labstack/echo"
)

func GetData22(c echo.Context) error {
	quote := new(models.Quote)

	username := c.Param("username")
	quote.Username = username
	password := c.Param("password")
	quote.Password = password
	quote.Time = time.Now().Format(time.RFC3339)
	log.Println("in GetData22", quote)

	return c.JSON(http.StatusOK, quote)
}
