package handlers

import (
	"backend/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ActualizarReporte(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		var reporte models.Reportes_Roomie
		if err := db.First(&reporte, idParam).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Reporte no encontrado"})
			return
		}

		var input struct {
			Estado string `json:"Estado"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		reporte.Estado = input.Estado

		if err := db.Save(&reporte).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, reporte)
	}
}
