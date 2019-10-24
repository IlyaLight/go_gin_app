package routers

import (
	"github.com/gin-gonic/gin"
	"go_gin_app/middleware"
	"go_gin_app/routers/api"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/user", api.PostUser)

	// Handle GET requests at /article/view/some_article_id
	r.GET("/article/view/:article_id", api.GetArticle)

	authorized := r.Group("/api")
	authorized.Use(middleware.Authentication())
	{
		//authorized.POST("/login", loginEndpoint)
		//authorized.POST("/submit", submitEndpoint)
		//authorized.POST("/read", readEndpoint)
		//
		//// nested group
		//testing := authorized.Group("testing")
		//testing.GET("/analytics", analyticsEndpoint)
	}

	return r
}
