package main

import (
	"errors"
	"github.com/iot-thermometer/server/internal/client"
	"github.com/iot-thermometer/server/internal/controller"
	"github.com/iot-thermometer/server/internal/dto"
	"github.com/iot-thermometer/server/internal/repository"
	"github.com/iot-thermometer/server/internal/service"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Info("Error loading .env file")
	}

	config := dto.Config{DSN: os.Getenv("DSN"), Broker: os.Getenv("BROKER")}

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}

	e := echo.New()
	e.HideBanner = true
	e.HidePort = true
	e.Use(middleware.CORS())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if err != nil {
				var appError dto.AppError
				switch {
				case errors.As(err, &appError):
					return echo.NewHTTPError(400, err.Error())
				}
			}
			return err
		}
	})

	clients := client.NewClients(config)
	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories, config)
	controllers := controller.NewControllers(services)
	controllers.Route(e)

	err = clients.Broker().Connect("sensors", services.Reading().Handle)
	if err != nil {
		logrus.Panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	logrus.Info("Starting server on port " + port)
	logrus.Fatal(e.Start(":" + port))
}
