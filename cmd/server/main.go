package main

import (
	"crm_system/config"
	"crm_system/internal/constants"
	"crm_system/internal/controllers"
	"crm_system/internal/repository"
	router "crm_system/internal/router"
	"crm_system/internal/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading environment variables")
	}

	app := fiber.New(fiber.Config{
		BodyLimit: constants.SIZE_LIMIT * 1024 * 1024, // this is the default limit of 4MB
	})

	// Initialize default config

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://192.168.0.139:3000", // no slash at end! , http://localhost:3000
		AllowCredentials: true,
	}))

	app.Static("/api/v1/public/static/", "./public/static/")

	// Initialize database
	config.InitDB()

	router := InitDI()
	router.Routes(app)

	// Fix: Properly start the server without forcing an exit
	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

func InitDI() *router.Router {

	//  Dependency Injection
	repositories := repository.NewRpository()
	services := services.NewService(repositories)
	controllers := controllers.NewController(services)
	r := router.NewRouter(controllers)

	// return router instance
	return r
}
