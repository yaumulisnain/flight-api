package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var (
		err     error
		logFile *os.File
		logPath string
	)

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	logPath, err = os.Getwd()

	if err != nil {
		log.Fatal("Error getting working directory")
		os.Exit(1)
	}

	fileName := path.Join(logPath, "logs/access.log")

	logFile, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Fatal("Failed to write/create log file")
		os.Exit(1)
	}

	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		r := fmt.Sprintf("%s%s Not Found", c.Request.Host, c.Request.URL)
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "message": r, "success": false})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"apps":    "Flight API Service",
			"version": "1.0.0",
		})
	})

	port := os.Getenv("PORT")
	router.Run(port)
}
