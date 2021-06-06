package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityCart struct {
	Id         string    `gorm:"primaryKey;" json:"id"`
	User_Id    string    `gorm:"type:varchar(255);not null" json:"user_id"`
	Order_Id   string    `gorm:"type:varchar(255);not null" json:"order_id"`
	Product_Id string    `gorm:"type:varchar(255);not null" json:"product_id"`
	Quantity   int       `gorm:"type:int;not null" json:"quantity"`
	Note       string    `gorm:"type:varchar(255);" json:"note"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (entity *EntityCart) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityCart) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
