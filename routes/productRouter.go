package routes

import (
	"carbon/go-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/products", controllers.GetProducts())
	router.GET("/product/:product_id", controllers.GetProduct())
}
