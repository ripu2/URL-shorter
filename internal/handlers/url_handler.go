package handlers

import (
	"net/http"

	"example.com/url-shorter/internal/models"
	"example.com/url-shorter/internal/services"
	"github.com/gin-gonic/gin"
)

func ShortenURLHandler(ctx *gin.Context) {
	var url models.URL
	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	val, err := services.SortUrl(&url)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"message": "URL shortened successfully babe", "link": "http://localhost:8080/" + val})
}

func RedirectURLHandler(ctx *gin.Context) {
	shortenedURL := ctx.Param("shortenedURL")
	url, err := services.GetRedirectURL(shortenedURL)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.Redirect(http.StatusFound, url)

}
