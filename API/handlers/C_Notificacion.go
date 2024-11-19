package handlers

import (
	"backend/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateNotificacion(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var notificacion models.Notificacion_Roomie
		if err := c.ShouldBindJSON(&notificacion); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&notificacion).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, notificacion)
	}
}
