package api

import (
	"github.com/gin-gonic/gin"
	"go_gin_app/models"
	"go_gin_app/service"
	"net/http"
)

func PostUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.AddUser(&newUser)
}

func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	if user, _ := service.VerifyUser(username, password); user != nil {
		c.JSON(200, user)
	}

}

func GetUSER(c *gin.Context) {

}
