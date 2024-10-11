package models

type Tipo struct {
	TipId int64 `json:"tip_id" gorm:"primaryKey;autoIncrement;not null"`
	TipNombre string `json:"tip_nombre" gorm:"not null"`
}

func (Tipo) TableName() string {
	return "tipo"
}