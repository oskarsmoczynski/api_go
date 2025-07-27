package service

import (
	"api/db/models"
	"api/service/schemas"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadTable(context *gin.Context, db *gorm.DB) {
	var body schemas.ReadTableSchema
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}
	body, err := body.Serialize()
	if err != nil {
		context.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	model, err := models.GetModelByName(body.Table, true)
	if err != nil {
		context.JSON(400, gin.H{"error": "Failed to read table from the database", "details": err.Error()})
		return
	}

	query := db
	for _, v := range body.OrderBy {
		query = query.Order(v)
	}
	if body.Limit > 0 {
		query = query.Limit(int(body.Limit))
	}
	for _, v := range body.Filters {
		exp, val1, val2 := buildFilters(v)
		if val2 != "" {
			query = query.Where(exp, val1, val2)
		} else {
			query = query.Where(exp, val1)
		}
	}
	if err := query.Find(model).Error; err != nil {
		context.JSON(400, gin.H{"error": "Failed to read table from the database", "details": err.Error()})
		return
	}

	context.JSON(200, gin.H{"data": model})
}

func CreateEntry(context *gin.Context, db *gorm.DB) {
	var body schemas.CreateEntrySchema
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}
	// TODO: body serialization

	var results []interface{}
	for _, values := range body.Values {
		entryModel, err := models.GetModelByName(body.Table, false)
		if err != nil {
			context.JSON(400, gin.H{"error": "Failed to create entry in the database", "detail": err.Error()})
			return
		}

		if err := models.ModelFromStruct(values, entryModel); err != nil {
			context.JSON(400, gin.H{"error": "Failed to create entry in the database", "detail": err.Error()})
			return
		}
				
		if err := db.Create(entryModel).Error; err != nil {
			context.JSON(400, gin.H{"error": "Failed to create entry in the database", "detail": err.Error()})
			return
		}

		results = append(results, entryModel)
	}

	context.JSON(200, gin.H{"data": results})

}
