package config

import (
	"fmt"
	"os"

	_ "github.com/golang-migrate/migrate/database/postgres"
	"github.com/trip/trip-service/internal/db"
)

func databaseConfig() error {
	pgUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		os.Getenv("PG_USER"),
		os.Getenv("PG_PASSWORD"),
		os.Getenv("PG_HOST"),
		os.Getenv("PG_PORT"),
		os.Getenv("PG_DB_NAME"),
		os.Getenv("PG_SSL_MODE"),
	)

	err := db.ConnectToPostgres(pgUrl)

	return err
}
