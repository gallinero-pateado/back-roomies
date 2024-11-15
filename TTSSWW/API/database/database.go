package database

import (
	"backend/API/config"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenGormDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DBURL()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Obtener la conexión subyacente *sql.DB para configurar el pool de conexiones
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Configurar el pool de conexiones
	sqlDB.SetMaxOpenConns(25)                 // Número máximo de conexiones abiertas
	sqlDB.SetMaxIdleConns(25)                 // Número máximo de conexiones inactivas
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Tiempo máximo de vida de una conexión

	// Puedes agregar logs para verificar que se ha configurado correctamente
	log.Println("Pool de conexiones configurado correctamente")

	return db, nil
}
