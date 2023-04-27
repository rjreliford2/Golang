package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("About to start server")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!  ererer")
	})
	e.GET("/GetData22/:username/:password", GetData22)
	e.Logger.Fatal(e.Start(":1323"))
}

// JSON response
type Quote struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Time     string `json:"time"`
}

// e.GET("/GetData22/:username/:password", GetData22)
func GetData22(c echo.Context) error {
	quote := new(Quote)

	username := c.Param("username")
	quote.Username = username
	password := c.Param("password")
	quote.Password = password
	quote.Time = time.Now().Format(time.RFC3339)
	log.Println("in GetData22", quote)

	return c.JSON(http.StatusOK, quote)
}
