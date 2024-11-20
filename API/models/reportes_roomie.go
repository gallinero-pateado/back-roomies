package models

import (
	"time"
)

type Reportes_Roomie struct {
	ID                  uint      `gorm:"primaryKey;autoIncrement"`
	UsuarioReportadoID  uint      `json:"UsuarioReportadoID"`
	UsuarioReportado    Usuario   `gorm:"foreignKey:UsuarioReportadoID;references:Id" json:"UsuarioReportado"`
	UsuarioReportanteID uint      `json:"UsuarioReportanteID"`
	UsuarioReportante   Usuario   `gorm:"foreignKey:UsuarioReportanteID;references:Id" json:"UsuarioReportante"`
	Motivo              string    `json:"Motivo"`
	FechaHora           time.Time `json:"FechaHora"`
	Estado              string    `json:"Estado"`
}

func (Reportes_Roomie) TableName() string {
	return "Reporte"
}
