package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go_gin_app/config"
	migration "go_gin_app/db"
	"log"
)

const dialect = "postgres"

var db *gorm.DB
var name int

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

// Setup initializes the database instance
func Setup() {

	dbArgs := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DatabaseSetting.User,
		config.DatabaseSetting.Password,
		config.DatabaseSetting.Host,
		config.DatabaseSetting.Port,
		config.DatabaseSetting.Name)

	var err error
	db, err = gorm.Open(dialect, dbArgs)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	db.LogMode(true)

	migration.Migration(dbArgs)

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

func GetDB() *gorm.DB {
	return db
}
