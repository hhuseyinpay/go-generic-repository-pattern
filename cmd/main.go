package main

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/hhuseyinpay/go-generic-repository-pattern/router"
)

func main() {
	log.Println("bismillah")

	app := fiber.New()
	router.Setup(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("%v", err)
	}

	log.Println("bye bye")
}
