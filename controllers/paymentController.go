package controllers

import (
	"carbon/go-commerce/database"
	"carbon/go-commerce/models"
	"context"
	"net/http"
	"time"

	"github.com/Chapa-Et/chapa-go"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/mongo"
)

var PaymentCollection *mongo.Collection = database.OpenCollection(*database.Client, "payment")

func AcceptPayment() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payment models.Payment
		c.BindJSON(&payment)

		amount, err := decimal.NewFromString(payment.Amount)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}

		chapa.InitConfig()
		var chapaAPI = chapa.New()

		request := &chapa.PaymentRequest{
			Amount:         amount,
			Currency:       payment.Currency,
			FirstName:      payment.First_name,
			LastName:       payment.Last_name,
			Email:          payment.Email,
			CallbackURL:    "https://webhook.site/077164d6-29cb-40df-ba29-8a00e59a7e60",
			TransactionRef: chapa.RandomString(20),
		}

		response, err := chapaAPI.PaymentRequest(request)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)

		defer cancel()
		PaymentCollection.InsertOne(ctx, payment)
		c.JSON(http.StatusOK, response)
	}
}

func VerifyTransaction() gin.HandlerFunc {
	return func(c *gin.Context) {
		var payment string
		c.BindJSON(&payment)

		chapa.InitConfig()
		var chapaAPI = chapa.New()

		response, err := chapaAPI.Verify(payment)

		if err != nil {
			c.JSON(http.StatusOK, gin.H{"error": err})
			return
		}
		c.JSON(http.StatusOK, response)
	}
}

func GetTransactions() gin.HandlerFunc {
	return func(c *gin.Context) {
		chapa.InitConfig()
		var chapaAPI = chapa.New()

		response, err := chapaAPI.GetTransactions()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": ""})
		}
		if response.Status != "success" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": response.Message})
		}
		c.JSON(http.StatusOK, response.Data)
	}
}
