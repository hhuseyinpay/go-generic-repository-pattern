package handlers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	"github.com/hhuseyinpay/go-generic-repository-pattern/models"
	"github.com/hhuseyinpay/go-generic-repository-pattern/repository"
)

func NewUserHandler(db *gorm.DB) UserHandler {
	repo := repository.NewUserRepository(db)
	uh := UserHandler{
		BaseHandler: BaseHandler[models.User]{
			baseRepository: repo,
		},
		repository: repo,
	}

	return uh
}

type UserHandler struct {
	BaseHandler[models.User]
	repository repository.IUserRepository
}

func (h UserHandler) GetByName(c *fiber.Ctx) error {
	name := c.Params("name")

	users, err := h.repository.GetByName(name)
	if err != nil {
		return errorResult(c, err)
	}

	return successResult(c, users)
}
