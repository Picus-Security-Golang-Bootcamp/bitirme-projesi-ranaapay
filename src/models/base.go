package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	Id        string    `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `gorm:"index" json:"deletedAt"`
	IsDeleted bool      `json:"isDeleted"`
}

func (b *Base) BeforeCreate(tx *gorm.DB) (err error) {
	b.Id = uuid.New().String()
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	b.IsDeleted = false
	return
}

func (b *Base) GetUpdatedAt() time.Time {
	return b.UpdatedAt
}
func (b *Base) SetUpdatedAt() {
	b.UpdatedAt = time.Now()
}
