package handlers

import (
	"backend/API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CrearReporte(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reporte models.Reportes_Roomie

		if err := c.ShouldBindJSON(&reporte); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Establecer la fecha y hora actual
		reporte.FechaHora = time.Now()
		reporte.Estado = "Pendiente"

		// Validar que los usuarios existen
		var usuarioReportado models.Usuario
		if err := db.First(&usuarioReportado, reporte.UsuarioReportadoID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario reportado no existe"})
			return
		}

		var usuarioReportante models.Usuario
		if err := db.First(&usuarioReportante, reporte.UsuarioReportanteID).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Usuario reportante no existe"})
			return
		}

		if err := db.Create(&reporte).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, reporte)
	}
}
