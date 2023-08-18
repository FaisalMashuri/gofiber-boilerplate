package main

import (
	"fmt"
	user_handler "gofiber-boilerplate/handler/user"
	"gofiber-boilerplate/internal/domain/user/repository"
	"gofiber-boilerplate/internal/domain/user/service"
	"log"

	"github.com/gofiber/fiber/v2/middleware/cors"

	"gofiber-boilerplate/router"

	"github.com/gofiber/fiber/v2"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"gofiber-boilerplate/config"
	"gofiber-boilerplate/infrastructure"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize log
	//logClient := infrastructure.NewLogCustom()

	// Connect to the database
	db, err := infrastructure.ConnectDatabase()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Initialize repositories
	userRepository := repository.NewSubmissionPartnerRepository(db)

	//initialize services
	userServices := service.NewUserServicee(&userRepository)

	//initialize handlers
	userHandler := user_handler.NewUserHandler(&userServices)
	// Initialize Fiber app
	app := fiber.New()

	app.Static("/uploads", "./uploads")

	app.Use(logger.New(logger.Config{
		TimeZone: "Asia/Jakarta",
	}))
	app.Use(cors.New())
	//middleware.LogMiddleware(app, logClient)

	// Set up routes
	routerApp := router.NewRouter(&router.DomainHandler{
		UserHandler: userHandler,
	})
	routerApp.SetupRoutes(app)

	// Start the server
	err = app.Listen(fmt.Sprintf(":%s", config.AppConfig.AppConfig.Port))
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
