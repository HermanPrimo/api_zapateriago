package models

type Pago struct {
	PagId int64 `json:"pag_id" gorm:"primaryKey;autoIncrement;not null"`
	PagNombre string `json:"pag_nombre" gorm:"not null"`
}

func (Pago) TableName() string {
	return "pago"
}