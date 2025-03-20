package controllers

import (
	"go-app/requests"
	"net/http"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type roleController struct {
	casbin *casbin.Enforcer
}

type RoleController interface {
	GetAll(*gin.Context)
	Add(*gin.Context)
	Delete(*gin.Context)
}

func NewRoleController(repo *casbin.Enforcer) RoleController {
	return &roleController {
		casbin: repo,
	}	
}

func (h roleController) GetAll(ctx *gin.Context) {
	var req requests.RoleRequest
	ctx.ShouldBindBodyWithJSON(&req)	

	ctx.JSON(http.StatusOK, req)
}

func (h roleController) Add(ctx *gin.Context) {
	var req requests.RoleRequest
	ctx.ShouldBindBodyWithJSON(&req)	

	ctx.JSON(http.StatusOK, req)
}

func (h roleController) Delete(ctx *gin.Context) {
	var req requests.RoleRequest
	ctx.ShouldBindBodyWithJSON(&req)	

	ctx.JSON(http.StatusOK, req)
}
