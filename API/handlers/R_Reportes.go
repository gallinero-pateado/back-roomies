package handlers

import (
	"backend/API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ObtenerReportes(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var reportes []models.Reportes_Roomie
		if err := db.Preload("UsuarioReportado").Preload("UsuarioReportante").Find(&reportes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, reportes)
	}
}

func ObtenerReportePorID(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		var reporte models.Reportes_Roomie
		if err := db.Preload("UsuarioReportado").Preload("UsuarioReportante").First(&reporte, idParam).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Reporte no encontrado"})
			return
		}
		c.JSON(http.StatusOK, reporte)
	}
}

func ObtenerReportesEnviadosPorUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioIdParam := c.Param("UsuarioId")
		usuarioId, err := strconv.Atoi(usuarioIdParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
			return
		}

		var reportes []models.Reportes_Roomie
		if err := db.Preload("UsuarioReportado").Preload("UsuarioReportante").
			Where("usuario_reportante_id = ?", usuarioId).Find(&reportes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, reportes)
	}
}
