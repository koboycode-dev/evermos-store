package resultsCart

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	ResultsCartRepository() (*[]model.EntityCart, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResults(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResultsCartRepository() (*[]model.EntityCart, string) {

	var Carts []model.EntityCart
	db := r.db.Model(&Carts)
	errorCode := make(chan string, 1)

	resultsCarts := db.Debug().Select("id, order_id, product_id, note, quantity, created_at, updated_at").Find(&Carts)

	if resultsCarts.Error != nil {
		errorCode <- util.NOT_FOUND
		return &Carts, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Carts, <-errorCode
}
