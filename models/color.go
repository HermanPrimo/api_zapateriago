package models

type Color struct {
	ColId int64 `json:"col_id" gorm:"primaryKey;autoIncrement;not null"`
	ColNombre string `json:"col_nombre" gorm:"not null"`
}

func (Color) TableName() string {
	return "color"
}