package handlers

import (
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/uptrace/bun"

	"github.com/hhuseyinpay/go-generic-repository-pattern/models"
	"github.com/hhuseyinpay/go-generic-repository-pattern/repository"
)

func NewUserHandler(db *bun.DB) UserHandler {
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

// GetByName godoc
// @Summary      Get user by name
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        name  path      string  true  "Name of User"
// @Success      200   {object}  models.User
// @Failure      400
// @Failure      404
// @Failure      500
// @Router       /users/{name} [get]
func (h UserHandler) GetByName(c *fiber.Ctx) error {
	name := c.Params("name")

	user, err := h.repository.GetByName(c.Context(), name)
	if err == sql.ErrNoRows {
		return notFoundResult(c)
	}
	if err != nil {
		return errorResult(c, err)
	}

	return successResult(c, user)
}
