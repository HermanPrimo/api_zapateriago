package models

type ZapatoModelo struct {
	ZapmodId       int64 `json:"zapmod_id" gorm:"primaryKey;autoIncrement;not null"`
	ZapmodFkzapato int64 `json:"zapmod_fkzapato" gorm:"column:zapmod_fkzapato;not null"`
	//Zapato         Zapato `json:"zapato" gorm:"foreignKey:ZapmodFkzapato;references:ZapId"`
	ZapmodFkmodelo int64 `json:"zapmod_fkmodelo" gorm:"column:zapmod_fkmodelo;not null"`
	//Modelo         Modelo `json:"modelo" gorm:"foreignKey:ZapmodFkmodelo;references:ModId"`
}

func (ZapatoModelo) TableName() string {
	return "zapato_modelo"
}
