package repository

import (
	"gorm.io/gorm"

	"github.com/hhuseyinpay/go-generic-repository-pattern/models"
)

type IUserRepository interface {
	IBaseRepository[models.User]
	GetByName(name string) (models.User, error)
}

type UserRepository struct {
	BaseRepository[models.User]
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		BaseRepository[models.User]{db: db},
	}
}

func (r UserRepository) GetByName(name string) (models.User, error) {
	var m models.User
	err := r.db.Where("name = ?", name).First(&m).Error
	return m, err
}
