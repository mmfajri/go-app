package repositories

import (
	"errors"
	"go-app/models"
	"log"

	"gorm.io/gorm"
)

type reportRepository struct {
	DB *gorm.DB
}

type ReportRepository interface {
	AddReport(*models.Report) (int, error)
	GetReport() (*[]models.Report, error)
	GetReportByUserId(*uint) (*[]models.Report, error)
	GetReportById(*uint) (*models.Report, error)
	UpdateReport(*models.Report) (int, error)
	DeleteReport(*uint) (int, error)
	Migrate() error
}

func NewReportRepository(db *gorm.DB) ReportRepository {
	return &reportRepository{ // Use pointer
		DB: db,
	}
}

func (u *reportRepository) Migrate() error {
	log.Print("[ReportRepository]...Migrate")
	return u.DB.AutoMigrate(&models.Report{})
}

func (u *reportRepository) GetReportById(id *uint) (*models.Report, error) {
	var data models.Report
	result := u.DB.Model(&models.Report{}).
		Where("id = ? AND is_deleted = ?", id, false).
		Scan(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return &data, nil
}

func (u *reportRepository) AddReport(data *models.Report) (int, error) {
	result := u.DB.Create(&data)
	return int(data.ID), result.Error
}

func (u *reportRepository) UpdateReport(data *models.Report) (int, error) {
	result := u.DB.Model(&models.Report{}).Where("id = ?", data.ID).Updates(data)
	return int(result.RowsAffected), result.Error
}

// SOFT DELETE
func (u *reportRepository) DeleteReport(id *uint) (int, error) {
	result := u.DB.Model(&models.Report{}).Where("id = ?", id).Updates(map[string]any{
		"is_deleted": true,
	})
	return int(result.RowsAffected), result.Error
}

func (u *reportRepository) GetReport() (*[]models.Report, error) {
	var data *[]models.Report
	result := u.DB.Where("is_deleted = ?", false).Find(&data)
	if data != nil {
		return nil, errors.New("Data Not Found")
	}
	return data, result.Error
}

func (u *reportRepository) GetReportByUserId(userId *uint) (*[]models.Report, error) {
	if userId == nil {
		return nil, gorm.ErrRecordNotFound
	}

	var data *[]models.Report
	result := u.DB.Where("is_deleted = ? and user_id = ?", false, *userId).Find(&data)
	return data, result.Error
}
