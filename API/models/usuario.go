package models

type Usuario struct {
	Id                uint           `gorm:"primaryKey;autoIncrement"`
	Firebase_usuario  string         `json:"Firebase_usuario"`
	Roomie            Usuario_Roomie `json:"Usuario_Roomie" gorm:"foreignKey:Id" `
	Correo            string         `json:"Correo"`
	Nombres           string         `json:"Nombres"`
	Apellidos         string         `json:"Apellidos"`
	Fecha_nacimiento  string         `json:"Fecha_Nacimiento"`
	Ano_ingreso       string         `json:"Ano_Ingreso"`
	Id_carrera        uint           `json:"Id_carrera"`
	Id_estado_usuario bool           `json:"Id_Estado_Usuario"`
	Foto_perfil       string         `json:"Foto_Perfil"`
	Fecha_creacion    string         `json:"Fecha_Creacion"`
	Rol               string         `json:"Rol"`
}

func (Usuario) TableName() string {
	return "Usuario"
}
