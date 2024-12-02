package handlers

import (
	"backend/API/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUsuarioByFirebase(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		firebaseUsuario := c.Param("Firebase_usuario")

		var usuario models.Usuario

		if err := db.Where("firebase_usuario = ?", firebaseUsuario).First(&usuario).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			} else {
				log.Printf("Error al obtener el usuario: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Usuario obtenido exitosamente",
			"usuario": usuario,
		})
	}
}

func GetUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("Id")

		var usuario models.Usuario

		if err := db.First(&usuario, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			} else {
				log.Printf("Error al obtener el usuario: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
			}
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Usuario obtenido exitosamente",
			"usuario": usuario,
		})
	}
}

func GetallUsuarios(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var usuarios []models.Usuario
		if err := db.Limit(100).Find(&usuarios).Error; err != nil {
			log.Printf("Error al obtener los usuarios: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los usuarios"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message":  "Usuarios obtenidos exitosamente",
			"usuarios": usuarios,
		})
	}
}
