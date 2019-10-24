package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"os"
)

type Database struct {
	User     string
	Password string
	Host     string
	Name     string
	Port     string
}

var DatabaseSetting = &Database{}

// Setup initialize the configuration instance
func Setup() {
	if os.Getenv("ENVIRONMENT") == "DEV" {
		godotenv.Load() //Загрузить файл .env
	}

	DatabaseSetting.User = os.Getenv("db_user")
	DatabaseSetting.Password = os.Getenv("db_pass")
	DatabaseSetting.Host = os.Getenv("db_host")
	DatabaseSetting.Name = os.Getenv("db_name")
	DatabaseSetting.Port = os.Getenv("db_port")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

}
