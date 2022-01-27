package repository

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/hhuseyinpay/go-generic-repository-pattern/models"
)

type IUserRepository interface {
	IBaseRepository[models.User]
	GetByName(ctx context.Context, name string) (models.User, error)
}

type UserRepository struct {
	BaseRepository[models.User]
}

func NewUserRepository(db *bun.DB) UserRepository {
	return UserRepository{
		BaseRepository[models.User]{db: db},
	}
}

func (r UserRepository) GetByName(ctx context.Context, name string) (models.User, error) {
	var m models.User
	err := r.db.NewSelect().Model(&m).Where("name = ?", name).Scan(ctx)
	return m, err
}
