package routes

import (
	"ShopsAPI/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	// User routes
	r.GET("/users", controllers.GetUsers)
	r.POST("/users", controllers.CreateUser)

	// Product routes
	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)
	r.GET("/products/:id", controllers.GetProduct)
	r.PUT("/products/:id", controllers.UpdateProduct)
	r.DELETE("/products/:id", controllers.DeleteProduct)
	// Repeat for other resources
}
