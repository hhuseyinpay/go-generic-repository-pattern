package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hhuseyinpay/go-generic-repository-pattern/database"
	"github.com/hhuseyinpay/go-generic-repository-pattern/models"
	"github.com/hhuseyinpay/go-generic-repository-pattern/repository"
)

func UserList(c *fiber.Ctx) error {
	userRepostiory := repository.NewUserRepository(database.DB)
	users, err := userRepostiory.GetAll()
	if err != nil {
		return errorResult(c, err)
	}

	return successResult(c, users)
}

func UserByName(c *fiber.Ctx) error {
	name := c.Params("name")

	userRepostiory := repository.NewUserRepository(database.DB)
	users, err := userRepostiory.GetByName(name)
	if err != nil {
		return errorResult(c, err)
	}

	return successResult(c, users)
}

// UserCreate registers a user
func UserCreate(c *fiber.Ctx) error {
	user := models.User{
		Name: c.Query("name"),
	}

	userRepostiory := repository.NewUserRepository(database.DB)
	var err error

	user, err = userRepostiory.Create(user)
	if err != nil {
		return errorResult(c, err)
	}

	return successResult(c, user)
}

func errorResult(c *fiber.Ctx, err error) error {
	return c.Status(500).JSON(fiber.Map{
		"success": false,
		"error":   err.Error(),
	})
}
func successResult[T any](c *fiber.Ctx, t T) error {
	return c.JSON(fiber.Map{
		"success": true,
		"data":    t,
	})
}
