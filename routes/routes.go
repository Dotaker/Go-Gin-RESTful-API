package routes

import (
	"net/http"

	article "CRUD-Operation/controllers/article"
	user "CRUD-Operation/controllers/user"

	"github.com/gin-gonic/gin"
)

// StartGin function
func StartGin() *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")
	{
		// Users Routes
		api.GET("/users", user.GetAllUser)
		api.POST("/users", user.CreateUser)
		api.GET("/users/:id", user.GetUser)
		api.PUT("/users/:id", user.UpdateUser)
		api.DELETE("/users/:id", user.DeleteUser)

		// Articles Routes
		api.GET("/articles", article.GetAllArticle)
		api.POST("/articles", article.CreateArticle)
		api.GET("/articles/:id", article.GetArticle)
		api.PUT("/articles/:id", article.UpdateArticle)
		api.DELETE("/articles/:id", article.DeleteArticle)

	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})

	return router
}
