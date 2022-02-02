package main

import (
	"fmt"
	"os"

	"github.com/uptrace/bun/migrate"
	"github.com/urfave/cli/v2"

	"github.com/hhuseyinpay/go-generic-repository-pattern/database"
	"github.com/hhuseyinpay/go-generic-repository-pattern/database/migrations"
)

func main() {
	app := &cli.App{
		Name: "bun",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "env",
				Value: "dev",
				Usage: "environment",
			},
		},
		Commands: commands,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}

var commands = []*cli.Command{
	{
		Name:  "init",
		Usage: "create migration tables",
		Action: func(c *cli.Context) error {
			migrator := getMigrator()
			return migrator.Init(c.Context)
		},
	},
	{
		Name:  "migrate",
		Usage: "migrate database",
		Action: func(c *cli.Context) error {
			migrator := getMigrator()

			group, err := migrator.Migrate(c.Context)
			if err != nil {
				return err
			}

			if group.ID == 0 {
				fmt.Printf("there are no new migrations to run\n")
				return nil
			}

			fmt.Printf("migrated to %s\n", group)
			return nil
		},
	},
	{
		Name:  "rollback",
		Usage: "rollback the last migration group",
		Action: func(c *cli.Context) error {
			migrator := getMigrator()

			group, err := migrator.Rollback(c.Context)
			if err != nil {
				return err
			}

			if group.ID == 0 {
				fmt.Printf("there are no groups to roll back\n")
				return nil
			}

			fmt.Printf("rolled back %s\n", group)
			return nil
		},
	}, {
		Name:  "status",
		Usage: "print migrations status",
		Action: func(c *cli.Context) error {
			migrator := getMigrator()

			ms, err := migrator.MigrationsWithStatus(c.Context)
			if err != nil {
				return err
			}
			fmt.Printf("migrations: %s\n", ms)
			fmt.Printf("unapplied migrations: %s\n", ms.Unapplied())
			fmt.Printf("last migration group: %s\n", ms.LastGroup())

			return nil
		},
	},
}

func getMigrator() *migrate.Migrator {
	db := database.DB()

	return migrate.NewMigrator(db, migrations.Migrations)
}
