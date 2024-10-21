package handlers

import (
    "backend/api/models"
    "backend/api/utils"
    "net/http"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

// FiltrarUsuarios maneja la solicitud de filtrado de usuarios
func FiltrarUsuarios(c *gin.Context) {
    ubicacion := c.Query("ubicacion")
    interes := c.Query("interes")
    preferencias := c.Query("preferencias")

    usuarios, err := models.FiltrarUsuarios(ubicacion, interes, preferencias)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, usuarios)
}