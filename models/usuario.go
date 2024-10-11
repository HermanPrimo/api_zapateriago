package models

import "time"

type Usuario struct {
	UsuId         int64     `json:"usu_id" gorm:"primaryKey;autoIncrement;not null"`
	UsuNombre     string    `json:"usu_nombre" gorm:"not null"`
	UsuPaterno    string    `json:"usu_paterno" gorm:"not null"`
	UsuMaterno    string    `json:"usu_materno" gorm:"not null"`
	UsuNacimiento time.Time `json:"usu_nacimiento" gorm:"not null"`
	UsuFkgenero   int64     `json:"usu_fkgenero" gorm:"column:usu_fkgenero;not null"`
	//Genero        Genero    `json:"genero" gorm:"foreignKey:UsuFkgenero;references:GenId"`
	UsuTelefono  string `json:"usu_telefono" gorm:"not null"`
	UsuCorreo    string `json:"usu_correo" gorm:"not null"`
	UsuDireccion string `json:"usu_direccion" gorm:"not null"`
}

func (Usuario) TableName() string {
	return "usuario"
}
