package main

import (
	"myapp/config"
	controllers "myapp/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	network := echo.New()
	network.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	/* urls */
	network.POST("/addNote", controllers.AddNote)
	network.POST("/getNote", controllers.GetNote)
	network.GET("/getNotes", controllers.GetNotes)
	network.DELETE("/deleteNote", controllers.DeleteNote)
	/* urls */
	port := ":" + config.GoDotEnvVariable("APP_PORT")
	network.Logger.Fatal(network.Start(port))
}
