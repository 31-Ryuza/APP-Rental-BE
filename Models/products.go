package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Code string `gorm:"type:varchar(300)" json:"code"`
	Name   string `gorm:"type:text" json:"name"`
	Price   string `gorm:"type:text" json:"price"`
	IsReady   bool `gorm:"type:text" json:"idReady"`
	Photo   string `gorm:"type:text" json:"photo"`
}