package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateDefaultRouter() *gin.Engine {
	return gin.Default()
}

func AppendRoute(router *gin.Engine, db *gorm.DB, method, endpoint string, fn func(*gin.Context, *gorm.DB)) {
	switch strings.ToUpper(method) {
	case http.MethodGet:
		router.GET(endpoint, func(c *gin.Context) {
			fn(c, db)
		})
	case http.MethodPost:
		router.POST(endpoint, func(c *gin.Context) {
			fn(c, db)
		})
	case http.MethodPut:
		router.PUT(endpoint, func(c *gin.Context) {
			fn(c, db)
		})
	case http.MethodPatch:
		router.PATCH(endpoint, func(c *gin.Context) {
			fn(c, db)
		})
	case http.MethodDelete:
		router.DELETE(endpoint, func(c *gin.Context) {
			fn(c, db)
		})
	default:
		err_msg := fmt.Sprintf("Invalid method: %v. Allowed methods: %v, %v, %v, %v, %v", method, http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete)
		panic(err_msg)
	}
}
