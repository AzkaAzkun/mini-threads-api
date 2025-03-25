package dto

type LikeCreate struct {
	UserId string
	PostId string `json:"post_id" binding:"required,uuid"`
}
