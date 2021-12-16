package controllers

import (
	"golang-products-api/database"
	"golang-products-api/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListProducts(context *gin.Context) {
	database := database.GetDatabase()

	var products []models.Product
	errorFromProductsList := database.Find(&products).Error

	if errorFromProductsList != nil {
		context.JSON(400, gin.H{
			"error": "Failed while attempting to list all products: " + errorFromProductsList.Error(),
		})

		return
	}

	
	context.JSON(200, products)
}

func GetProductById(context *gin.Context) {
	paramId := context.Param("id")
	productId, errorFromIdParse := strconv.Atoi(paramId)

	if errorFromIdParse != nil {
		context.JSON(400, gin.H{
			"error": "Product ID needs to be an integer",
		})

		return
	}

	database := database.GetDatabase()

	var product models.Product
	errorFromProduct := database.First(&product, productId).Error

	if errorFromProduct != nil {
		context.JSON(400, gin.H{
			"error": "Product not found: " + errorFromProduct.Error(),
		})

		return
	}

	context.JSON(200, product)
}

func CreateProduct(context *gin.Context) {
	database := database.GetDatabase()

	var product models.Product

	errorFromJSONBind := context.ShouldBindJSON(&product)

	if errorFromJSONBind != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind JSON object: " + errorFromJSONBind.Error(),
		})

		return
	}

	errorFromProduct := database.Create(&product).Error

	if errorFromProduct != nil {
		context.JSON(400, gin.H{
			"error": "Create product failed: " + errorFromProduct.Error(),
		})

		return
	}

	context.JSON(200, product)
}

func UpdateProduct(context *gin.Context) {
	database := database.GetDatabase()

	var product models.Product

	errorFromJSONBind := context.ShouldBindJSON(&product)

	if errorFromJSONBind != nil {
		context.JSON(400, gin.H{
			"error": "Cannot bind JSON object: " + errorFromJSONBind.Error(),
		})

		return
	}

	errorFromProduct := database.Save(&product).Error

	if errorFromProduct != nil {
		context.JSON(400, gin.H{
			"error": "Update product failed: " + errorFromProduct.Error(),
		})

		return
	}

	context.JSON(200, product)
}

func DeleteProduct(context *gin.Context) {
	paramId := context.Param("id")
	productId, errorFromIdParse := strconv.Atoi(paramId)

	if errorFromIdParse != nil {
		context.JSON(400, gin.H{
			"error": "Product ID needs to be an integer",
		})

		return
	}

	database := database.GetDatabase()

	errorFromDelete := database.Delete(&models.Product{}, productId).Error

	if errorFromDelete != nil {
		context.JSON(400, gin.H{
			"error": "Delete product failed: " + errorFromDelete.Error(),
		})

		return
	}

	context.Status(200)
}