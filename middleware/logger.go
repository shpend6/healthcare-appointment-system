package middleware

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var logger *log.Logger

func init() {
	// Create log directory if not exists
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		log.Fatalf("could not create log directory: %v", err)
	}

	// Open log file
	file, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("failed to open log file: %v", err)
	}

	// MultiWriter to log to file and stdout
	logger = log.New(file, "", log.LstdFlags)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		duration := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		path := c.Request.URL.Path

		logger.Printf("[HTTP] %s %s %d %s", method, path, status, duration)
	}
}
