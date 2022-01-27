package migrations

import (
	"context"

	"github.com/uptrace/bun"

	"github.com/hhuseyinpay/go-generic-repository-pattern/database/migrations/20220118103135_initial"
)

func init() {
	m := []interface{}{
		&migrations.User{},
	}
	up := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			if _, err := db.NewAddColumn().
				Model(&migrations.User{}).
				ColumnExpr("phone JSONB").
				Exec(ctx); err != nil {
				return err
			}
			return nil
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			for _, i := range m {
				if _, err := db.NewDropColumn().Model(i).ColumnExpr("phone").Exec(ctx); err != nil {
					return err
				}
			}
			return nil
		})
	}

	if err := Migrations.Register(up, down); err != nil {
		panic(err)
	}
}
