package entity

import (
	"FP-DevOps/utils"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `json:"id" form:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()" `
	Username string    `json:"username" form:"username" gorm:"uniqueIndex;not null"`
	Password string    `json:"password" form:"password"`

	Timestamp
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	return nil
}
