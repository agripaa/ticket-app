package models

type Product struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	NameProduct string `gorm:"varchar(255)" json:"name_product"`
	Desc        string `gorm:"varchar(255)" json:"desc"`
	Price       string    `gorm:"varchar(255)" json:"price"`
}
