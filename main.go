package main

import (
	"fmt"
	"log"
	"os"

	"github.com/AzkaAzkun/mini-threads-api/controller"
	"github.com/AzkaAzkun/mini-threads-api/database"
	"github.com/AzkaAzkun/mini-threads-api/repository"
	"github.com/AzkaAzkun/mini-threads-api/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("Failed to loading env file")
	}
	db := database.New()

	if len(os.Args) > 1 {
		database.Commands(db)
		return
	}

	app := gin.Default()
	var (
		userRepository      repository.IUserRepository      = repository.NewUser(db)
		postRepository      repository.IPostRepository      = repository.NewPost(db)
		postImageRepository repository.IPostImageRepository = repository.NewPostImage(db)
		commentRepository   repository.ICommentRepository   = repository.NewComment(db)
		likeRepository      repository.ILikeRepository      = repository.NewLike(db)

		userService      service.IUserService      = service.NewUser(userRepository, db)
		postService      service.IPostService      = service.NewPost(postRepository, db)
		postImageService service.IPostImageService = service.NewPostImage(postImageRepository, db)
		commentService   service.ICommentService   = service.NewComment(commentRepository, postRepository, db)
		likeService      service.ILikeService      = service.NewLike(likeRepository, postRepository, db)

		userController      controller.UserController      = controller.NewUser(userService)
		postController      controller.PostController      = controller.NewPost(postService)
		postImageController controller.PostImageController = controller.NewPostImage(postImageService)
		commentController   controller.CommentController   = controller.NewComment(commentService)
		likeController      controller.LikeController      = controller.NewLike(likeService)
	)

	apiGroup := app.Group("/api")
	controller.UserRoute(apiGroup.Group("/users"), userController)
	controller.CommentRoute(apiGroup.Group("/posts/:post_id/comments"), commentController)
	controller.PostRoute(apiGroup.Group("/posts"), postController)
	controller.PostImageRoute(apiGroup.Group("/posts"), postImageController)
	controller.LikeRoute(apiGroup.Group("/likes"), likeController)
	app.Static("/api/assets", "./assets")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	var serve string
	if os.Getenv("APP_ENV") == "localhost" {
		serve = "127.0.0.1:" + port
	} else {
		serve = ":" + port
	}

	fmt.Printf("Starting server on %s\n", serve)
	if err := app.Run(serve); err != nil {
		log.Fatalf("error running server: %v", err)
	}
}
