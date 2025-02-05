package routes

import (
	controller "carbon/go-commerce/controllers"
	"carbon/go-commerce/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.Use(middleware.Authenticate())
	router.POST("/users", controller.GetUsers())
	router.POST("/user/:user_id", controller.GetUser())
}
