package service

import (
	"api/db/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AddUserBody struct {
	UserID 		uint 		`json:"user_id" binding:"required"`
	Name 		string 		`json:"name" binding:"required"`
	Email		string 		`json:"email" binding:"required,email"`
	CreatedAt 	time.Time 	`json:"created_at"`
}

func AddUser(context *gin.Context, database *gorm.DB) {
	var body AddUserBody

	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if body.CreatedAt.IsZero() {
		body.CreatedAt = time.Now()
	}

	newUser := models.User{
		UserId: body.UserID,
		Name: body.Name,
		Email: body.Email,
		CreatedAt: body.CreatedAt,
	}

	if err := database.Create(&newUser).Error; err != nil {
		context.JSON(500, gin.H{
				"error": 	"Failed to create user", 
				"details": 	err.Error(),
			})
		return
	}

	context.JSON(201, gin.H{
		"message": "User created successfully", 
		"user": 	newUser,
	})
}