package config

import (
	"fmt"

	"github.com/trip/trip-service/internal/db"
)

func Config() error {
	err := databaseConfig()
	if err != nil {
		return fmt.Errorf("Config-> %s", err.Error())
	}

	return nil
}

func ConfigDefer() {
	db.Close()
}
