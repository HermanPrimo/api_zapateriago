package models

type Puesto struct {
	PueId int64 `json:"pue_id" gorm:"primaryKey;autoIncrement;not null"`
	PueNombre string `json:"pue_nombre" gorm:"not null"`
}

func (Puesto) TableName() string {
	return "puesto"
}