package main

import (
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

func addRoutes(router *gin.Engine) {

	router.Use(static.Serve("/", static.LocalFile("./views", true)))

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "test. What did you think this would be? ;)"})
		})

		// api.GET("/{urlShortCode}", redirect)

		api.POST("/shortenURL", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Pong"})
		})

	}

}

type URLRequestObject struct {
	URL string `json:"url"`
}

type URLCollection struct {
	ActualURL string

	ShortURL string
}
type SuccessResponse struct {
	Code     int
	Message  string
	Response URLCollection
}
