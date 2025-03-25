package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	UserId uuid.UUID `gorm:"not null" json:"user_id"`

	Title     string `gorm:"not null" json:"title"`
	Body      string `gorm:"not null" json:"body"`
	LikeCount int    `gorm:"" json:"like_count"`

	Comment   []Comment   `json:"-"`
	Like      []Like      `json:"like"`
	PostImage []PostImage `json:"post_image"`

	CreatedAt time.Time      `gorm:"type:timestamp without time zone" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp without time zone" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
