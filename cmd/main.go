package main

import (
	"github.com/iot-thermometer/server/internal/controller"
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/repository"
	"github.com/iot-thermometer/server/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Panic("Error loading .env file")
	}

	config := dto.Config{DSN: os.Getenv("DSN")}

	db, err := gorm.Open(mysql.Open(config.DSN), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories, config)
	controllers := controller.NewControllers(services)
	controllers.Route(e)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4444"
	}
	logrus.Info("Starting server on port " + port)
	logrus.Fatal(e.Start(":" + port))
}
