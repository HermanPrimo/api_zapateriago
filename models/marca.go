package models

type Marca struct {
	MarId int64 `json:"mar_id" gorm:"primaryKey;autoIncrement;not null"`
	MarNombre string `json:"mar_nombre" gorm:"not null"`
}

func (Marca) TableName() string {
	return "marca"
}