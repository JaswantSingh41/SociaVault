package main

import (
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

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

// method to find the platform form url
func (ud *UniversalDownloader) DetectPlatform(url string) string {
	url = strings.ToLower(url)
	switch {
	case strings.Contains(url, "youtube.com") || strings.Contains(url, "youtu.be"):
		return "youtube"
	case strings.Contains(url, "instagram.com"):
		return "instagram"
	case strings.Contains(url, "facebook.com") || strings.Contains(url, "fb.watch"):
		return "facebook"
	case strings.Contains(url, "twitter.com") || strings.Contains(url, "x.com"):
		return "twitter"
	case strings.Contains(url, "tiktok.com"):
		return "tiktok"
	case strings.Contains(url, "reddit.com"):
		return "reddit"
	case strings.Contains(url, "twitch.tv"):
		return "twitch"
	case strings.Contains(url, "vimeo.com"):
		return "vimeo"
	default:
		return "unknown"
	}
}

// method for create the file name of downloaded content
func (ud *UniversalDownloader) CreateFilename(filename string, maxLength int) string {
	re := regexp.MustCompile(`[<>:"/\\|?*]`)
	filename = re.ReplaceAllString(filename, "_")
	filename = strings.TrimSpace(filename)
	if len(filename) > maxLength {
		filename = filename[:maxLength]
	}
	return filename
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	// ud := NewUniversalDownload()
	router.POST("/api/download", func(c *gin.Context) {
		var req DownloadRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, DownloadResponse{Status: "error", Message: err.Error()})
			return
		}
	})
	router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to SociaVault backend!",
		})
	})

	router.Run("localhost:8080")
}
