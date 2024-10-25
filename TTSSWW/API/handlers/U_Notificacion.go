package handlers

import (
	"backend/API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateNotificacion(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("Id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var notificacion models.Notificacion_Roomie
		if err := db.First(&notificacion, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Notificación no encontrada"})
			return
		}

		if err := c.ShouldBindJSON(&notificacion); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Save(&notificacion).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, notificacion)
	}
}
