package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"

	"github.com/amacneil/dbmate/pkg/dbmate"
	_ "github.com/amacneil/dbmate/pkg/driver/postgres"
	"github.com/mochammadshenna/saza-aluminium/util/helper"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "dbmate"
	app.Usage = "A lightweight, framework-independent database migration tool."
	app.Version = dbmate.Version

	app.Commands = []*cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "Generate a new migration file",
			Action: action(func(db *dbmate.DB, c *cli.Context) error {
				name := c.Args().First()
				return db.NewMigration(name)
			}),
		},
		{
			Name:  "up",
			Usage: "Create database (if necessary) and migrate to the latest version",
			Flags: []cli.Flag{
				&cli.BoolFlag{
					Name:    "verbose",
					Aliases: []string{"v"},
					EnvVars: []string{"DBMATE_VERBOSE"},
					Usage:   "print the result of each statement execution",
				},
			},
			Action: action(func(db *dbmate.DB, c *cli.Context) error {
				db.Verbose = c.Bool("verbose")
				return db.CreateAndMigrate()
			}),
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		re := regexp.MustCompile("([a-zA-Z]+://[^:]+:)[^@]+@")
		errText := re.ReplaceAllString(fmt.Sprintf("Error: %s\n", err), "${1}********@")

		_, _ = fmt.Fprint(os.Stderr, errText)
		os.Exit(2)
	}
}

func action(f func(*dbmate.DB, *cli.Context) error) cli.ActionFunc {
	return func(c *cli.Context) error {
		dbConfig := "root"
		dbPassword := "root"
		dbHost := "localhost"
		dbPort := 5432
		dbName := "sayakaya_db"

		link := fmt.Sprintf(
			"%s://%s:%s@%s:%d/%s?sslmode=disable",
			"postgres",
			dbConfig,
			dbPassword,
			dbHost,
			dbPort,
			dbName,
		)

		u, err := url.Parse(link)
		helper.PanicError(err)

		db := dbmate.New(u)
		db.AutoDumpSchema = !c.Bool("no-dump-schema")
		db.SchemaFile = c.String("schema-file")
		db.MigrationsDir = "./scripts/migrations"

		return f(db, c)
	}
}
