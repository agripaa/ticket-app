package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Name     string `gorm:"varchar(255)" json:"name"`
	Username string `gorm:"varchar(255)" json:"username"`
	Password string `gorm:"varchar(255)" json:"password"`
}
