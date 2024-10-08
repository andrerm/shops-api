package controllers

import (
	"ShopsAPI/config"
	"ShopsAPI/middleware"
	"ShopsAPI/models"
	"ShopsAPI/utils"
	"net/http"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes the given password using bcrypt.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPasswordHash compares a hashed password with its plain-text version.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// GetUsers retrieves all users with pagination
func GetUsers(c *gin.Context) {
	var users []models.User
	pagination := c.MustGet("pagination").(middleware.Pagination)

	var total int64
	config.DB.Model(&models.User{}).Count(&total)

	offset := (pagination.Page - 1) * pagination.Limit
	config.DB.Offset(offset).Limit(pagination.Limit).Find(&users)

	utils.RespondSuccess(c, users, gin.H{
		"page":  pagination.Page,
		"limit": pagination.Limit,
		"total": total,
	})
}

// // CreateUser creates a new user
// func CreateUser(c *gin.Context) {
// 	var input struct {
// 		Name     string `json:"name" binding:"required"`
// 		Email    string `json:"email" binding:"required"`
// 		Password string `json:"password" binding:"required"`
// 	}
// 	if err := c.ShouldBindJSON(&input); err != nil {
// 		utils.RespondError(c, http.StatusBadRequest, "Invalid request data", err.Error())
// 		return
// 	}

// 	// Hash the password
// 	hashedPassword, err := HashPassword(input.Password)
// 	if err != nil {
// 		utils.RespondError(c, http.StatusInternalServerError, "Failed to hash password", err.Error())
// 		return
// 	}

// 	user := models.User{
// 		UserID:   uuid.New(), // Generate new UUID for UserID
// 		Name:     input.Name,
// 		Email:    input.Email,
// 		Password: hashedPassword,
// 	}

// 	// Log the user struct
// 	log.Println("Creating user:", user)

// 	if err := config.DB.Create(&user).Error; err != nil {
// 		log.Println("Error creating user:", err)
// 		utils.RespondError(c, http.StatusInternalServerError, "Failed to create user", err.Error())
// 		return
// 	}
// 	utils.RespondSuccess(c, user, nil)
// }

func CreateUser(c *gin.Context) {
	var input struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		RoleName string `json:"role_name" binding:"required"`
		StoreID  string `json:"store_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to hash password", err.Error())
		return
	}

	user := models.User{
		UserID:   uuid.New(),
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create user", err.Error())
		return
	}

	var role models.Role
	if err := config.DB.Where("role_name = ?", input.RoleName).First(&role).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to find role", err.Error())
		return
	}

	userRole := models.UserRole{
		UserRoleID: uuid.New(),
		UserID:     user.UserID,
		RoleID:     role.RoleID,
		StoreID:    uuid.MustParse(input.StoreID),
	}

	if err := config.DB.Create(&userRole).Error; err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to assign role to user", err.Error())
		return
	}

	utils.RespondSuccess(c, user, nil)
}

// GetUser retrieves a single user by ID
func GetUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("user_id = ?", c.Param("id")).First(&user).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found", err.Error())
		return
	}
	utils.RespondSuccess(c, user, nil)
}

// UpdateUser updates an existing user
func UpdateUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("user_id = ?", c.Param("id")).First(&user).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found", err.Error())
		return
	}

	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	// If password is being updated, hash it
	if input.Password != "" {
		hashedPassword, err := HashPassword(input.Password)
		if err != nil {
			utils.RespondError(c, http.StatusInternalServerError, "Failed to hash password", err.Error())
			return
		}
		input.Password = hashedPassword
	}

	config.DB.Model(&user).Updates(input)
	utils.RespondSuccess(c, user, nil)
}

// DeleteUser deletes an existing user
func DeleteUser(c *gin.Context) {
	var user models.User
	if err := config.DB.Where("user_id = ?", c.Param("id")).First(&user).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "User not found", err.Error())
		return
	}

	config.DB.Delete(&user)
	utils.RespondSuccess(c, true, nil)
}
