package dto

type CommentCreate struct {
	UserId string
	PostId string
	Body   string `json:"body" binding:"required"`
}

type CommentUpdate struct {
	CommentId string
	Body      string `json:"body" binding:"required"`
}
