package models

type Usuario_Roomie struct {
	Id           uint   `gorm:"primaryKey;autoIncrement"`
	Genero       string `json:"Genero"`
	Biografia    string `json:"Biografia"`
	Intereses    string `json:"Intereses"`
	Preferencias string `json:"Preferencias"`
	Ubicacion    string `json:"Ubicacion"`
}

func (Usuario_Roomie) TableName() string {
	return "Usuario_Roomie"
}
