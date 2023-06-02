package models

type Book struct {
	Id          int64 `gorm:"primaryKey" json:"id"`
	UserID      int64
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Description string `gorm:"type:text" json:"description"`
}
