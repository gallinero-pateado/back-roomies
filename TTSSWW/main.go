package main

import (
	"backend/API/config"
	"backend/API/database"
	"backend/API/handlers"
	//"backend/API/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Hola mundo")

	// Conectar a la base de datos
	db, err := database.OpenGormDB() //abro la conexion a la base de datos
	if err != nil {
		log.Fatalf("Error al conectarse a la BD: %v", err)
	}
	fmt.Print(config.DBURL())

	fmt.Printf("db: %v\n", db)

	//Migramos las tablas a la bd
	//db.AutoMigrate(&models.Usuario{}, &models.Usuario_Roomie{})

	router := gin.Default()

	// Create
	router.POST("/Usuario", handlers.CreateUsuario(db))
	router.POST("/UsuarioRoomie", handlers.CreateUsuarioRoomie(db))
	router.POST("/favorites", handlers.CreateFavorito(db))

	// Read
	router.GET("/Usuario/:Id", handlers.GetUsuario(db))              // Lectura de un usuario por ID
	router.GET("/Usuarios", handlers.GetallUsuarios(db))             // Lectura de todos los usuarios
	router.GET("/UsuarioRoomie/:Id", handlers.GetUsuarioRoomie(db))  // Lectura de un roomie por ID
	router.GET("/UsuarioRoomies", handlers.GetallUsuariosRoomie(db)) // Lectura de todos los roomies
	router.GET("/favorites/:Id", handlers.GetFavoritos(db))          // Lectura de los favoritos de un usuario por id

	// Update
	router.PUT("/Usuario/:Id", handlers.UpdateUsuario(db))             // Actualización de un usuario por ID
	router.PUT("/UsuarioRoomie/:Id", handlers.UpdateUsuarioRoomie(db)) // Actualización de un roomie por ID

	//Delete
	router.DELETE("/Usuario/:Id", handlers.DeleteUsuario(db))             //no funca si no se elimina la rommie antes
	router.DELETE("/UsuarioRoomie/:Id", handlers.DeleteUsuarioRommie(db)) //funca
	router.DELETE("/favorites/:Id", handlers.DeleteFavorito(db))
	//Indico el puerto
	router.Run(":8080")

}
