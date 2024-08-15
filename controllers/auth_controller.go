package controllers

import (
	"ShopsAPI/config"
	"ShopsAPI/middleware"
	"ShopsAPI/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		AppType  string `json:"app_type" binding:"required"`
	}

	// Validate request payload
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Find the user by email
	var user models.User
	if err := config.DB.Where("email = ?", credentials.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Fetch the user's role from the user_roles table
	var userRole models.UserRole
	if err := config.DB.Preload("Role").Where("user_id = ?", user.UserID).First(&userRole).Error; err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "User role not found"})
		return
	}

	log.Println("User ID:", user.UserID)
	log.Println("User Role Query:", userRole)
	log.Println("Fetched Role:", userRole.Role.RoleName)

	var role models.Role
	if err := config.DB.Where("role_id = ?", userRole.RoleID).First(&role).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Role"})
		return
	}
	// Generate JWT token with role included
	// token, err := middleware.GenerateToken(user.Email, role.RoleName, credentials.AppType)
	token, err := middleware.GenerateToken(user.Email, userRole.Role.RoleName, credentials.AppType)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// Respond with the generated token
	c.JSON(http.StatusOK, gin.H{"token": token})
}
