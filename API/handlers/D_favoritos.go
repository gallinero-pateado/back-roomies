package handlers

import (
	"backend/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteFavorito(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("Id")
		if err := db.Delete(&models.Favorito_Roomie{}, id).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Favorito eliminado"})
	}
}
