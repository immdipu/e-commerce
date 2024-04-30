package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/immdipu/e-commerce-go/database"
	"github.com/immdipu/e-commerce-go/models"
	tokens "github.com/immdipu/e-commerce-go/token"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password *string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*password), 14)
	if err != nil {
		log.Panic("something went wrong while hasing password", err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword, givenPassword *string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(*userPassword), []byte(*givenPassword))
	message := ""
	valid := true
	if err != nil {
		valid = false
		message = "email or password is incorrect"
	}
	return valid, message
}

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

	token, err := tokens.GenerateToken(user.Email)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to generate token",
		})
		return
	}

	password := HashPassword(&user.Password)

	user.ID = primitive.NewObjectID()
	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Token = token
	user.Password = password
	_, err = database.User.InsertOne(c, user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to createnew user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"message": "user created successfully",
		},
		"data": gin.H{
			"_id":        user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"email":      user.Email,
			"phone":      user.Phone,
			"token":      user.Token,
			"created_at": user.CreatedAt,
			"updated_at": user.UpdatedAt,
		},
	})
}

func Login(c *gin.Context) {
	var user models.User
	var foundUser models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": gin.H{
				"success": false,
				"message": "Something went wrong.",
			},
		})
		return
	}

	err = database.User.FindOne(c, bson.M{"email": user.Email}).Decode(&foundUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{
				"success": false,
				"message": "can't find the user",
			},
		})
		return
	}

	isVerified, _ := VerifyPassword(&foundUser.Password, &user.Password)

	if !isVerified {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{
				"success": false,
				"message": "username or password is incorrect",
			},
		})
		return
	}

	token, err := tokens.GenerateToken(user.Email)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": gin.H{
				"success": false,
				"message": "Cannot generate token",
			},
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": gin.H{
			"message": "user created successfully",
		},
		"data": gin.H{
			"_id":        foundUser.ID,
			"first_name": foundUser.FirstName,
			"last_name":  foundUser.LastName,
			"email":      foundUser.Email,
			"phone":      foundUser.Phone,
			"token":      token,
			"created_at": foundUser.CreatedAt,
			"updated_at": foundUser.UpdatedAt,
		},
	})
}
