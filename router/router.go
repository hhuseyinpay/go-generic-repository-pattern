package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/uptrace/bun"

	"github.com/hhuseyinpay/go-generic-repository-pattern/handlers"
)

func Setup(app fiber.Router, db *bun.DB) {
	app.Use(logger.New())
	api := app.Group("/api")

	uh := handlers.NewUserHandler(db)
	api.Get("/users/create", uh.PostCreate) // Get method used here, just for simplicity
	api.Get("/users", uh.GetAll)
	api.Get("/users/:name", uh.GetByName)
}
