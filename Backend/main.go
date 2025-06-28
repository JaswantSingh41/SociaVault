package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type DownloadRequest struct {
	URL string `json:"url" binding:"required"`
}

type DownloadResponse struct {
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Platform string      `json:"platform,omitempty"`
	Title    string      `json:"title,omitempty"`
	Uploader string      `json:"uploader,omitempty"`
	Type     string      `json:"type,omitempty"`
	JobID    string      `json:"job_id,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type UniversalDownloader struct {
	DownloadDir string
}

func NewUniversalDownload() *UniversalDownloader {
	downloadDir := "./downloads"
	if err := os.MkdirAll(downloadDir, 0755); err != nil {
		log.Printf("Failed to create download directory:%v", err)
	}
	return &UniversalDownloader{
		DownloadDir: downloadDir,
	}
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	// ud := NewUniversalDownload()
	router.POST("/api/download", func(c *gin.Context) {
		var req DownloadRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, DownloadResponse{Status: "error", Message: err.Error()})
		}
	})
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to SociaVault backend!",
		})
	})

	router.Run("localhost:8080")
}
