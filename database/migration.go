package database

import (
	"context"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/migrate"

	"github.com/hhuseyinpay/go-generic-repository-pattern/database/migrations"
)

func Migrate(ctx context.Context, db *bun.DB) error {
	m := migrations.Migrations
	migrator := migrate.NewMigrator(db, m)

	err := migrator.Init(context.Background())
	if err != nil {
		return err
	}
	// group, err := migrator.Rollback(context.Background())

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}

	if group.ID == 0 {
		fmt.Printf("there are no new migrations to run\n")
		return nil
	}

	fmt.Printf("migrated to %s\n", group)
	return nil
}
