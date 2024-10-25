package handlers

import (
	"backend/API/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateMensaje(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var mensaje models.Mensaje_Roomie
		if err := c.ShouldBindJSON(&mensaje); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Validar que el Emisor y Receptor existan
		var emisor, receptor models.Usuario_Roomie
		if err := db.First(&emisor, mensaje.EmisorId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Emisor no encontrado"})
			return
		}
		if err := db.First(&receptor, mensaje.ReceptorId).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Receptor no encontrado"})
			return
		}

		// Establecer la fecha y hora de envío al momento actual
		mensaje.FechaHoraEnvio = time.Now()
		mensaje.Estado = "No Leído"

		// Crear el mensaje en la base de datos
		if err := db.Create(&mensaje).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Crear una notificación para el receptor
		notificacion := models.Notificacion_Roomie{
			IdMensaje:        mensaje.Id,
			IdReceptor:       mensaje.ReceptorId,
			FechaHoraMensaje: mensaje.FechaHoraEnvio,
			Estado:           "No Leída",
		}

		// Crear la notificación en la base de datos
		if err := db.Create(&notificacion).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear la notificación: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, mensaje)
	}
}
