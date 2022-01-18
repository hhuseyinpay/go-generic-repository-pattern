package repository

import (
	"context"

	"github.com/uptrace/bun"
)

type IBaseRepository[T any] interface {
	Create(ctx context.Context, t T) (T, error)
	GetByID(ctx context.Context, id int64) (T, error)
	GetAll(ctx context.Context) ([]T, error)
	Update(ctx context.Context, t T) error
	Delete(ctx context.Context, id int64) error
}

type BaseRepository[T any] struct {
	db *bun.DB
}

func (r BaseRepository[T]) Create(ctx context.Context, t T) (T, error) {
	_, err := r.db.NewInsert().Model(&t).Exec(ctx)
	return t, err
}

func (r BaseRepository[T]) GetByID(ctx context.Context, id int64) (T, error) {
	var t T
	err := r.db.NewSelect().Model(&t).Where("id = ?", id).Scan(ctx)
	return t, err
}

func (r BaseRepository[T]) GetAll(ctx context.Context) ([]T, error) {
	var t []T
	err := r.db.NewSelect().Model(&t).Where("").Scan(ctx)
	return t, err
}

func (r BaseRepository[T]) Update(ctx context.Context, t T) error {
	_, err := r.db.NewUpdate().
		Model(&t).
		OmitZero().
		WherePK().
		Exec(ctx)
	return err
}

func (r BaseRepository[T]) Delete(ctx context.Context, id int64) error {
	var t T
	_, err := r.db.NewDelete().Model(&t).Where("id = ?", id).Exec(ctx)
	return err
}
