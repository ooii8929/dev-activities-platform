package routes

import (
	"dev-platform/controllers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/", controllers.IndexHome)
	router.GET("/index", controllers.IndexHome)
	router.GET("/login", controllers.IndexLogin)

	router.GET("/users/:id", controllers.UserGet)
	router.GET("/users", controllers.UserGetList)
	router.POST("/login", controllers.UserPost)
	router.PUT("/users/:id", controllers.UserPut)
	router.PATCH("/users/:id", controllers.UserPut)
	router.DELETE("/users/:id", controllers.UserDelete)
}