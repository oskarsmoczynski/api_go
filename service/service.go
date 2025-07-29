package service

import (
	"api/db/models"
	"api/service/schemas"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ReadTable(context *gin.Context, db *gorm.DB) {
	var body schemas.ReadTableSchema
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}
	if err := body.Serialize(); err != nil {
		context.JSON(400, gin.H{"error": "Invalid request", "details": err.Error()})
		return
	}

	model := models.GetModelByName(body.Table, true)
	if model == nil {
		context.JSON(400, gin.H{"error": "Failed to read table from the database"})
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
	if err := body.Serialize(); err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	var results []interface{}
	for _, values := range body.Values {
		entryModel := models.GetModelByName(body.Table, false)
		if entryModel == nil {
			context.JSON(400, gin.H{"error": "Failed to create entry in the database"})
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

func UpdateEntry(context *gin.Context, db *gorm.DB) {
	var body schemas.UpdateEntrySchema
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	if err := body.Serialize(); err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	model := models.GetModelByName(body.Table, false)
    if model == nil {
        context.JSON(400, gin.H{"error": "table not found"})
        return
    }

    query := db
    for _, v := range body.Filters {
        exp, val1, val2 := buildFilters(v)
        if val2 != "" {
            query = query.Where(exp, val1, val2)
        } else {
            query = query.Where(exp, val1)
        }
    }

    updates := make(map[string]interface{})
    for _, v := range body.Values {
        parts := strings.Split(v, "=")
        if len(parts) == 2 {
            updates[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
        }
    }

    if err := query.Model(model).Updates(updates).Error; err != nil {
        context.JSON(500, gin.H{"error": "failed to update entry", "details": err.Error()})
        return
    }

    context.JSON(200, gin.H{"message": "entry updated successfully"})
}

func DeleteEntry(context *gin.Context, db *gorm.DB) {
	var body schemas.DeleteEntrySchema
	if err := context.ShouldBindJSON(&body); err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return	
	}
	if err := body.Serialize(); err != nil {
		context.JSON(400, gin.H{"error": "invalid request", "details": err.Error()})
		return
	}

	model := models.GetModelByName(body.Table, false)
	if model == nil {
		context.JSON(400, gin.H{"error": "Failed to read table from the database"})
		return
	}

	query := db
	for _, v := range body.Filters {
		exp, val1, val2 := buildFilters(v)
		if val2 != "" {
			query = query.Where(exp, val1, val2)
		} else {
			query = query.Where(exp, val1)
		}
	}

	if err := query.Delete(model).Error; err != nil {
		context.JSON(500, gin.H{"error": "Failed to delete entries from the database", "details": err.Error()})
		return
	}

	context.JSON(204, gin.H{})
}
