package controllers

import (
	"fmt"
	"go-app/models"
	"go-app/repositories"
	"go-app/utils"
	"net/http"
	"strconv"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
)

type userController struct {
	userRepo repositories.UserRepository
}

type UserController interface {
	AddUser(enforcer *casbin.Enforcer) gin.HandlerFunc
	GetUser(*gin.Context)
	GetAllUser(*gin.Context)
	SignInUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
}

func NewUserController(repo repositories.UserRepository) UserController {
	return userController {
		userRepo: repo,
	}
}

func (h userController) GetAllUser(ctx *gin.Context) {
	fmt.Println(ctx.Get("userID"))
	users, err := h.userRepo.GetAllUser()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (h userController) GetUser(ctx *gin.Context) {
	id := ctx.Param("user")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userRepo.GetUser(intID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h userController) SignInUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindBodyWithJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	dbUser, err := h.userRepo.GetByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if isTrue := utils.ComparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID)
		token := utils.GenerateToken(dbUser.ID)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully SignIn", "token": token})
		return
	} 
	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Password not matched"})
	return
}

func (h userController) AddUser(enforcer *casbin.Enforcer) gin.HandlerFunc {
	return func(ctx *gin.Context){
		var user models.User
		if err := ctx.ShouldBindJSON(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		utils.HashPassword(&user.Password)
		user, err := h.userRepo.AddUser(user)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		enforcer.AddGroupingPolicy(fmt.Sprint(user.ID), user.Role)
		user.Password = ""
		ctx.JSON(http.StatusOK, user)
	}
}

func (h userController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("user")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.ID = uint(intID)
	user, err = h.userRepo.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h userController) DeleteUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("user")
	intID, _ := strconv.Atoi(id)
	user.ID = uint(intID)
	user, err := h.userRepo.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

