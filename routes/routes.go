package routes

import (
	"github.com/gin-gonic/gin"

	controllers "dev-platform/controllers"
)

func PublicRoutes(g *gin.RouterGroup) {

	// g.GET("/login", controllers.LoginGetHandler())
	g.POST("/login", controllers.LoginPostHandler())
	// g.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// g.GET("/", controllers.IndexGetHandler())


}
