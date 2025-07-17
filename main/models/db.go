package models

type LmLicencia struct {
	Id  string `gorm:"primaryKey;column:id"`
	Num string `gorm:"column:numero"`
}

func (LmLicencia) TableName() string {
	return "lm_licencia"
}