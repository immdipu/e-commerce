package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/immdipu/e-commerce-go/middleware"
	"github.com/immdipu/e-commerce-go/routes"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	log.Fatal(router.Run(":" + port))
}
