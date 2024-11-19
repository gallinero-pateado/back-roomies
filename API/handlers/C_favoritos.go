package handlers

import (
	"backend/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateFavorito(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var favorito models.Favorito_Roomie
		if err := c.ShouldBindJSON(&favorito); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Crear el nuevo favorito en la base de datos
		if err := db.Create(&favorito).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, favorito)
	}
}
