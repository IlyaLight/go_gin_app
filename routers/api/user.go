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
