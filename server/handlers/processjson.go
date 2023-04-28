package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"time"

	"github.com/labstack/echo"
)

func Processjson(c echo.Context) error {
	incomingData := new(models.IncomingData)
	if err := json.NewDecoder(c.Request().Body).Decode(incomingData); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid json format")
	}
	returnData := make(models.ReturnData)
	for group, users := range *incomingData {
		returnusers := make([]models.ReturnUser, 0, len(users))
		for _, user := range users {
			numericmonth, err := convmonth(user.BdayMonth)
			if err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "invalid json format")
			}
			userreturn := models.ReturnUser{
				Name:         user.Name,
				BdayMonth:    user.BdayMonth,
				NumericMonth: numericmonth,
			}
			returnusers = append(returnusers, userreturn)
		}
		returnData[group] = returnusers
	}
	return c.JSON(http.StatusOK, returnData)
}

func convmonth(month string) (int, error) {
	parsedMonth, err := time.Parse("January", month)
	if err != nil {
		return 0, err
	}
	return int(parsedMonth.Month()), nil
}
