package controllers

import (
	"myapp/messages"
	"myapp/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddTask(c echo.Context) error {
	db := ConnectDatabase()

	task := models.Task{
		Title:       c.QueryParam("title"),
		Description: c.QueryParam("description"),
	}
	if len(task.Title) != 0 {
		db.Create(&task)
		return c.JSON(http.StatusOK, task)

	}
	return c.JSON(http.StatusBadRequest, map[string]string{
		"error": messages.CantBeNull,
	})
}

func GetTask(c echo.Context) error {
	var task models.Task

	db := ConnectDatabase()
	result := db.First(&task, c.QueryParam("id"))
	statusCode := 200
	if result.Error != nil {
		statusCode = 400
		return c.JSON(statusCode, map[string]string{
			"error": messages.DoesntExist,
		})
	}
	return c.JSON(statusCode, task)
}

func GetTasks(c echo.Context) error {
	var tasks []models.Task
	db := ConnectDatabase()
	db.Order("created_at desc").Find(&tasks)
	return c.JSON(http.StatusOK, tasks)
}

func DeleteTask(c echo.Context) error {
	var tasks []models.Task
	db := ConnectDatabase()
	db.Delete(&tasks, "id = ?", c.QueryParam("id"))

	return c.JSON(http.StatusOK, true)
}
