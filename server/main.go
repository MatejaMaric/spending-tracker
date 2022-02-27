package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/MatejaMaric/spending-tracker/server/models"
	"github.com/MatejaMaric/spending-tracker/server/routes"
)

func main() {
	server := gin.Default()

	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "User-Agent", "Referrer", "Host", "Token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           86400,
	}
	server.Use(cors.New(corsConfig))

	models.Connect()
	routes.Setup(server)

	server.Run("0.0.0.0:3000")
}
