package middleware

import (
	"ShopsAPI/config"
	"ShopsAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeRole(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		storeID := c.Query("store_id") // Assuming store_id is passed as a query parameter
		if !exists || storeID == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated or store not specified"})
			c.Abort()
			return
		}

		var userRole models.UserRole
		if err := config.DB.Joins("Role").Where("user_id = ? AND store_id = ?", userID, storeID).First(&userRole).Error; err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "User does not have the required role"})
			c.Abort()
			return
		}

		if userRole.Role.RoleName != requiredRole {
			c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}
