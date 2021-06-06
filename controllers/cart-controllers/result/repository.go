package resultCart

import (
	model "github.com/firmanJS/store-app/models"
	util "github.com/firmanJS/store-app/utils"
	"gorm.io/gorm"
)

type Repository interface {
	ResultCartRepository(input *model.EntityCart) (*model.EntityCart, string)
}

type repository struct {
	db *gorm.DB
}

func NewRepositoryResult(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (r *repository) ResultCartRepository(input *model.EntityCart) (*model.EntityCart, string) {

	var Carts model.EntityCart
	db := r.db.Model(&Carts)
	errorCode := make(chan string, 1)

	resultCarts := db.Debug().Select("id, order_id, product_id, note, quantity, created_at, updated_at").Where("id = ?", input.Id).Find(&Carts)

	if resultCarts.RowsAffected < 1 {
		errorCode <- util.NOT_FOUND
		return &Carts, <-errorCode
	} else {
		errorCode <- "nil"
	}

	return &Carts, <-errorCode
}
