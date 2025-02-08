package controllers

import (
	"carbon/go-commerce/database"
	"carbon/go-commerce/helpers"
	"carbon/go-commerce/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.OpenCollection(*database.Client, "user")

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		var user models.User
		var foundUser models.User

		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer cancel()

		passwordIsValid, msg := VerifyPassword(*user.Password, *foundUser.Password)

		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}

		token, refreshToken, err := helpers.GenerateAllTokens(*foundUser.Email, *foundUser.First_name, *foundUser.Last_name, *foundUser.User_type, *&foundUser.User_id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
		}

		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)
		c.JSON(http.StatusOK, &foundUser)
	}
}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func VerifyPassword(userPassword string, providedPassword string) (check bool, msg string) {
	error := bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
	check = true
	msg = ""

	if error != nil {
		check = false
		msg = fmt.Sprintf("Email or Password is Incorrect")
	}
	return check, msg
}
