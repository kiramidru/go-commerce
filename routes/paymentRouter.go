package routes

import (
	"carbon/go-commerce/controllers"

	"github.com/gin-gonic/gin"
)

func PaymentRoutes(router *gin.Engine) {
	router.GET("/accept_payment", controllers.AcceptPayment())
	router.GET("/verify_transaction", controllers.VerifyTransaction())
	router.GET("/get_transactions", controllers.GetTransactions())
}
