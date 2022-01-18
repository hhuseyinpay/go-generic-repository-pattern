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
			for _, i := range m {
				if _, err := db.NewCreateTable().Model(i).IfNotExists().Exec(ctx); err != nil {
					return err
				}
			}
			return nil
		})
	}

	down := func(ctx context.Context, db *bun.DB) error {
		return db.RunInTx(ctx, nil, func(ctx context.Context, tx bun.Tx) error {
			for _, i := range m {
				if _, err := db.NewDropTable().Model(i).Exec(ctx); err != nil {
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
