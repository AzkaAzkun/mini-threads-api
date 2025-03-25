package controller

import (
	"net/http"

	"github.com/AzkaAzkun/mini-threads-api/dto"
	"github.com/AzkaAzkun/mini-threads-api/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUser(userService service.IUserService) UserController {
	return UserController{userService}
}

func (c *UserController) RegisterAccount(ctx *gin.Context) {
	var param dto.UserCreate
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := c.userService.RegisterAccount(ctx.Request.Context(), param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully",
		"user_id": userId,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	var param dto.UserLoginRequest
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := c.userService.Login(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "User login successfully",
		"data":    res,
	})
}

func UserRoute(routes *gin.RouterGroup, controller UserController) {
	routes.POST("/register", controller.RegisterAccount)
	routes.POST("/login", controller.Login)
}
