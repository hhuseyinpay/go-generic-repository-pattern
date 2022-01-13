package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/hhuseyinpay/go-generic-repository-pattern/handlers"
)

func Setup(app fiber.Router) {
	app.Use(logger.New())
	api := app.Group("/api")

	api.Get("/users/create", handlers.UserCreate) // aynen kanka Post methodu olması gerektiğini bende biliyorum ama kim uğraşacak postmanla ;)
	api.Get("/users", handlers.UserList)
	api.Get("/users/:name", handlers.UserByName)
}
