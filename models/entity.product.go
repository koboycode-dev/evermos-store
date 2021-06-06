package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type EntityProduct struct {
	Id         string    `gorm:"primaryKey;" json:"id"`
	Id_Product string    `gorm:"type:varchar(50);unique;not null" json:"id_product"`
	Name       string    `gorm:"type:varchar(50);unique;not null" json:"name"`
	Price      int       `gorm:"type:int;not null" json:"price"`
	Quantity   int       `gorm:"type:int;not null" json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (entity *EntityProduct) BeforeCreate(db *gorm.DB) error {
	entity.Id = uuid.New().String()
	entity.CreatedAt = time.Now().Local()
	return nil
}

func (entity *EntityProduct) BeforeUpdate(db *gorm.DB) error {
	entity.UpdatedAt = time.Now().Local()
	return nil
}
