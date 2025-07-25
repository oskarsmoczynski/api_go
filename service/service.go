package service

import (
	"api/db/models"
	"api/service/schemas"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AddUser(context *gin.Context, db *gorm.DB) {
	var body schemas.AddUserBody
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	body, err := body.Serialize()
	if err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	newUser := models.User{}.FromStruct(body)
	if err := db.Select("Name", "Email", "CreatedAt").Create(&newUser).Error; err != nil {
		context.JSON(500, gin.H{
			"error":   "Failed to create user",
			"details": err.Error(),
		})
		return
	}

	context.JSON(201, gin.H{
		"message": "User created successfully",
		"user":    newUser,
	})
}
