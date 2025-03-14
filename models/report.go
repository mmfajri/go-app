package models

import (
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	Name 			string 	`json:"name"`
	Content			string 	`json:"content"`
	IsDeleted 		bool 	`json:"is_deleted"`

	User			User	`gorm:"foreignKey:UserID;constraint:OnUpdate:Cascade,OnDelete:SET NULL"`
	UserID			*uint 	`json:"user_id"`
}
	
func (Report) TableName() string {
	return "tbl_m_report"
}
