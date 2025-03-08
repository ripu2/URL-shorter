package middleware

import (
	"example.com/url-shorter/internal/services"
	"github.com/gin-gonic/gin"
)

// Redis cache middleware
func CacheMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		cacheKey := ctx.Param("shortenedURL")
		// Check Redis cache
		if cachedData, err := services.GetValueFromCache(cacheKey); err == nil && cachedData != "" {
			ctx.Abort() // Return cached response
			return
		}

		// If not cached, continue processing the request
		// Capture response and cache it
		ctx.Writer = &cacheWriter{
			ResponseWriter: ctx.Writer,
			key:            cacheKey,
		}
		ctx.Next()
	}
}

// Custom response writer to cache response
type cacheWriter struct {
	gin.ResponseWriter
	key string
}

func (cw *cacheWriter) Write(data []byte) (int, error) {
	_ = services.SetValueInCache(cw.key, string(data)) // Store response in cache
	return cw.ResponseWriter.Write(data)
}
