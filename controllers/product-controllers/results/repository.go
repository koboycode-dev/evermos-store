package resultsProduct

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	ResultsProductRepository() (*[]model.EntityProduct, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResultsProductRepository() (*[]model.EntityProduct, string) {

	var Products []model.EntityProduct
	db := r.db.Model(&Products)
	errorCode := make(chan string, 1)

	resultsProducts := db.Debug().Select("id, id_product, name, price, quantity, created_at, updated_at").Find(&Products)

	if resultsProducts.Error != nil {
		errorCode <- util.NOT_FOUND
		return &Products, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Products, <-errorCode
}
