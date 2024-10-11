package models

type Empleado struct {
	EmpId       int64 `json:"emp_id" gorm:"primaryKey;autoIncrement;not null"`
	EmpFkpuesto int64 `json:"emp_fkpuesto" gorm:"column:emp_fkpuesto;not null"`
	//Puesto       Puesto  `json:"puesto" gorm:"foreignKey:EmpFkpuesto;references:PueId"`
	EmpFkestado int64 `json:"emp_fkestado" gorm:"column:emp_fkestado;not null"`
	//Estado       Estado  `json:"estado" gorm:"foreignKey:EmpFkestado;references:EstId"`
	EmpFkusuario int64 `json:"emp_fkusuario" gorm:"column:emp_fkusuario;not null"`
	//Usuario      Usuario `json:"usuario" gorm:"foreignKey:EmpFkusuario;references:UsuId"`
	EmpRfc string `json:"emp_rfc" gorm:"not null"`
}

func (Empleado) TableName() string {
	return "empleado"
}
