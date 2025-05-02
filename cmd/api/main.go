package main

import (
	"healthcare-appointment-system/internal/database"
	"healthcare-appointment-system/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Connect()

	r := gin.Default()

	handlers.SetupRoutes(r)

	r.Run()
}
