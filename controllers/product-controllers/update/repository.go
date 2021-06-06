package updateProduct

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	UpdateProductRepository(input *model.EntityProduct) (*model.EntityProduct, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryUpdate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) UpdateProductRepository(input *model.EntityProduct) (*model.EntityProduct, string) {

	var Products model.EntityProduct
	db := r.db.Model(&Products)
	errorCode := make(chan string, 1)

	Products.Id = input.Id

	checkProductId := db.Debug().Select("name").Where("id = ?", input.Id).Find(&Products)

	if checkProductId.RowsAffected < 1 {
		errorCode <- util.NOT_FOUND
		return &Products, <-errorCode
	}

	Products.Name = input.Name
	Products.Price = input.Price
	Products.Quantity = input.Quantity

	updateProduct := db.Debug().Select("name", "price", "quantity", "updated_at").Where("id = ?", input.Id).Updates(Products)

	if updateProduct.Error != nil {
		errorCode <- util.FAILED
		return &Products, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Products, <-errorCode
}
