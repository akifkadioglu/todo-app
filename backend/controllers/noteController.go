package controllers

import (
	"github.com/labstack/echo/v4"
	"myapp/messages"
	"myapp/models"
	"net/http"
)

func AddNote(c echo.Context) error {

	note := models.Note{
		Title:       c.QueryParam("title"),
		Description: c.QueryParam("description"),
	}
	db := ConnectDatabase()
	db.Create(&note)
	return c.JSON(http.StatusOK, note)
}

func GetNote(c echo.Context) error {
	var note models.Note

	db := ConnectDatabase()
	result := db.First(&note, c.QueryParam("id"))
	statusCode := 200
	if result.Error != nil {
		statusCode = 400
		return c.JSON(statusCode, map[string]string{
			"error": messages.DoesntExist,
		})
	}
	return c.JSON(statusCode, note)
}

func GetNotes(c echo.Context) error {
	var notes []models.Note
	db := ConnectDatabase()
	db.Find(&notes)
	return c.JSON(http.StatusOK, notes)
}

func DeleteNote(c echo.Context) error {
	var note models.Note
	db := ConnectDatabase()
	db.Delete(&note, c.QueryParam("id"))

	return c.JSON(http.StatusOK, "deleted")
}
