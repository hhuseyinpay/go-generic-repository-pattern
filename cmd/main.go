package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/hhuseyinpay/go-generic-repository-pattern/database"
	_ "github.com/hhuseyinpay/go-generic-repository-pattern/docs"
	"github.com/hhuseyinpay/go-generic-repository-pattern/router"
)

// @title Generic Repository Pattern
// @version 1.0
// @description This is a sample project for generic repostitory pattern
// @termsOfService http://swagger.io/terms/
// @contact.email hhuseyinpay@gmail.com
// @host localhost:3000
// @BasePath /api/
func main() {
	log.Println("bismillah")
	db, err := database.New()
	if err != nil {
		panic(err)
	}

	err = database.Migrate(context.Background(), db)
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	router.Setup(app, db)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("%v", err)
	}

	log.Println("hoşcagal")
}
