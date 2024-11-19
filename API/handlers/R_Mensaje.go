package handlers

import (
	"backend/API/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMensaje(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("Id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var mensaje models.Mensaje_Roomie
		if err := db.Preload("Emisor").Preload("Receptor").First(&mensaje, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Mensaje no encontrado"})
			return
		}
		// Actualizar el estado a "Leído"
		mensaje.Estado = "Leído"
		db.Save(&mensaje)

		c.JSON(http.StatusOK, mensaje)
	}
}

func GetMensajesRecibidosPorUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioIdParam := c.Param("UsuarioId")
		usuarioId, err := strconv.Atoi(usuarioIdParam)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
			return
		}

		var mensajes []models.Mensaje_Roomie
		if err := db.Preload("Emisor").Preload("Receptor").Where("receptor_id = ?", usuarioId).Find(&mensajes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, mensajes)
	}
}

func GetMensajesEnviadosPorUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		usuarioIdParam := c.Param("UsuarioId")
		usuarioId, err := strconv.Atoi(usuarioIdParam)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario inválido"})
			return
		}

		var mensajes []models.Mensaje_Roomie
		if err := db.Preload("Emisor").Preload("Receptor").Where("emisor_id = ?", usuarioId).Find(&mensajes).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, mensajes)
	}
}
