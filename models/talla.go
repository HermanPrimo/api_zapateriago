package models

type Talla struct {
	TalId int64 `json:"tal_id" gorm:"primaryKey;autoIncrement;not null"`
	TalNombre string `json:"tal_nombre" gorm:"not null"`
}

func (Talla) TableName() string {
	return "talla"
}