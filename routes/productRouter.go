package routes

import (
	controller "carbon/go-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(router *gin.Engine) {
	router.GET("/items", controller.GetProducts())
	router.GET("/item/:item_id", controller.GetProduct())
}
