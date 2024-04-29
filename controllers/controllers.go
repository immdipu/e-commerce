package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/immdipu/e-commerce-go/database"
	"github.com/immdipu/e-commerce-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Signup(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	count, err := database.User.CountDocuments(c, bson.M{"email": user.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to check user",
		})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user already exists",
		})
		return
	}

	count, err = database.User.CountDocuments(c, bson.M{"phone": user.Phone})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to check user",
		})
		return
	}

	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "phone number already exists",
		})
		return
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	res, err := database.User.InsertOne(c, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create new user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": res,
	})
}
