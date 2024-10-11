package models

type Estado struct {
	EstId int64 `json:"est_id" gorm:"primaryKey;autoIncrement;not null"`
	EstNombre string `json:"est_nombre" gorm:"not null"`
}

func (Estado) TableName() string {
	return "estado"
}