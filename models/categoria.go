package models

type Categoria struct {
	CatId int64 `json:"cat_id" gorm:"primaryKey;autoIncrement;not null"`
	CatNombre string `json:"cat_nombre" gorm:"not null"`
}

func (Categoria) TableName() string {
	return "categoria"
}