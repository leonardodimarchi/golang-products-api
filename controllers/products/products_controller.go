package controllers

import "github.com/gin-gonic/gin"

func ListProducts(context *gin.Context) {
	context.JSON(200, gin.H{
		"value": "ok",
	})
}