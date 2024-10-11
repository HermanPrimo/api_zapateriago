package models

import "time"

type Venta struct {
	VenId         int64     `json:"ven_id" gorm:"primaryKey;autoIncrement;not null"`
	VenFolio      string    `json:"ven_folio" gorm:"not null"`
	VenFecha      time.Time `json:"ven_fecha" gorm:"not null"`
	VenFkempleado int64     `json:"ven_fkempleado" gorm:"column:ven_fkempleado;not null"`
	//Empleado      Empleado  `json:"empleado" gorm:"foreignKey:VenFkempleado;references:EmpId"`
	VenFkusuario int64 `json:"ven_fkusuario" gorm:"column:ven_fkusuario;not null"`
	//Usuario       Usuario   `json:"usuario" gorm:"foreignKey:VenFkusuario;references:UsuId"`
	VenFkpago int64 `json:"ven_fkpago" gorm:"column:ven_fkpago;not null"`
	//Pago          Pago      `json:"pago" gorm:"foreignKey:VenFkpago;references:PagId"`
}

func (Venta) TableName() string {
	return "venta"
}
