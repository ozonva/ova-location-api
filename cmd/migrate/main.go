package main

import (
	"flag"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/ozonva/ova-location-api/config"
	"github.com/rs/zerolog/log"
	"os"

	"github.com/pressly/goose/v3"
)

var (
	flags      = flag.NewFlagSet("goose", flag.ExitOnError)
	dir        = flags.String("dir", "migrations", "directory with migration files")
	table      = flags.String("table", "migrations", "migrations table name")
	verbose    = flags.Bool("v", false, "enable verbose mode")
	help       = flags.Bool("h", false, "print help")
	version    = flags.Bool("version", false, "print version")
	sequential = flags.Bool("s", false, "use sequential numbering for new migrations")
)

func main() {
	flags.Usage = usage
	flags.Parse(os.Args[1:])

	if *version {
		fmt.Println(goose.VERSION)
		return
	}
	if *verbose {
		goose.SetVerbose(true)
	}
	if *sequential {
		goose.SetSequential(true)
	}
	goose.SetTableName(*table)

	args := flags.Args()
	if len(args) == 0 || *help {
		flags.Usage()
		return
	}

	command := args[0]
	switch command {
	case "create":
		if err := goose.Run("create", nil, *dir, args[1:]...); err != nil {
			log.Fatal().Err(err).Msg("Не удалось создать миграцию")
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatal().Err(err).Msg("Не удалось исправить порядок")
		}
		return
	}

	cfg := config.Get()
	db, err := sqlx.Connect("pgx", cfg.Db.GetDsn())
	if err != nil {
		log.Fatal().Err(err).Msg("Не удалось подключиться к БД")
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal().Err(err).Msg("Не удалось закрыть подключение к БД")
		}
	}()

	if err := goose.SetDialect(cfg.Db.Dialect); err != nil {
		log.Fatal().Err(err).Msg("Неизвестный язык запросов")
	}

	if err := goose.Run(command, db.DB, *dir, args[1:]...); err != nil {
		log.Fatal().Err(err).Msg("Не удалось выполнить команду")
	}
}

func usage() {
	fmt.Println(usagePrefix)
	flags.PrintDefaults()
	fmt.Println(usageCommands)
}

var (
	usagePrefix = `Usage: migrate [OPTIONS] COMMAND

Examples:
    migrate status
    migrate create init sql
    migrate create add_some_column sql
    migrate create fetch_user_data go
    migrate up

Options:`

	usageCommands = `
Commands:
    up                   Migrate the DB to the most recent version available
    up-by-one            Migrate the DB up by 1
    up-to VERSION        Migrate the DB to a specific VERSION
    down                 Roll back the version by 1
    down-to VERSION      Roll back to a specific VERSION
    redo                 Re-run the latest migration
    reset                Roll back all migrations
    status               Dump the migration status for the current DB
    version              Print the current version of the database
    create NAME [sql|go] Creates new migration file with the current timestamp
    fix                  Apply sequential ordering to migrations`
)
