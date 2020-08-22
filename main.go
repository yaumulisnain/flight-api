package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

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
