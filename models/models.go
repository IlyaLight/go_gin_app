package models

import (
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"go_gin_app/config"
	migration "go_gin_app/db"
	//"log"
)

const dialect = "postgres"

var db *sqlx.DB
var name int

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
}

// SetupConnection initializes the database instance
func SetupConnection() {

	dbArgs := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.DatabaseSetting.User,
		config.DatabaseSetting.Password,
		config.DatabaseSetting.Host,
		config.DatabaseSetting.Port,
		config.DatabaseSetting.Name)
	db = sqlx.MustConnect("pgx", dbArgs)
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)

	migration.Migration(dbArgs)
	//addTestDataToDb()

	//var err error
	//db, err = gorm.Open(dialect, dbArgs)
	//if err != nil {
	//	log.Fatalf("models.SetupConnection err: %v", err)
	//}
	//db.LogMode(true)
	//db.DB().SetMaxIdleConns(10)
	//db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection (unnecessary)
func CloseDB() {
	defer db.Close()
}

func GetDB() *sqlx.DB {
	return db
}
