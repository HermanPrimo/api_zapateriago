package models

type ZapatoTalla struct {
	ZaptalId       int64 `json:"zaptal_id" gorm:"primaryKey;autoIncrement;not null"`
	ZaptalFkzapato int64 `json:"zaptal_fkzapato" gorm:"column:zaptal_fkzapato;not null"`
	//Zapato         Zapato `json:"zapato" gorm:"foreignKey:ZaptalFkzapato;references:ZapId"`
	ZaptalFktalla int64 `json:"zaptal_fktalla" gorm:"column:zaptal_fktalla;not null"`
	//Talla          Talla  `json:"talla" gorm:"foreignKey:ZaptalFktalla;references:TalId"`
}

func (ZapatoTalla) TableName() string {
	return "zapato_talla"
}
