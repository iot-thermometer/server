package main

import (
	"github.com/iot-thermometer/server/internal/controller"
	"github.com/iot-thermometer/server/internal/repository"
	"github.com/iot-thermometer/server/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"os"
)

func main() {
	db := gorm.DB{}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories)
	controllers := controller.NewControllers(services)
	controllers.Route(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	logrus.Info("Starting server on port " + port)
	logrus.Fatal(e.Start(":" + port))
}
