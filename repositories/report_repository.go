package repositories

import (
	"go-app/models"
	"log"

	"gorm.io/gorm"
)

type reportRepository struct {
	DB *gorm.DB
}

type ReportRepository interface {
	AddReport(models.Report) (int, error)
	GetReport() ([]models.Report, error)
	GetReportByUserId(uint) ([]models.Report, error)
	UpdateReport(models.Report) (int, error)
	DeleteReport(uint) (int, error)
	Migrate() error
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return reportRepository{
		DB: db,
	}
}

func (u reportRepository) Migrate() error {
	log.Print("[ReportRepository]...Migrate")
	return u.DB.AutoMigrate(&models.Report{})
}

func (u reportRepository) AddReport(data models.Report) (int, error){
	result := u.DB.Create(&data)
	return int(data.ID), result.Error
}

func (u reportRepository) UpdateReport(data models.Report) (int, error) {
	result := u.DB.Model(&models.Report{}).Where("id = ?",data.ID).Updates(data)
	return int(result.RowsAffected), result.Error
}

// SOFT DELETE
func (u reportRepository) DeleteReport(id uint) (int, error) {
	result := u.DB.Model(&models.Report{}).Where("id = ?", id).Update("is_deleted = ?", true)
	return int(result.RowsAffected), result.Error
}


func (u reportRepository) GetReport() ([]models.Report, error) {
	var data []models.Report
	result := u.DB.Where("is_deleted = ?", false).Find(&data) 
	return data, result.Error
}

func (u reportRepository) GetReportByUserId(userId uint) ([]models.Report, error) {
	var data []models.Report
	result := u.DB.Where("is_deleted = ? and user_id = ?", false, userId).Find(&data)
	return data, result.Error
}
