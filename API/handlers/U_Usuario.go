package handlers

import (
	"backend/API/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUsuario(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("Id")
		var usuarioExistente models.Usuario

		// Intentar obtener el usuario existente
		if err := db.First(&usuarioExistente, id).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			} else {
				log.Printf("Error al obtener el usuario: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener el usuario"})
			}
			return
		}

		// Parsear el cuerpo de la solicitud al modelo Usuario
		var datosActualizados models.Usuario
		if err := c.ShouldBindJSON(&datosActualizados); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		// Procesar la actualización en una goroutine
		go func(usuarioExistente models.Usuario, nuevosDatos models.Usuario) {
			// Actualizar los campos necesarios
			usuarioExistente.Nombres = nuevosDatos.Nombres
			usuarioExistente.Apellidos = nuevosDatos.Apellidos
			usuarioExistente.Correo = nuevosDatos.Correo
			usuarioExistente.Fecha_nacimiento = nuevosDatos.Fecha_nacimiento
			usuarioExistente.Ano_ingreso = nuevosDatos.Ano_ingreso
			usuarioExistente.Id_carrera = nuevosDatos.Id_carrera
			usuarioExistente.Id_estado_usuario = nuevosDatos.Id_estado_usuario
			usuarioExistente.Foto_perfil = nuevosDatos.Foto_perfil
			usuarioExistente.Rol = nuevosDatos.Rol
			// Añade aquí otros campos que desees actualizar

			// Guardar los cambios en la base de datos
			if err := db.Save(&usuarioExistente).Error; err != nil {
				log.Printf("Error al actualizar el usuario: %v", err)
				return
			}
			log.Printf("Usuario actualizado exitosamente: %+v", usuarioExistente)
		}(usuarioExistente, datosActualizados)

		// Responder al cliente inmediatamente
		c.JSON(http.StatusAccepted, gin.H{"message": "Solicitud de actualización recibida y en proceso"})
	}
}
