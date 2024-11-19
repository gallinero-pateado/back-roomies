package models

import (
	"time"
)

type Favorito_Roomie struct {
	IdFavorito        uint      `gorm:"primaryKey;autoIncrement" json:"id_favorito"`
	UsuarioId         int       `json:"usuario_id" binding:"required"`
	UsuarioFavoritoId int       `json:"usuario_favorito_id" binding:"required"`
	FechaFavorito     time.Time `json:"fecha_favorito" gorm:"autoCreateTime"`
}

func (Favorito_Roomie) TableName() string {
	return "favoritos"
}
