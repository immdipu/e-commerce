package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/immdipu/e-commerce-go/controllers"
)

func UserRoutes(router *gin.Engine) {
	router.POST("users/signup", controllers.Signup)
	// router.POST("/users/login", controllers.Login())
	// router.POST("/admin/addproduct", controllers.ProductViewerAdmin())
	// router.GET("/users/productview", controllers.SearchProduct())
	// router.GET("/users/search", controllers.SearchProductByQuery())
}
