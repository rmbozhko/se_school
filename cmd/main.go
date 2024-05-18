package main

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/robfig/cron"
	"log"
	"os"
	"se_school/api"
	db "se_school/db/sqlc"
	"se_school/util"
	"strconv"
)

// @title GSES BTC application
// @version 1.0.0
// @Tags rate "Отримання поточного курсу USD до UAH"
// @host gses2.app:8080
// @basePath /api
// @schemes http
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("failed to connect to database")
	}

	runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(conn)
	smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Fatal("SMTP port should be an integer")
	}

	dialer := util.NewEmailDialer(os.Getenv("SMTP_HOST"), smtpPort, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))
	server := api.NewServer(store, dialer)

	job := cron.New()
	err = job.AddFunc("4 * * * *", server.SendRatesEmail)
	if err != nil {
		return
	}

	job.Start()

	if err = server.Start(config.ServerAddress); err != nil {
		log.Fatal("Cannot start the server at the address {}:\n{}", config.ServerAddress, err)
	}
}

func runDBMigration(migrationURL string, dbSource string) {
	m, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}
	version, _, err := m.Version()
	if err != nil && !errors.Is(err, migrate.ErrNilVersion) {
		log.Fatal("failed to fetch migration version:", err)
	}

	if errors.Is(err, migrate.ErrNilVersion) {
		if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal("failed to run migrate up:", err)
		}
		log.Println("db migrated successfully")
	} else {
		log.Printf("db already migrated to version: %d\n", version)
	}
}
