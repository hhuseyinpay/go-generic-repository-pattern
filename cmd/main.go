package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/hhuseyinpay/go-generic-repository-pattern/database"
	"github.com/hhuseyinpay/go-generic-repository-pattern/router"
)

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

	log.Println("bye bye")
}
