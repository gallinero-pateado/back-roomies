package models

import (
	"time"
)

type Mensaje_Roomie struct {
	Id             uint           `gorm:"primaryKey;autoIncrement"`
	EmisorId       uint           `json:"EmisorId"`
	ReceptorId     uint           `json:"ReceptorId"`
	Emisor         Usuario_Roomie `gorm:"foreignKey:EmisorId;references:Id" json:"Emisor"`
	Receptor       Usuario_Roomie `gorm:"foreignKey:ReceptorId;references:Id" json:"Receptor"`
	Asunto         string         `json:"Asunto"`
	Contenido      string         `json:"Contenido"`
	FechaHoraEnvio time.Time      `json:"FechaHoraEnvio"`
	Estado         string         `json:"Estado"`
}

func (Mensaje_Roomie) TableName() string {
	return "Mensajes"
}
