package routes

import (
	"carbon/go-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.GET("/login", controllers.Login())
	router.GET("/signup", controllers.Signup())
}
