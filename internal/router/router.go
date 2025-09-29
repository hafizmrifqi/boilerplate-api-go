package router

import (
	"boilerplate-api-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "OK", "message": "API is running"})
	})

	// Contoh route placeholder — ganti sesuai kebutuhan proyek
	api := r.Group("/api")
	{
		// Contoh: api.GET("/users", handlers.GetUsers)
		// Tambahkan route di sini saat mulai proyek baru
	}

	return r
}