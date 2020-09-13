package models

type MTExtendTableMeta struct {
	ExtendTable string `json:"extendTable" gorm:"primary_key"`
	MaxColumnNum int `json:"maxColumnNum"`
}
