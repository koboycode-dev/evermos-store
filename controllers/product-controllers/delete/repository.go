package deleteProduct

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	DeletedProductRepository(input *model.EntityProduct) (*model.EntityProduct, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryDelete(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) DeletedProductRepository(input *model.EntityProduct) (*model.EntityProduct, string) {

	var dProducts model.EntityProduct
	db := r.db.Model(&dProducts)
	errorCode := make(chan string, 1)

	checkdProductId := db.Debug().Select("*").Where("id = ?", input.Id).Find(&dProducts)

	if checkdProductId.RowsAffected < 1 {
		errorCode <- util.NOT_FOUND
		return &dProducts, <-errorCode
	}

	deletedProductId := db.Debug().Select("name").Where("id = ?", input.Id).Find(&dProducts).Delete(&dProducts)

	if deletedProductId.Error != nil {
		errorCode <- util.FAILED
		return &dProducts, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &dProducts, <-errorCode
}
