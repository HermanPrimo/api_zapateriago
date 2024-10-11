package models

type Genero struct {
	GenId int64 `json:"gen_id" gorm:"primaryKey;autoIncrement;not null"`
	GenNombre string `json:"gen_nombre" gorm:"not null"`
}

func (Genero) TableName() string {
	return "genero"
}