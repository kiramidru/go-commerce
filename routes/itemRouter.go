package routes

import (
	controller "carbon/go-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func ItemRoutes(router *gin.Engine) {
	router.GET("/items", controller.GetItems())
	router.GET("/item/:item_id", controller.GetItem())
}
