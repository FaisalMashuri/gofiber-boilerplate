package router

import (
	"github.com/Saucon/errcntrct"
	"github.com/gofiber/fiber/v2"
	"gofiber-boilerplate/config"
	user_handler "gofiber-boilerplate/handler/user"
	"gofiber-boilerplate/infrastructure"
	"gofiber-boilerplate/middleware"
)

// DomainHandler untuk register handler seriap domain
type DomainHandler struct {
	// Daftar Handler
	/*
		example :
		user.UserHandler
	*/
	user_handler.UserHandler
}

type RouterStruct struct {
	DomainHandler DomainHandler
}

func NewRouter(handler *DomainHandler) RouterStruct {
	return RouterStruct{
		DomainHandler: *handler,
	}
}

func (r *RouterStruct) SetupRoutes(app *fiber.App) {
	logger := infrastructure.NewLogCustom()

	if err := errcntrct.InitContract(config.AppConfig.ErrorContract.JSONPathFile); err != nil {
		logger.Fatal(err, "main : init contract", nil)
	}

	// Define routes with auth
	v1 := app.Group("/api/v1")
	// use middleware auth
	v1.Use()
	v1.Route("/", func(router fiber.Router) {
		//user
		user := router.Group("/user")
		user.Post("/", middleware.GetCredential, r.DomainHandler.UserHandler.Create)
	})
}
