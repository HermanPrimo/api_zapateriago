package models

type DetalleVenta struct {
	DetvenId       int64   `json:"detven_id" gorm:"primaryKey;autoIncrement;not null"`
	DetvenCantidad int     `json:"detven_cantidad" gorm:"not null"`
	DetvenMonto    float64 `json:"detven_monto" gorm:"not null"`
	DetvenFkventa  int64   `json:"detven_fkventa" gorm:"column:detven_fkventa;not null"`
	//Venta          Venta   `json:"venta" gorm:"foreignKey:DetvenFkventa;references:VenId"`
	DetvevFkzapato int64 `json:"detvev_fkzapato" gorm:"column:detvev_fkzapato;not null"`
	//Zapato         Zapato  `json:"zapato" gorm:"foreignKey:DetvevFkzapato;references:ZapId"`
}

func (DetalleVenta) TableName() string {
	return "detalle_venta"
}
