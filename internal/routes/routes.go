package routes

import (
	"example.com/url-shorter/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(server *gin.Engine) {
	server.POST("/", handlers.ShortenURLHandler)
	server.GET("/:shortenedURL", handlers.RedirectURLHandler)
}
