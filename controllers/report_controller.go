package controllers

import (
	"go-app/models"
	"go-app/repositories"
	"go-app/requests"
	"go-app/responses"
	"net/http"

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
	return &reportController{
		reportRepo: repo,
	}
}

func (h *reportController) Delete(ctx *gin.Context) {
	var req requests.ReportById

	err := ctx.ShouldBindBodyWithJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	affectedRows, err := h.reportRepo.DeleteReport(&req.IdReport)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK,
		responses.BaseResponse{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       affectedRows,
		})
}

func (h *reportController) Get(ctx *gin.Context) {
	var req requests.ReportById
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.reportRepo.GetReportById(&req.IdReport)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,
		responses.BaseResponse{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       data,
		})

}

func (h *reportController) Update(ctx *gin.Context) {
	var req requests.ReportRequest
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idReport := uint(req.Id)
	data, err := h.reportRepo.GetReportById(&idReport)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data.Name = req.Name
	data.Content = req.Content

	affectedRows, err := h.reportRepo.UpdateReport(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,
		responses.BaseResponse{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       affectedRows,
		})

}

func (h *reportController) GetAll(ctx *gin.Context) {
	datas, err := h.reportRepo.GetReport()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK,
		responses.BaseResponse{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       datas,
		})
}

func (h *reportController) Add(ctx *gin.Context) {
	var req requests.ReportRequest
	ctx.ShouldBindBodyWithJSON(&req)

	if req.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Name/Title cannot empty"})
		return
	}

	var data = models.Report{
		Name:      req.Name,
		Content:   req.Content,
		IsDeleted: false,
	}

	affectedRows, err := h.reportRepo.AddReport(&data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK,
		responses.BaseResponse{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       affectedRows,
		})
}
