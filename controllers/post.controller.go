package controllers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/pimenvibritania/rest-api/models"
)

func FetchAllPost(c echo.Context) error {
	result, err := models.FetchAllPost()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
