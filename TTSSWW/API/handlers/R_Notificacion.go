package handlers

import (
	"backend/API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetNotificacion(db *gorm.DB) gin.HandlerFunc {
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

		c.JSON(http.StatusOK, notificacion)
	}
}

func GetAllNotificaciones(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var notificaciones []models.Notificacion_Roomie
		if err := db.Find(&notificaciones).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, notificaciones)
	}
}

func GetNotificacionesPorUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioIdParam := c.Param("UsuarioId")
		usuarioId, err := strconv.Atoi(usuarioIdParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
			return
		}

		var notificaciones []models.Notificacion_Roomie
		if err := db.Where("id_receptor = ?", usuarioId).Find(&notificaciones).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, notificaciones)
	}
}
