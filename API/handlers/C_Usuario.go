package handlers

import (
	"backend/API/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var usuario models.Usuario
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Procesar la creación del usuario en una goroutine
		go func(u models.Usuario) {
			if err := db.Create(&u).Error; err != nil {
				// Manejar el error (por ejemplo, registrarlo)
				log.Printf("Error al crear el usuario: %v", err)
				return
			}
			log.Printf("Usuario creado exitosamente: %+v", u)
		}(usuario)

		// Respuesta al cliente
		c.JSON(http.StatusAccepted, gin.H{"message": "Solicitud recibida y en proceso"})
	}
}
