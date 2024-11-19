package handlers

import (
	"backend/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {

		id := informacion.Param("Id")
		var usuario models.Usuario

		// Buscar el usuario por ID
		if err := db.First(&usuario, id).Error; err != nil {
			informacion.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return // Retorna el error si no encuentra el usuario
		}
		// Eliminar el usuario
		if err := db.Delete(&usuario).Error; err != nil {
			informacion.JSON(http.StatusInternalServerError, gin.H{"error": "Error al eliminar el usuario"})
			return // Retorna el error si la eliminación falla
		}

		informacion.JSON(http.StatusOK, usuario)
	}
}

/*




 */
