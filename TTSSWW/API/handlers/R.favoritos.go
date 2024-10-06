package handlers

import (
	"backend/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetFavoritos(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var favoritos []models.Favorito

		// Obtener el usuario_id desde los parámetros de la URL
		usuarioId := c.Param("Id")

		// Buscar los favoritos del usuario específico
		if err := db.Where("usuario_id = ?", usuarioId).Find(&favoritos).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, favoritos)
	}
}
