package models

type User struct {
	ID       uint   `gorm:"primarykey" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	Email    string `gorm:"unique" json:"email"`
}
