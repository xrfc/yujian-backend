package model

type AuthDTO struct {
	UserName string `gorm:"column:user_name" json:"user_name"`
	Password string `gorm:"column:password" json:"password"`
}
