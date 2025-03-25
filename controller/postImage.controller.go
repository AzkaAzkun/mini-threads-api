package controller

import (
	"net/http"

	"github.com/AzkaAzkun/mini-threads-api/dto"
	"github.com/AzkaAzkun/mini-threads-api/middleware"
	"github.com/AzkaAzkun/mini-threads-api/service"
	"github.com/gin-gonic/gin"
)

type PostImageController struct {
	postImageService service.IPostImageService
}

func NewPostImage(postImageService service.IPostImageService) PostImageController {
	return PostImageController{postImageService}
}

func (c *PostImageController) Create(ctx *gin.Context) {
	var param dto.PostImageCreate
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	param.PostId = ctx.Param("post_id")
	postImageId, err := c.postImageService.Create(ctx.Request.Context(), param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":       "Post image create successfully",
		"post_image_id": postImageId,
	})
}

func (c *PostImageController) Delete(ctx *gin.Context) {
	postImageId := ctx.Param("post_image_id")

	err := c.postImageService.Delete(ctx, postImageId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post image deleted successfully",
	})
}

func PostImageRoute(routes *gin.RouterGroup, controller PostImageController) {
	routes.POST("/:post_id/images", middleware.Authenticate(), controller.Create)
	routes.DELETE("/:post_id/images/:post_image_id", middleware.Authenticate(), controller.Delete)
}
