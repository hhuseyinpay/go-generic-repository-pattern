package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/hhuseyinpay/go-generic-repository-pattern/repository"
)

type BaseHandler[T any] struct {
	baseRepository repository.IBaseRepository[T]
}

func (h BaseHandler[T]) GetAll(ctx *fiber.Ctx) error {
	m, err := h.baseRepository.GetAll(ctx.Context())
	if err != nil {
		return errorResult(ctx, err)
	}

	return successResult(ctx, m)
}

func (h BaseHandler[T]) PostCreate(ctx *fiber.Ctx) error {
	var m T
	err := ctx.QueryParser(&m)
	if err != nil {
		return errorResult(ctx, err)
	}
	m, err = h.baseRepository.Create(ctx.Context(), m)
	if err != nil {
		return errorResult(ctx, err)
	}

	return successResult(ctx, m)
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
