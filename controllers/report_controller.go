package controllers

import (
	"go-app/repositories"

	"github.com/gin-gonic/gin"
)

type reportController struct {
	reportRepo repositories.ReportRepository
}

type ReportController interface {
	Add(*gin.Context)
	GetAll(*gin.Context)
	Get(*gin.Context)
	Update(*gin.Context)
	Delete(*gin.Context)
}

func NewReportController(repo repositories.ReportRepository) ReportController {
	return reportController {
		reportRepo: repo,
	}
}

