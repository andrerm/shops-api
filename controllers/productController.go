package controllers

import (
	"ShopsAPI/config"
	"ShopsAPI/middleware"
	"ShopsAPI/models"
	"ShopsAPI/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct creates a new product
func CreateProduct(c *gin.Context) {
	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	product := models.Product{
		Name:      input.Name,
		Price:     input.Price,
		Stock:     input.Stock,
		StoreID:   input.StoreID,
		CreatedBy: input.CreatedBy,
		UpdatedBy: input.UpdatedBy,
	}
	config.DB.Create(&product)
	utils.RespondSuccess(c, product, nil)
}

// GetProducts retrieves all products with pagination
func GetProducts(c *gin.Context) {
	var products []models.Product
	pagination := c.MustGet("pagination").(middleware.Pagination)

	var total int64
	config.DB.Model(&models.Product{}).Count(&total)

	offset := (pagination.Page - 1) * pagination.Limit
	config.DB.Offset(offset).Limit(pagination.Limit).Find(&products)

	utils.RespondSuccess(c, products, gin.H{
		"page":  pagination.Page,
		"limit": pagination.Limit,
		"total": total,
	})
}

// GetProduct retrieves a single product by ID
func GetProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.Where("product_id = ?", c.Param("id")).First(&product).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Product not found", err.Error())
		return
	}
	utils.RespondSuccess(c, product, nil)
}

// UpdateProduct updates an existing product
func UpdateProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.Where("product_id = ?", c.Param("id")).First(&product).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Product not found", err.Error())
		return
	}

	var input models.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid request data", err.Error())
		return
	}

	config.DB.Model(&product).Updates(input)
	utils.RespondSuccess(c, product, nil)
}

// DeleteProduct deletes an existing product
func DeleteProduct(c *gin.Context) {
	var product models.Product
	if err := config.DB.Where("product_id = ?", c.Param("id")).First(&product).Error; err != nil {
		utils.RespondError(c, http.StatusNotFound, "Product not found", err.Error())
		return
	}

	config.DB.Delete(&product)
	utils.RespondSuccess(c, true, nil)
}
