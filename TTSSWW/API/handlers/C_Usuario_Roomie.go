package handlers

import (
	"backend/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateUsuarioRoomie(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		var usuarioRoomie models.Usuario_Roomie
		if err := informacion.ShouldBindJSON(&usuarioRoomie); err != nil {
			informacion.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		// creamos el Usuario en la bd y controlamos el error
		if err := db.Create(&usuarioRoomie).Error; err != nil {
			informacion.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		informacion.JSON(http.StatusOK, usuarioRoomie)
	}
}
