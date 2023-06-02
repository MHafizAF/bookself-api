package models

type User struct {
	Id       int64  `gorm:"primaryKey" json:"id"`
	Username string `gorm:"type:varchar(255); not null; unique" json:"username"`
	Password string `gorm:"type:varchar(255); not null" json:"password"`
	Books    []Book
}
