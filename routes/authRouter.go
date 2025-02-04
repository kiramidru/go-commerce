package routes

import (
	controller "carbon/go-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/login", controller.Login())
	router.GET("/signup", controller.Signup())
}
