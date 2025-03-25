package dto

import "mime/multipart"

type PostImageCreate struct {
	PostId string
	Image  *multipart.FileHeader `form:"image" binding:"required"`
}
