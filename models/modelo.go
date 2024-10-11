package models

type Modelo struct {
	ModId int64 `json:"mod_id" gorm:"primaryKey;autoIncrement;not null"`
	ModNombre string `json:"mod_nombre" gorm:"not null"`
}

func (Modelo) TableName() string {
	return "modelo"
}