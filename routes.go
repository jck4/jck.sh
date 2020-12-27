package main

import (
	"github.com/gin-gonic/gin"
	controllers "github.com/jck4/jck.sh/controllers"
	"github.com/gin-gonic/contrib/static"
	github "github.com/jck4/jck.sh/middleware"
)

func addRoutes(r *gin.Engine) {

	r.Use(static.Serve("/", static.LocalFile("./views", true)))


	r.GET("/url", controllers.FindURLs)
	r.POST("/url", controllers.CreateURL)


	r.GET("/login", github.LoginHandler)
    r.GET("/auth", github.Auth)

}
