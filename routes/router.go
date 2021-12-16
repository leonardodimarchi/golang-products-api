package routes

import (
	controllers "golang-products-api/controllers/products"

	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1") 
	{
		products := main.Group("products") 
		{
			products.GET("/", controllers.ListProducts)
			products.GET("/:id", controllers.GetProductById)
			products.POST("/", controllers.CreateProduct)
			products.PUT("/", controllers.UpdateProduct)
			products.DELETE("/:id", controllers.DeleteProduct)
		}
	}

	return router
}