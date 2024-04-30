package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/immdipu/e-commerce-go/database"
	"github.com/immdipu/e-commerce-go/routes"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	database.DBSet()
	defer database.CloseDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.Default()
	routes.UserRoutes(router)
	// router.Use(middleware.Authentication())
	log.Fatal(router.Run(":" + port))
}
