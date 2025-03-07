package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name		string `json:"name"`
	Email		string `json:"email" gorm:"unique"`
	Role		string `json:"role" gorm:"-"`
	Password	string `json:"password"`	
}

func (User) TableName() string {
	return "tbl_m_user"
} 
