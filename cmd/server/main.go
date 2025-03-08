package main

import (
	"fmt"
	"os"

	db "example.com/url-shorter/config/db"
	redis "example.com/url-shorter/config/redis"
	"example.com/url-shorter/internal/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	server := gin.Default()
	routes.SetupRoutes(server)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	db.InitDB()
	redis.InitRedisClient()

	defer func() {
		db.DB.Close()
		redis.RedisClient.Close()
	}()

	defer db.DB.Close()
	fmt.Println("🚀 Server is up an running on http://localhost:8080")
	server.Run(":" + port) //localhost:8080
}
