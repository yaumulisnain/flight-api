package main

import (
	"flight-api/app/core"
	v1 "flight-api/app/v1"
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

	core.InitDB()
	defer core.GetDB().Close()

	router := gin.Default()

	router.NoRoute(func(c *gin.Context) {
		r := fmt.Sprintf("%s%s Not Found", c.Request.Host, c.Request.URL)
		c.JSON(http.StatusNotFound, gin.H{"code": http.StatusNotFound, "error": r, "message": http.StatusText(http.StatusNotFound)})
	})

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"apps":    "Flight API Service",
			"version": "1.0.0",
		})
	})

	v1.Route(router)

	port := os.Getenv("PORT")
	router.Run(port)
}
