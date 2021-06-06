package resultProduct

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	ResultProductRepository(input *model.EntityProduct) (*model.EntityProduct, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResult(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResultProductRepository(input *model.EntityProduct) (*model.EntityProduct, string) {

	var Products model.EntityProduct
	db := r.db.Model(&Products)
	errorCode := make(chan string, 1)

	resultProducts := db.Debug().Select("id, Id_Product, name, price, quantity, created_at, updated_at").Where("id = ?", input.Id).Find(&Products)

	if resultProducts.RowsAffected < 1 {
		errorCode <- util.NOT_FOUND
		return &Products, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Products, <-errorCode
}
