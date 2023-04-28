package main

import (
	"fmt"
	"net/http"
	"server/handlers"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("About to start server")
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hi there")
	})
	e.GET("/GetData22/:username/:password", handlers.GetData22)
	e.GET("/heartbeat", handlers.HeartBeat)
	e.POST("/processjson", handlers.ProcessJson)
	e.Logger.Fatal(e.Start(":1323"))

}
