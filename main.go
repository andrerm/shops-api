package main

import (
	"ShopsAPI/config"
	"ShopsAPI/middleware"
	"ShopsAPI/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(middleware.ResponseMiddleware())
	r.Use(middleware.PaginationMiddleware(1, 10)) // Default page 1, limit 10

	config.ConnectDatabase()
	routes.SetupRoutes(r)
	r.Run(":8080")
}
