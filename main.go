package main

import (
	"carbon/go-commerce/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.PaymentRoutes(router)
	routes.ProductRoutes(router)
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.Run(":" + port)
}
