package handlers

import (
	"goapi2/cmd/models"
	"goapi2/cmd/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateMeasurements(c echo.Context) error {
	meansurements := models.Measurements{}

	c.Bind(&meansurements)

	newMeansurements, err := repositories.CreateMeasurements(meansurements)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, newMeansurements)
}

func HandleUpdateMeasurement(c echo.Context) error {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	meansurements := models.Measurements{}
	c.Bind(&meansurements)

	updatedMeansurements, err := repositories.UpdateMeasurement(meansurements, idInt)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, updatedMeansurements)

}
