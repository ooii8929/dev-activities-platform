package main

import (
	"dev-platform/globals"
	"dev-platform/routes"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)



func main() {

    router := gin.Default()

		router.Static("/static", "./static")

		router.Use(sessions.Sessions("session", cookie.NewStore(globals.Secret)))

    public := router.Group("/")
		routes.PublicRoutes(public)
    
    router.Run(":9090")

}
