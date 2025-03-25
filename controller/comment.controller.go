package controller

import (
	"net/http"

	"github.com/AzkaAzkun/mini-threads-api/dto"
	"github.com/AzkaAzkun/mini-threads-api/middleware"
	"github.com/AzkaAzkun/mini-threads-api/service"
	"github.com/gin-gonic/gin"
)

type CommentController struct {
	commentService service.ICommentService
}

func NewComment(commentService service.ICommentService) CommentController {
	return CommentController{commentService}
}

func (c *CommentController) Create(ctx *gin.Context) {
	var param dto.CommentCreate
	if err := ctx.ShouldBind(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	param.UserId = ctx.MustGet("user_id").(string)
	param.PostId = ctx.Param("post_id")

	commentId, err := c.commentService.Create(ctx.Request.Context(), param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":    "Comment registered successfully",
		"comment_id": commentId,
	})
}

func (c *CommentController) GetAll(ctx *gin.Context) {
	postId := ctx.Param("post_id")
	res, err := c.commentService.GetAllByPostId(ctx, postId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "get all comment record by post id successfully",
		"data":    res,
	})
}

func (c *CommentController) Edit(ctx *gin.Context) {
	var param dto.CommentUpdate
	if err := ctx.ShouldBindJSON(&param); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	param.CommentId = ctx.Param("comment_id")
	err := c.commentService.Edit(ctx, param)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comment updated successfully",
	})
}

func (c *CommentController) Delete(ctx *gin.Context) {
	commentId := ctx.Param("comment_id")

	err := c.commentService.Delete(ctx, commentId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Comment deleted successfully",
	})
}

func CommentRoute(routes *gin.RouterGroup, controller CommentController) {
	routes.POST("/", middleware.Authenticate(), controller.Create)
	routes.GET("/", middleware.Authenticate(), controller.GetAll)
	routes.PUT("/:comment_id", middleware.Authenticate(), controller.Edit)
	routes.DELETE("/:comment_id", middleware.Authenticate(), controller.Delete)
}
