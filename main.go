package main

import (
	"log"
	"mods/config"
	"mods/controller"
	"mods/middleware"
	"mods/repository"
	"mods/routes"
	"mods/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	db := config.SetupDatabaseConnection()

	carRepository := repository.NewCarRepository(db)
	carService := service.NewCarService(carRepository)
	CarController := controller.NewCarController(carService)

	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	routes.Routes(server, CarController)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
