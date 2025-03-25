package entity

import (
	"time"

	"github.com/AzkaAzkun/mini-threads-api/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"id"`

	Email    string `gorm:"uniqueIndex:not null" json:"email"`
	Name     string `gorm:"not null" json:"name"`
	Password string `gorm:"not null" json:"password"`

	CreatedAt time.Time      `gorm:"type:timestamp without time zone" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp without time zone" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
