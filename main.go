package main

import (
	"github.com/sirupsen/logrus"
	"go_gin_app/config"
	"go_gin_app/models"
	"go_gin_app/routers"
)

func init() {
	config.Setup()
	models.Setup()
}

//var router *gin.Engine

func main() {

	// Initialize the routes
	router := routers.InitRouter()

	router.LoadHTMLGlob("templates/*")

	logrus.Info("run router")
	// Start serving the application
	router.Run()

	logrus.Info("Hello, world!")

}
