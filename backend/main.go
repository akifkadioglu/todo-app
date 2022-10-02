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
	network.POST("/addTask", controllers.AddTask)
	network.POST("/getTask", controllers.GetTask)
	network.GET("/getTasks", controllers.GetTasks)
	network.DELETE("/deleteTask", controllers.DeleteTask)
	/* urls */
	port := ":" + config.GoDotEnvVariable("APP_PORT")
	network.Logger.Fatal(network.Start(port))
}
