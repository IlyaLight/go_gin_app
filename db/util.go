package db

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/viper"
	"log"
)

func Migration(dbUrl string) {
	m, err := migrate.New(
		viper.GetString("migration.path"),
		dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			log.Print("migrate up error: no change")
			return
		}
		log.Fatal(err)
	}
}
