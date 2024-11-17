package main

import (
	"backend/API/database"
	"backend/API/handlers"

	//"backend/API/models"

	"backend/API/config"
	"fmt"
	"log"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	// Cargar variables de entorno desde .env
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}

	// Conectar a la base de datos
	db, err := database.OpenGormDB() //abro la conexion a la base de datos
	if err != nil {
		log.Fatalf("Error al conectarse a la BD: %v", err)
	}

	fmt.Print(config.DBURL())

	// Configurar CORS
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}

	fmt.Printf("db: %v\n", db)

	//Migramos las tablas a la bd
	//db.AutoMigrate(&models.Favorito_Roomie{})

	router := gin.Default()
	router.Use(cors.New(config))

	// Create
	router.POST("/Usuario", handlers.CreateUsuario(db))
	router.POST("/UsuarioRoomie", handlers.CreateUsuarioRoomie(db))
	router.POST("/FavoritosRoomie", handlers.CreateFavorito(db))
	router.POST("/NotificacionesRoomie", handlers.CreateNotificacion(db))
	router.POST("/MensajesRoomie", handlers.CreateMensaje(db))

	// Read
	router.GET("/Usuario/:Firebase_usuario", handlers.GetUsuarioByFirebase(db))
	router.GET("/UsuarioId/:Id", handlers.GetUsuario(db))                                                  // Lectura de un usuario por ID
	router.GET("/Usuarios", handlers.GetallUsuarios(db))                                                   // Lectura de todos los usuarios
	router.GET("/UsuarioRoomie/:Id", handlers.GetUsuarioRoomie(db))                                        // Lectura de un roomie por ID
	router.GET("/UsuarioRoomies", handlers.GetallUsuariosRoomie(db))                                       // Lectura de todos los roomies
	router.GET("/FavoritosRoomie/:Id", handlers.GetFavoritos(db))                                          // Lectura de los favoritos de un usuario por ID
	router.GET("/NotificacionesRoomie/:Id", handlers.GetNotificacion(db))                                  // Lectura de notificaciones por ID
	router.GET("/NotificacionesRoomie", handlers.GetAllNotificaciones(db))                                 // Lectura de todas las notificaciones
	router.GET("/NotificacionesRoomie/UsuarioRoomie/:UsuarioId", handlers.GetNotificacionesPorUsuario(db)) // Lectura de todas las notificaciones de un usuario
	router.GET("/MensajesRoomie/:Id", handlers.GetMensaje(db))                                             // Lectura de mensaje por el ID del mensaje
	router.GET("/MensajesRoomie/UsuarioRoomie/:UsuarioId", handlers.GetMensajesPorUsuario(db))             // Lectura de todos los mensajes de un usuario
	router.GET("/UsuariosconRoomie", handlers.GetUsuariosConRoomie(db))

	//router.GET("/filtrar_usuario", handlers.FilterUsers(db))                     //para filtrar usuarios no se como conectarlo bien a los datos que me pidieron.

	// Update
	router.PUT("/Usuario/:Id", handlers.UpdateUsuario(db))                   // Actualización de un usuario por ID
	router.PUT("/UsuarioRoomie/:Id", handlers.UpdateUsuarioRoomie(db))       // Actualización de un roomie por ID
	router.PUT("/NotificacionesRoomie/:Id", handlers.UpdateNotificacion(db)) // Actualización de notificación por ID
	router.PUT("/MensajesRoomie/:Id", handlers.UpdateMensaje(db))            // Actualización de mensaje por ID

	//Delete
	router.DELETE("/Usuario/:Id", handlers.DeleteUsuario(db))             //no funca si no se elimina la rommie antes
	router.DELETE("/UsuarioRoomie/:Id", handlers.DeleteUsuarioRommie(db)) //funca
	router.DELETE("/FavoritosRoomie/:Id", handlers.DeleteFavorito(db))
	router.DELETE("/NotificacionesRoomie/:Id", handlers.DeleteNotificacion(db))
	router.DELETE("/MensajesRoomie/:Id", handlers.DeleteMensaje(db))

	//Indico el puerto
	router.Run(":8080")

}
