package repository

import (
	"gorm.io/gorm"
)

type IBaseRepository[T any] interface {
	Create(t T) (T, error)
	GetByID(id int64) (T, error)
	GetAll() ([]T, error)
	Update(t T) error
	Delete(t T) error
}

type BaseRepository[T any] struct {
	db *gorm.DB
}

func (r BaseRepository[T]) Create(t T) (T, error) {
	err := r.db.Create(&t).Error
	return t, err
}

func (r BaseRepository[T]) GetByID(id int64) (T, error) {
	var t T
	err := r.db.Where("id = ?", id).First(&t).Error
	return t, err
}

func (r BaseRepository[T]) GetAll() ([]T, error) {
	var t []T
	err := r.db.Find(&t).Error
	return t, err
}

func (r BaseRepository[T]) Update(t T) error {
	err := r.db.Save(&t).Error
	return err
}

func (r BaseRepository[T]) Delete(t T) error {
	err := r.db.Delete(&t).Error
	return err
}
