package models

type Zapato struct {
	ZapId          int64   `json:"zap_id" gorm:"primaryKey;autoIncrement;not null"`
	ZapNombre      string  `json:"zap_nombre" gorm:"not null"`
	ZapDescripcion string  `json:"zap_descripcion" gorm:"not null"`
	ZapCodigo      string  `json:"zap_codigo" gorm:"not null"`
	ZapCantidad    int     `json:"zap_cantidad" gorm:"not null"`
	ZapPrecio      float64 `json:"zap_precio" gorm:"not null"`
	ZapFkmarca     int64   `json:"zap_fkmarca" gorm:"column:zap_fkmarca;not null"`
	//Marca          Marca     `json:"marca" gorm:"foreignKey:ZapFkmarca;references:MarId"`
	ZapFktipo int64 `json:"zap_fktipo" gorm:"column:zap_fktipo;not null"`
	//Tipo           Tipo      `json:"tipo" gorm:"foreignKey:ZapFktipo;references:TipId"`
	ZapFkcolor int64 `json:"zap_fkcolor" gorm:"column:zap_fkcolor;not null"`
	//Color          Color     `json:"color" gorm:"foreignKey:ZapFkcolor;references:ColId"`
	ZapFkcategoria int64 `json:"zap_fkcategoria" gorm:"column:zap_fkcategoria;not null"`
	//Categoria      Categoria `json:"categoria" gorm:"foreignKey:ZapFkcategoria;references:CatId"`
}

func (Zapato) TableName() string {
	return "zapato"
}
