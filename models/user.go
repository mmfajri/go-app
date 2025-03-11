package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name		string `json:"name"`
	Email		string `json:"email" gorm:"unique"`
	Role		string `json:"role" gorm:"-"`
	Password	string `json:"password"`	
	IsDeleted 	bool   `json:"is_deleted"`	

	Reports		[]Report `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (User) TableName() string {
	return "tbl_m_user"
} 
