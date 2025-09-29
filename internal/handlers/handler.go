package handlers

import (
	"boilerplate-api-go/pkg/response"

	"github.com/gin-gonic/gin"
)

// Contoh handler — hapus atau ganti saat mulai proyek baru
func GetExample(c *gin.Context) {
	response.Success(c, gin.H{
		"message": "This is a placeholder. Replace with your logic!",
	})
}