package createProduct

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	CreateProductRepository(input *model.EntityProduct) (*model.EntityProduct, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryCreate(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) CreateProductRepository(input *model.EntityProduct) (*model.EntityProduct, string) {

	var products model.EntityProduct
	db := r.db.Model(&products)
	errorCode := make(chan string, 1)

	checkProductExist := db.Debug().Select("id_product").Where("id_product = ?", input.Id_Product).Find(&products)

	if checkProductExist.RowsAffected > 0 {
		errorCode <- util.CONFLICT
		return &products, <-errorCode
	}

	products.Id_Product = input.Id_Product
	products.Name = input.Name
	products.Price = input.Price
	products.Quantity = input.Quantity

	addNewProduct := db.Debug().Create(&products)
	db.Commit()

	if addNewProduct.Error != nil {
		errorCode <- util.FAILED
		return &products, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &products, <-errorCode
}
