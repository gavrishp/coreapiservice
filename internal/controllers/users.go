package controllers

import (
	"net/http"

	"github.com/gavrishp/coreapiservicetest/internal/dto"
	"github.com/gavrishp/coreapiservicetest/internal/services"
	"github.com/gin-gonic/gin"
)

type IUserService interface {
	AddUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type UsersController struct {
	usersService *services.UsersService
}

func NewUsersController(usersService *services.UsersService) *UsersController {
	return &UsersController{usersService: usersService}
}

func (c *UsersController) AddUser(ctx *gin.Context) {
	body := dto.AddUserRequestBody{}

	// getting request's body
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := c.usersService.AddUser(&body)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusCreated, &response)
}

func (c *UsersController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	response, err := c.usersService.GetUser(userId)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func (c *UsersController) GetUsers(ctx *gin.Context) {

	response, err := c.usersService.GetUsers()
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, &response)
}

func (c *UsersController) DeleteUser(ctx *gin.Context) {
	userId := ctx.Param("id")

	err := c.usersService.DeleteUser(userId)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *UsersController) UpdateUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	body := dto.UpdateUserRequestBody{}

	// getting request's body
	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	response, err := c.usersService.UpdateUser(userId, &body)
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusCreated, &response)
}
