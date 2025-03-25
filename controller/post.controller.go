package controller

import (
	"net/http"

	"github.com/AzkaAzkun/mini-threads-api/dto"
	"github.com/AzkaAzkun/mini-threads-api/middleware"
	"github.com/AzkaAzkun/mini-threads-api/service"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService service.IPostService
}

func NewPost(postService service.IPostService) PostController {
	return PostController{postService}
}

func (c *PostController) Create(ctx *gin.Context) {
	var param dto.PostCreate
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "failed to get files"})
		return
	}

	param.Image = form.File["images"]
	param.UserId = ctx.MustGet("user_id").(string)

	postId, err := c.postService.Create(ctx.Request.Context(), param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Post registered successfully",
		"post_id": postId,
	})
}

func (c *PostController) GetAll(ctx *gin.Context) {
	res, err := c.postService.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get all post record successfully",
		"data":    res,
	})
}

func (c *PostController) GetById(ctx *gin.Context) {
	postId := ctx.Param("post_id")

	res, err := c.postService.GetById(ctx, postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get post by id record successfully",
		"data":    res,
	})
}

func (c *PostController) Edit(ctx *gin.Context) {
	var param dto.PostUpdate
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	param.PostId = ctx.Param("post_id")
	err := c.postService.Edit(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post updated successfully",
	})
}

func (c *PostController) Delete(ctx *gin.Context) {
	postId := ctx.Param("post_id")

	err := c.postService.Delete(ctx, postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}

func PostRoute(routes *gin.RouterGroup, controller PostController) {
	routes.POST("/", middleware.Authenticate(), controller.Create)
	routes.GET("/", middleware.Authenticate(), controller.GetAll)
	routes.GET("/:post_id", middleware.Authenticate(), controller.GetById)
	routes.PUT("/:post_id", middleware.Authenticate(), controller.Edit)
	routes.DELETE("/:post_id", middleware.Authenticate(), controller.Delete)
}
