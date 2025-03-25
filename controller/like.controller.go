package controller

import (
	"net/http"

	"github.com/AzkaAzkun/mini-threads-api/dto"
	"github.com/AzkaAzkun/mini-threads-api/middleware"
	"github.com/AzkaAzkun/mini-threads-api/service"
	"github.com/gin-gonic/gin"
)

type LikeController struct {
	likeService service.ILikeService
}

func NewLike(likeService service.ILikeService) LikeController {
	return LikeController{likeService}
}

func (c *LikeController) Create(ctx *gin.Context) {
	var param dto.LikeCreate
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	param.UserId = ctx.MustGet("user_id").(string)
	likeId, err := c.likeService.Create(ctx.Request.Context(), param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Like create successfully",
		"like_id": likeId,
	})
}

func (c *LikeController) Delete(ctx *gin.Context) {
	likeId := ctx.Param("like_id")

	err := c.likeService.Delete(ctx, likeId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Like deleted successfully",
	})
}

func LikeRoute(routes *gin.RouterGroup, controller LikeController) {
	routes.POST("", middleware.Authenticate(), controller.Create)
	routes.DELETE("/:like_id", middleware.Authenticate(), controller.Delete)
}
