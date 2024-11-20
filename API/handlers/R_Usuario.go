package handlers

import (
	"backend/API/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsuarioByFirebase(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		firebaseUsuario := c.Param("Firebase_usuario")

		var usuario models.Usuario

		if err := db.Where("firebase_usuario = ?", firebaseUsuario).First(&usuario).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}
		c.JSON(http.StatusOK, usuario)
	}
}

func GetUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		id := informacion.Param("Id")

		var usuario models.Usuario

		if err := db.Limit(1).First(&usuario, id).Error; err != nil {
			informacion.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		}
		informacion.JSON(http.StatusOK, usuario)
	}
}

func GetallUsuarios(db *gorm.DB) gin.HandlerFunc {
	return func(informacion *gin.Context) {
		var usuarios []models.Usuario
		err := db.Limit(100).Find(&usuarios).Error
		if err != nil {
			informacion.JSON(http.StatusInternalServerError, gin.H{"error": "Error al buscar todos los usuarios"})
			return
		}
		informacion.JSON(http.StatusOK, usuarios)
	}
}