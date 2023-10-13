package server

import (
	"Loco/controller"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// InitializeRoutes function provides entry point for application wide routing
func InitializeRoutes(db *gorm.DB, router *gin.Engine) {
	baseRoute := router.Group("")
	controller.NewTransactionController(db, baseRoute)
}
