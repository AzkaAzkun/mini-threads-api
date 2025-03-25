package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostImage struct {
	ID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`
	PostId uuid.UUID `gorm:"not null" json:"post_id"`

	ImagePath string `gorm:"not null" json:"image_path"`

	CreatedAt time.Time      `gorm:"type:timestamp without time zone" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp without time zone" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}
