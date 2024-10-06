package handlers

import (
	"backend/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUsuarioRoomie(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		id := informacion.Param("Id")
		var usuario models.Usuario_Roomie

		if err := db.First(&usuario, id).Error; err != nil {
			informacion.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}

		// Binding (parsear) el cuerpo de la solicitud (JSON) al modelo Usuario
		if err := informacion.ShouldBindJSON(&usuario); err != nil {
			informacion.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		if err := db.Save(&usuario).Error; err != nil {
			informacion.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el usuario"})
			return
		}
		informacion.JSON(http.StatusOK, gin.H{"message": "Usuario actualizado exitosamente", "usuario": usuario})
	}
}
