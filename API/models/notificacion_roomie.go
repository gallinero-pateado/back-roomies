package models

import (
	"time"
)

type Notificacion_Roomie struct {
	Id               uint           `gorm:"primaryKey;autoIncrement"`
	IdMensaje        uint           `json:"IdMensaje"`
	IdReceptor       uint           `json:"IdReceptor"`
	Mensaje          Mensaje_Roomie `gorm:"foreignKey:IdMensaje;references:Id" json:"Mensaje_Roomie,omitempty"`
	Receptor         Usuario_Roomie `gorm:"foreignKey:IdReceptor;references:Id" json:"Receptor,omitempty"`
	FechaHoraMensaje time.Time      `json:"FechaHoraMensaje"`
	Estado           string         `json:"Estado"`
}

func (Notificacion_Roomie) TableName() string {
	return "Notificaciones"
}
