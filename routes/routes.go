package routes

import (
	"ShopsAPI/controllers"
	"ShopsAPI/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// Public routes
	r.POST("/login", controllers.Login)

	// User routes
	userRoutes := r.Group("/users")
	// userRoutes.Use(middleware.Authenticate())
	userRoutes.GET("/", controllers.GetUsers)
	userRoutes.POST("/", controllers.CreateUser)
	userRoutes.GET("/:id", controllers.GetUser)
	userRoutes.PUT("/:id", controllers.UpdateUser)
	userRoutes.DELETE("/:id", controllers.DeleteUser)

	// Product routes
	productRoutes := r.Group("/products")
	productRoutes.Use(middleware.Authenticate())
	productRoutes.GET("/", controllers.GetProducts)
	productRoutes.POST("/", middleware.Authorize("admin", "mobile"), controllers.CreateProduct)
	productRoutes.GET("/:id", controllers.GetProduct)
	productRoutes.PUT("/:id", middleware.Authorize("admin", "mobile"), controllers.UpdateProduct)
	productRoutes.DELETE("/:id", middleware.Authorize("admin", "mobile"), controllers.DeleteProduct)
}
