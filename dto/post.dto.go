package dto

import "mime/multipart"

type PostCreate struct {
	UserId string
	Title  string                  `form:"title" binding:"required"`
	Body   string                  `form:"body" binding:"required"`
	Image  []*multipart.FileHeader `form:"image"`
}

type PostUpdate struct {
	PostId string
	Title  string `json:"title" binding:"required"`
	Body   string `json:"body" binding:"required"`
}
