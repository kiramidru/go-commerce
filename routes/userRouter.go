package routes

import (
	controller "carbon/go-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	router.POST("/users", controller.GetUsers())
	router.POST("/user/:user_id", controller.GetUser())
}
