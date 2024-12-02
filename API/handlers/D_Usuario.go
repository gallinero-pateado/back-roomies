package handlers

import (
	"backend/API/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("Id")
		var usuario models.Usuario

		// Buscar el usuario por ID
		if err := db.First(&usuario, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			} else {
				log.Printf("Error al buscar el usuario: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar el usuario"})
			}
			return
		}

		// Eliminar el usuario
		if err := db.Delete(&usuario).Error; err != nil {
			log.Printf("Error al eliminar el usuario: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
			return
		}

		// Responder con un mensaje de éxito
		c.JSON(http.StatusOK, gin.H{
			"message": "Usuario eliminado exitosamente",
			"usuario": usuario,
		})
	}
}
